package usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/db"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/hypixel"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

func NewUserHandler(mc *coflnet.McConnectApi, p *coflnet.PaymentApi, userDb *db.UserRepo, discordHandler *discord.DiscordHandler) *UserHandler {
	u := &UserHandler{}
	u.tracer = otel.Tracer("user-handler")
	u.mcConnectApi = mc
	u.paymentApi = p
	u.userRepo = userDb
	u.discordHandler = discordHandler

	return u
}

type UserHandler struct {
	tracer         trace.Tracer
	mcConnectApi   *coflnet.McConnectApi
	paymentApi     *coflnet.PaymentApi
	userRepo       *db.UserRepo
	discordHandler *discord.DiscordHandler
}

func (u *UserHandler) RefreshUserByUUID(ctx context.Context, uuid string) error {
	ctx, span := u.tracer.Start(ctx, "refresh-user-by-uuid")
	defer span.End()
	span.SetAttributes(attribute.String("uuid", uuid))

	users, err := u.userRepo.SearchUserByUUID(ctx, uuid)
	if err != nil {
		if _, ok := err.(*model.UserNotFoundError); !ok {
			slog.Error(fmt.Sprintf("failed to search user by uuid; traceId: %s", span.SpanContext().TraceID()), err)
			span.RecordError(err)
			return err
		}
	}

	// create a empty user if there is no user with the given uuid
	if len(users) == 0 {
		slog.Info(fmt.Sprintf("user with uuid %s not found; traceId: %s", uuid, span.SpanContext().TraceID()))
		user, err := u.createUserFromMinecraftUUID(ctx, uuid)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to create user from minecraft uuid; traceId: %s", span.SpanContext().TraceID()), err)
			span.RecordError(err)
			return err
		}
		slog.Info(fmt.Sprintf("created user with uuid %s; traceId: %s", uuid, span.SpanContext().TraceID()))
		users = []*model.User{user}
	} else {
		slog.Info(fmt.Sprintf("a user with the uuid %s was found, no need to create a new one; traceId: %s", uuid, span.SpanContext().TraceID()))
	}

	// update the given users
	slog.Info(fmt.Sprintf("updating users with uuid %s; traceId: %s", uuid, span.SpanContext().TraceID()))
	for _, user := range users {
		err := u.RefreshUser(ctx, user)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to refresh user by id; traceId: %s", span.SpanContext().TraceID()), err)
			span.RecordError(err)
			return err
		}
	}

	return nil
}

func (u *UserHandler) checkForDeprecatedUsernameDiscriminator(ctx context.Context, user discordgo.User) error {
	ctx, span := u.tracer.Start(ctx, "check-for-deprecated-username-discriminator")
	defer span.End()

	searchTerm := fmt.Sprintf("%s#%s", user.Username, user.Discriminator)
	span.SetAttributes(attribute.String("search-term", searchTerm))

	dbUser, err := u.userRepo.SearchByDiscordTag(ctx, searchTerm)
	if err != nil {
		if _, ok := err.(*model.UserNotFoundError); !ok {
			slog.Error(fmt.Sprintf("failed to search for deprecated user; traceId: %s", span.SpanContext().TraceID()), err)
			span.RecordError(err)
			return err
		}
	}

	if dbUser == nil {
		slog.Debug("no deprecated user found")
		span.SetAttributes(attribute.Bool("deprecated-user-found", false))
		return nil
	}

	span.SetAttributes(attribute.Bool("deprecated-user-found", true))
	slog.Info(fmt.Sprintf("found deprecated user with id %d set discord id %s", dbUser.UserId, user.ID))

	// set the new discord id
	ids := []string{user.ID}
	err = u.userRepo.SetDiscordIdForUser(ctx, dbUser.UserId, ids)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to set new discord id for user; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return err
	}

	// remove the old discord id
	err = u.userRepo.UnsetDiscordTag(ctx, dbUser.UserId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to unset deprecated discord id for user; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return err
	}

	return nil
}

func (u *UserHandler) createUserFromDiscordUser(ctx context.Context, user *discordgo.User) (*model.User, error) {
	ctx, span := u.tracer.Start(ctx, "create-user-from-discord-user")
	defer span.End()

	modelUser := &model.User{
		DiscordIds:         []string{user.ID},
		PreferredDiscordId: user.ID,
		LastRefresh:        time.Now(),
	}

	err := u.userRepo.InsertEmptyModelUser(ctx, modelUser)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to insert empty model user; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return nil, err
	}

	return modelUser, nil
}

