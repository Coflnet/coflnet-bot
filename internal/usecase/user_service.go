package usecase

import (
	"coflnet-bot/internal/db"
	"coflnet-bot/internal/gen"
	mcconnectgen "coflnet-bot/internal/gen/mcconnect"
	paymentgen "coflnet-bot/internal/gen/payment"
	proxygen "coflnet-bot/internal/gen/proxy"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"
)

type UserService struct {
	tracer          trace.Tracer
	paymentClient   *paymentgen.ClientWithResponses
	proxyClient     *proxygen.ClientWithResponses
	mcConnectClient *mcconnectgen.ClientWithResponses
}

func NewUserService(paymentUrl, proxyUrl, mcConnectUrl string) (*UserService, error) {
	paymentClient, err := paymentgen.NewClientWithResponses(paymentUrl, paymentgen.WithRequestEditorFn(addJsonAcceptHeaderFn))
	if err != nil {
		return nil, err
	}
	proxyClient, err := proxygen.NewClientWithResponses(proxyUrl, proxygen.WithRequestEditorFn(addJsonAcceptHeaderFn))
	if err != nil {
		return nil, err
	}
	mcConnectClient, err := mcconnectgen.NewClientWithResponses(mcConnectUrl, mcconnectgen.WithRequestEditorFn(addJsonAcceptHeaderFn))
	if err != nil {
		return nil, err
	}

	return &UserService{
		tracer:          otel.Tracer("user-service"),
		paymentClient:   paymentClient,
		proxyClient:     proxyClient,
		mcConnectClient: mcConnectClient,
	}, nil
}

func (s *UserService) LoadUserByUUID(ctx context.Context, uuid string) (*db.User, error) {
	ctx, span := s.tracer.Start(ctx, "load-user")
	defer span.End()

	// load the external id of the user
	externalId, err := s.LoadExternalId(ctx, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error loading mc user", "err", err)
		return nil, err
	}
	slog.Debug(fmt.Sprintf("found the external id %s for the uuid %s", externalId, uuid))

	// check if the user exists
	dbUser, err := db.UserByExternalId(ctx, externalId)
	if err != nil {
		span.RecordError(err)
		slog.Error("Error loading user", "err", err)
		return nil, err
	}

	user := &db.User{
		ExternalId: externalId,
	}
	if dbUser != nil {
		user = dbUser
		slog.Info(fmt.Sprintf("user loaded from db is no nil, use existing id %d", dbUser.ID))
	} else {
		slog.Info("user loaded from db is nil, create new user")
	}

	// merge the uuid in the object
	s.mergeDBUserWithUUID(user, uuid)

	// only refresh once a minute
	if user.LastUpdated != nil && user.LastUpdated.After(time.Now().Add(-1*time.Minute)) {
		slog.Info(fmt.Sprintf("user with id %s was refreshed less than an hour ago", user.ExternalId))
		return user, nil
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// update the premium status of the player parallel to the hypixel information
	go func() {
		defer wg.Done()
		err = s.UpdatePremiumStatusOfUser(ctx, user)
		if err != nil {
			span.RecordError(err)
			slog.Error("Error updating premium status of user", "err", err)
		}
		slog.Info("updated premium status of user")
	}()

	// update the hypixel information of the player
	go func() {
		defer wg.Done()
		err = s.LoadHypixelInformationOfUUID(ctx, user, uuid)
		if err != nil {
			span.RecordError(err)
			slog.Error("error loading hypixel information of user", "err", err)
		}
		slog.Info("loaded hypixel information of user")
	}()

	wg.Wait()
	slog.Info(fmt.Sprintf("updated user with uuid: %s", uuid))

	// save the user to the database
	err = db.SaveUser(ctx, user)
	if err != nil {
		span.RecordError(err)
		slog.Error("error saving user", "err", err)
		return nil, err
	}

	return &db.User{
		ExternalId: externalId,
	}, nil
}

func (s *UserService) mergeDBUserWithUUID(user *db.User, uuid string) {
	if user.MinecraftAccounts == nil {
		user.MinecraftAccounts = make([]db.MinecraftAccount, 0)
	}

	if slices.ContainsFunc(user.MinecraftAccounts, func(acc db.MinecraftAccount) bool {
		return acc.MinecraftUUID == uuid
	}) {
		slog.Debug(fmt.Sprintf("user with id %s already has minecraft account with uuid %s, skipping", user.ExternalId, uuid))
		return
	}

	// set the new one as preferred
	for _, acc := range user.MinecraftAccounts {
		if acc.Preferred {
			acc.Preferred = false
			break
		}
	}

	slog.Info(fmt.Sprintf("adding minecraft account %s to user %s", uuid, user.ExternalId))
	user.MinecraftAccounts = append(user.MinecraftAccounts, db.MinecraftAccount{
		MinecraftUUID: uuid,
		UserID:        user.ID,
		Preferred:     true,
	})
}

// LoadExternalId LoadMcUser loads a user from the mcconnect service
// this only checks for the externalId of the user
func (s *UserService) LoadExternalId(ctx context.Context, uuid string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "load-external-id")
	defer span.End()

	response, err := s.mcConnectClient.GetConnectMinecraftMcUuidWithResponse(ctx, uuid)
	if err != nil {
		return "", err
	}

	slog.Debug(fmt.Sprintf("send request for uuid %s, got response code: %d", uuid, response.StatusCode()))
	slog.Debug(string(response.Body))

	if response.StatusCode() > 399 {
		return "", fmt.Errorf("expected response code 200, got %d", response.StatusCode())
	}

	// sometimes this returns a 204, which is an error in the mcconnect service
	if response.StatusCode() == 204 || response.JSON200 == nil {
		return "", &UserNotFoundError{SearchTerm: uuid}
	}

	slog.Debug(fmt.Sprintf("found external id: %s", *response.JSON200.ExternalId))

	return *response.JSON200.ExternalId, nil
}

