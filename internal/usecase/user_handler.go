package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/db"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

func NewUserHandler(mc *coflnet.McConnectApi, p *coflnet.PaymentApi, userDb *db.UserRepo) *UserHandler {
	u := &UserHandler{}
	u.tracer = otel.Tracer("user-handler")
	u.mcConnectApi = mc
	u.paymentApi = p
	u.userRepo = userDb

	return u
}

type UserHandler struct {
	tracer       trace.Tracer
	mcConnectApi *coflnet.McConnectApi
	paymentApi   *coflnet.PaymentApi
	userRepo     *db.UserRepo
}

func (u *UserHandler) RefreshUserById(ctx context.Context, id int) error {
	return errors.New("not implemented")
}

func (u *UserHandler) RefreshUserByDiscordId() error {
	return errors.New("not implemented")
}

func (u *UserHandler) RefreshUserByDiscordUser(ctx context.Context, user *discordgo.User) error {
	ctx, span := u.tracer.Start(ctx, "refresh-user-by-discord-user")
	defer span.End()

	span.SetAttributes(attribute.String("discord-id", user.ID))
	span.SetAttributes(attribute.String("discord-username", user.Username))
	span.SetAttributes(attribute.String("discord-discriminator", user.Discriminator))

	// check if there is a user associated with the discord username and discriminator
	// if yes replace it with the new system

	err := u.checkForDeprecatedUsernameDiscriminator(ctx, *user)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to check for deprecated username and discriminator; traceId: %s", span.SpanContext().TraceID()), err)
		span.RecordError(err)
		return err
	}

	// check if there is a user associated with the discord id
	dbUser, err := u.userRepo.SearchUserByDiscordId(ctx, user.ID)
	if err != nil {
		// check if it is a user not found error
		if _, ok := err.(*model.UserNotFoundError); ok {
			dbUser, err = u.createUserFromDiscordUser(ctx, user)
			if err != nil {
				slog.Error(fmt.Sprintf("failed to create user from discord user; traceId: %s", span.SpanContext().TraceID()), err)
				span.RecordError(err)
				return err
			}
		}
	}

	// update all the properties of the user
	if dbUser == nil || dbUser.UserId == 0 {
		slog.Warn(fmt.Sprintf("user with discord username %s has no cofl user id, can not continue", user.Username))
		return nil
	}

	slog.Warn(fmt.Sprintf("user with discord username %s has cofl user id %d, can continue", user.Username, dbUser.UserId))

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

func (u *UserHandler) RefreshUser(ctx context.Context, user *model.User) error {
	_, span := u.tracer.Start(ctx, "refresh-user")
	defer span.End()
	span.SetAttributes(attribute.Int("user-id", user.UserId))
	span.SetAttributes(attribute.String("last-refresh", user.LastRefresh.String()))

	// if the user was last refreshed less than 5 minutes ago, skip
	if user.LastRefresh.After(time.Now().Add(-5 * time.Minute)) {
		slog.Debug("skipping refresh of user %d; user was refreshed at %v", user.UserId, user.LastRefresh)
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
	slog.Debug(fmt.Sprintf("user %d owns %d products", user.UserId, len(ownedProducts)))

	// load the hypixel information
	slog.Warn("loading hypixel information is not implemented")

	// load the roles the user should have
	slog.Warn("loading roles is not implemented")

	// apply the roles to the user
	slog.Warn("applying roles is not implemented")

	// update the user in the database
	slog.Warn("updating user in database is not implemented")

	return errors.New("not implemented")
}