func (u *UserHandler) createUserFromMinecraftUUID(ctx context.Context, uuid string) (*model.User, error) {
	ctx, span := u.tracer.Start(ctx, "create-user-from-minecraft-uuid")
	defer span.End()
	span.SetAttributes(attribute.String("minecraft-uuid", uuid))

	mcUser, err := u.mcConnectApi.PlayerByUUID(ctx, uuid)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get minecraft user by uuid calling mc connect; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return nil, err
	}

	if mcUser == nil {
		slog.Error("mc connect api returned nil user", "uuid", uuid)
		return nil, errors.New("mc connect api returned nil user")
	}

	if mcUser.Id == nil {
		slog.Error("mc user has no id", "uuid", uuid)
		return nil, errors.New("user has not id")
	}

	modelUser := &model.User{
		MinecraftUuids: []string{uuid},
		PreferredUUID:  uuid,
		LastRefresh:    time.Now(),
		UserId:         int(*mcUser.Id),
	}

	err = u.userRepo.InsertEmptyModelUser(ctx, modelUser)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to insert empty model user; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return nil, err
	}

	return modelUser, nil
}

func (u *UserHandler) RefreshUser(ctx context.Context, user *model.User) error {
	_, span := u.tracer.Start(ctx, "refresh-user")
	defer span.End()
	span.SetAttributes(attribute.Int("user-id", user.UserId))
	span.SetAttributes(attribute.String("last-refresh", user.LastRefresh.String()))

	// if the user was last refreshed less than 5 minutes ago, skip
	if user.LastRefresh.After(time.Now().Add(-5 * time.Minute)) {
		slog.Info("skipping refresh of user %d; user was refreshed at %v", user.UserId, user.LastRefresh)
		span.SetAttributes(attribute.Bool("skipped", true))
		return nil
	}

	// load the  mc connet information
	mcConnectUser, err := u.mcConnectApi.GetPlayer(ctx, user.UserId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to load user from mc connect api; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return err
	}
	slog.Debug("loaded user from mc connect api: %v", mcConnectUser)

	// load the payment information
	ownedProducts, err := u.paymentApi.OwningTimesOfUser(ctx, user.UserId)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to load user from payment api; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
	}

	slog.Info(fmt.Sprintf("user %d owns %d products", user.UserId, len(ownedProducts)))

	// load the hypixel information
	slog.Warn("loading hypixel information..")
	discord_members := make(chan *discordgo.Member)

	// load the discord names from the hypixel api
	go func() {
		wg := &sync.WaitGroup{}
		for _, uuid := range user.MinecraftUuids {
			wg.Add(1)
			go func(uuid string) {
				defer wg.Done()

				hypixelData, err := hypixel.PlayerData(uuid)
				if err != nil {
					slog.Warn("failed to load hypixel data", "uuid", uuid, "error", err)
					span.RecordError(err)
					return
				}

				slog.Info("loaded hypixel data", "uuid", uuid)
				discord_name := hypixelData.Player.SocialMedia.Links.Discord

				// check if a user with that discord name exists
				member, err := u.discordHandler.SearchDiscordUser(discord_name)
				if err != nil {
					slog.Error("failed to search discord user", "discord_name", discord_name, "error", err)
					span.RecordError(err)
					return
				}

				discord_members <- member
			}(uuid)
		}
		wg.Wait()
		close(discord_members)
	}()

	for member := range discord_members {
		if member == nil {
			continue
		}

		slog.Info("found discord member for user", "member_name", member.User.Username, "uuid", user.UUID(), "member_id", member.User.ID)

		// add the discord id to the existing user
		if slices.Contains(user.DiscordIds, member.User.ID) {
			slog.Debug("user already has that discord id", "discord_id", member.User.ID, "uuid", user.UUID())
			continue
		}

		user.DiscordIds = append(user.DiscordIds, member.User.ID)
	}

	err = u.userRepo.SetDiscordIdForUser(ctx, user.UserId, user.DiscordIds)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to set discord ids for user; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return err
	}
	span.SetAttributes(attribute.Int("discord-ids-set", len(user.DiscordIds)))

	// load the roles the user should have
	slog.Warn("loading roles is not implemented")

	// apply the roles to the user
	slog.Warn("applying roles is not implemented")

	// update the user in the database
	slog.Warn("updating user in database is not implemented")

	return errors.New("not implemented")
}
