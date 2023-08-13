package coflnet

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/mc_connect"
)

const (
	mcConnectApiName = "mc-connect-api"
)

func NewMcConnectApi() *McConnectApi {
	r := &McConnectApi{}
	r.tracer = otel.Tracer(mcConnectApiName)

	apiClientUrl, err := utils.McConnectBaseUrl()
	if err != nil {
		slog.Error("error getting mc connect api url", err)
		return r
	}
	r.apiClient, err = mc_connect.NewClient(apiClientUrl)
	if err != nil {
		slog.Error("error creating mc connect api client", err)
		return r
	}

	return r
}

type McConnectApi struct {
	apiClient *mc_connect.Client
	tracer    trace.Tracer
}

func (a *McConnectApi) GetPlayer(ctx context.Context, id int) (*mc_connect.User, error) {
	if a.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	ctx, span := a.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))

	user, err := a.apiClient.ConnectUserUserIdGet(ctx, mc_connect.ConnectUserUserIdGetParams{
		UserId: strconv.Itoa(id),
	})

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return user, nil
}

func (a *McConnectApi) PlayerByUUID(ctx context.Context, uuid string) (*mc_connect.User, error) {
	if a.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	ctx, span := a.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.String("uuid", uuid))

	user, err := a.apiClient.ConnectMinecraftMcUuidGet(ctx, mc_connect.ConnectMinecraftMcUuidGetParams{
		McUuid: uuid,
	})

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if user.ID.IsSet() {
		span.SetAttributes(attribute.Int("id", int(user.ID.Value)))
	} else {
		span.SetAttributes(attribute.Int("id", -1))
		return nil, errors.New("user id not set")
	}

	return user, nil
}
