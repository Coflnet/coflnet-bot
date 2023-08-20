package coflnet

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/Coflnet/coflnet-bot/codegen/mcconnect"
	"github.com/Coflnet/coflnet-bot/internal/utils"
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
	r.apiClient, err = mcconnect.NewClientWithResponses(apiClientUrl)
	if err != nil {
		slog.Error("error creating mc connect api client", err)
		return r
	}

	return r
}

type McConnectApi struct {
	apiClient *mcconnect.ClientWithResponses
	tracer    trace.Tracer
}

func (a *McConnectApi) GetPlayer(ctx context.Context, id int) (*mcconnect.User, error) {
	if a.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	ctx, span := a.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))

	response, err := a.apiClient.GetConnectUserUserIdWithResponse(ctx, strconv.Itoa(id))
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	switch response.StatusCode() {
	case 200:
		span.SetAttributes(attribute.Int("status", 200))
		slog.Debug("got user from mc connect api", "id", id)
		user := response.JSON200
		return user, nil
	default:
		span.SetAttributes(attribute.Int("status", response.StatusCode()))
		slog.Warn("error getting user from mc connect api", "id", id, "status", response.StatusCode())
		return nil, errors.New("error getting user from mc connect api")
	}
}

func (a *McConnectApi) PlayerByUUID(ctx context.Context, uuid string) (*mcconnect.User, error) {
	if a.apiClient == nil {
		return nil, errors.New("mc connect api client not initialized")
	}

	ctx, span := a.tracer.Start(ctx, "get-player-from-mc-connect-api")
	defer span.End()

	span.SetAttributes(attribute.String("uuid", uuid))

	user, err := a.apiClient.GetConnectMinecraftMcUuidWithResponse(ctx, uuid)

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	switch user.StatusCode() {
	case 200:
		slog.Debug("got user from mc connect api", "uuid", uuid)
		span.SetAttributes(attribute.Int("status", 200))
		return user.JSON200, nil
	default:
		slog.Warn("error getting user from mc connect api", "uuid", uuid, "status", user.StatusCode())
		span.SetAttributes(attribute.Int("status", user.StatusCode()))
		return nil, errors.New("error getting user from mc connect api")
	}
}
