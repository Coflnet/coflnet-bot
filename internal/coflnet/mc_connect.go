package coflnet

import (
	"context"
	"errors"
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
	var err error
	r := &McConnectApi{}

	r.apiClient, err = mc_connect.NewClient(utils.McConnectBaseUrl())
	r.tracer = otel.Tracer(mcConnectApiName)

	if err != nil {
		panic(err)
	}

	return r
}

type McConnectApi struct {
	apiClient *mc_connect.Client
	tracer    trace.Tracer
}

func (a *McConnectApi) GetPlayer(ctx context.Context, id int) (*mc_connect.User, error) {
	ctx, span := a.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))

	if a.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	user, err := a.apiClient.ConnectUserUserIdGet(ctx, mc_connect.ConnectUserUserIdGetParams{
		UserId: strconv.Itoa(id),
	})

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return user, nil
}

func (m *McConnectApi) PlayerByUUID(ctx context.Context, uuid string) (*mc_connect.User, error) {
	ctx, span := m.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.String("uuid", uuid))

	if m.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	user, err := m.apiClient.ConnectMinecraftMcUuidGet(ctx, mc_connect.ConnectMinecraftMcUuidGetParams{
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