func (s *UserService) LoadProductsOfUser(ctx context.Context, externalId string) (*time.Time, *time.Time, error) {
	ctx, span := s.tracer.Start(ctx, "load-payment-information")
	defer span.End()

	// TODO use this to check
	const premium = "premium"
	const premiumPlus = "premium_plus"

	body := []string{premium, premiumPlus}
	response, err := s.paymentClient.PostUserUserIdOwnsUntilWithResponse(ctx, externalId, body)
	if err != nil {
		return nil, nil, err
	}

	if response == nil {
		return nil, nil, fmt.Errorf("response is nil")
	}

	if response.StatusCode() > 399 {
		return nil, nil, fmt.Errorf("expected response code 200, got %d", response.StatusCode())
	}

	if response.JSON200 == nil {
		slog.Debug(fmt.Sprintf("response: %s", string(response.Body)))
		return nil, nil, fmt.Errorf("response of payment service is nil")
	}

	productMap := response.JSON200
	if productMap == nil {
		return nil, nil, fmt.Errorf("response of payment service is nil")
	}

	var premiumUntil *time.Time
	var premiumPlusUntil *time.Time

	if until, ok := (*productMap)[premium]; ok {
		slog.Debug(fmt.Sprintf("user has premium until %s", until.Format(time.RFC3339)))
		premiumUntil = &until
	} else {
		slog.Debug("user has no premium")
	}

	if until, ok := (*productMap)[premiumPlus]; ok {
		slog.Debug(fmt.Sprintf("user has premium plus until %s", until.Format(time.RFC3339)))
		premiumPlusUntil = &until
	} else {
		slog.Debug("user has no premium plus")
	}

	return premiumUntil, premiumPlusUntil, nil
}

func (s *UserService) UpdatePremiumStatusOfUser(ctx context.Context, user *db.User) error {
	ctx, span := s.tracer.Start(ctx, "update-premium-status")
	defer span.End()

	premiumUntil, premiumPlusUntil, err := s.LoadProductsOfUser(ctx, user.ExternalId)
	if err != nil {
		return err
	}

	slog.Info("loaded products of user")
	if premiumUntil != nil {
		slog.Info("premium until: %v", premiumUntil)
		user.PremiumUntil = premiumUntil
	}
	if premiumPlusUntil != nil {
		slog.Info("premium plus until: %v", premiumPlusUntil)
		user.PremiumPlusUntil = premiumPlusUntil
	}

	return nil
}

func (s *UserService) LoadHypixelInformationOfUUID(ctx context.Context, user *db.User, uuid string) error {
	ctx, span := s.tracer.Start(ctx, "load-hypixel-information")
	defer span.End()

	path := fmt.Sprintf("/player?uuid=%s", uuid)
	slog.Debug(fmt.Sprintf("loading hypixel information for uuid: %s with path: %s", uuid, path))
	params := &proxygen.GetProxyHypixelParams{
		Path: strPtr(path),
	}

	response, err := s.proxyClient.GetProxyHypixelWithResponse(ctx, params)
	if err != nil {
		return err
	}

	hypixelResponse := gen.HypixelPlayerResponse{}
	err = json.Unmarshal(response.Body, &hypixelResponse)
	if err != nil {

		// try something
		var str string
		err = json.Unmarshal(response.Body, &str)
		if err != nil {
			span.RecordError(err)
			return err
		}

		err = json.Unmarshal([]byte(str), &hypixelResponse)
		if err != nil {
			span.SetAttributes(attribute.String("hypixel-response", string(response.Body)))
			fmt.Println(string(response.Body))
			span.RecordError(err)
			return err
		}
	}

	// discordName is the username#discriminator
	discordName := hypixelResponse.Player.SocialMedia.Links.Discord
	// remove the discriminator
	if strings.Contains(discordName, "#") {
		discordName = discordName[:len(discordName)-5]
	}

	discordMember, err := SearchDiscordUser(ctx, discordName)
	if err != nil {
		span.RecordError(err)
		return err
	}

	slog.Info(fmt.Sprintf("loaded discord user with name %s and id %s for uuid %s", discordMember.User.Username, discordMember.User.ID, uuid))

	if user.DiscordAccounts == nil {
		user.DiscordAccounts = make([]db.DiscordAccount, 0)
	}

	if slices.ContainsFunc(user.DiscordAccounts, func(acc db.DiscordAccount) bool {
		return acc.DiscordID == discordMember.User.ID
	}) {
		slog.Info("user already has discord account, skipping")
	} else {

		// set the new one as preferred
		for _, acc := range user.DiscordAccounts {
			if acc.Preferred {
				acc.Preferred = false
				break
			}
		}

		slog.Info("adding discord account to user")
		user.DiscordAccounts = append(user.DiscordAccounts, db.DiscordAccount{
			DiscordID: discordMember.User.ID,
			UserID:    user.ID,
			Preferred: true,
		})
	}

	return nil
}

type UserNotFoundError struct {
	SearchTerm string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with search term %s not found", e.SearchTerm)
}
func addJsonAcceptHeaderFn(_ context.Context, req *http.Request) error {
	// header has to be set always, otherwise response parsing is not working
	req.Header.Set("accept", "application/json")
	return nil
}
