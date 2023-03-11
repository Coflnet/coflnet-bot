package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

type UserHandler struct {
    tracer trace.Tracer
    mcConnectApi *coflnet.McConnectApi
}

func NewUserHandler() (*UserHandler, error) {
    var err error
	u := &UserHandler{}

    u.tracer = otel.Tracer(flipProcessorTracerName)
    u.mcConnectApi, err = coflnet.NewMcConnectClient()

    if err != nil {
        slog.Error("failed to initialize mc connect api client", err)
        return nil, err
    }

	return u, nil
}

func (u *UserHandler) RefreshUserId(ctx context.Context, id int) error {
    _, span := u.tracer.Start(ctx, "refresh-user-id")
    defer span.End()

    user, err := mongo.UserById(ctx, id)
    if err != nil {
        slog.Error("failed to load user", err)
        span.RecordError(err)
    }

    return u.RefreshUser(ctx, user)
}

func (u *UserHandler) RefreshUserDiscordId() error {
    return errors.New("not implemented")
}

func (u *UserHandler) RefreshUser(ctx context.Context, user *model.User) error {
    _, span := u.tracer.Start(ctx, "refresh-user")
    defer span.End()

    // if the user was last refreshed less than 5 minutes ago, skip
    if user.LastRefresh.After(time.Now().Add(-5 * time.Minute)) {
        slog.Debug("skipping refresh of user %d; user was refreshed at %v", user.UserId, user.LastRefresh)
        return nil
    }

    // load the  mc connet information
    mcConnectUser, err := u.mcConnectApi.GetPlayer(ctx, user.UserId)
    if err != nil {
        slog.Error("failed to load user from mc connect api", err)
        span.RecordError(err)
        return err
    }
    slog.Debug("loaded user from mc connect api: %v", mcConnectUser)

    // load the payment information
    slog.Warn("loading payment information is not implemented")

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
