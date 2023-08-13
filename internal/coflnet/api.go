package coflnet

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ApiClient struct {
	tracer    trace.Tracer
	apiClient *api.Client
}

func NewApiClient() *ApiClient {
	c := &ApiClient{
		tracer: otel.Tracer("api-client"),
	}

	apiUrl, err := utils.ApiBaseUrl()
	if err != nil {
		slog.Error("missing api base url, continue without that functionality", slog.String("err", err.Error()))
		return c
	}

	c.apiClient, err = api.NewClient(apiUrl)
	if err != nil {
		panic(err)
	}

	return c
}

func (m *ApiClient) SearchUUIDForPlayer(ctx context.Context, username string) ([]string, error) {
	ctx, span := m.tracer.Start(ctx, "search-uuid-for-player")
	defer span.End()

	span.SetAttributes(attribute.String("player_name", username))

	result := make([]string, 0)

	playerResults, err := m.apiClient.APISearchPlayerPlayerNameGet(ctx, api.APISearchPlayerPlayerNameGetParams{
		PlayerName: username,
	})

	if err != nil {
		slog.Error("failed to get player info", err)
		span.RecordError(err)
		return result, err
	}

	span.SetAttributes(attribute.Int("player_count", len(playerResults)))

	if len(playerResults) == 0 {
		slog.Warn("no player found")
		return result, errors.New(fmt.Sprintf("no player with name %s found", username))
	}

	for _, player := range playerResults {
		result = append(result, player.UUID.Value)
	}

	span.SetAttributes(attribute.Int("valid_uuids", len(result)))
	span.SetAttributes(attribute.String("uuids", strings.Join(result, ",")))

	return result, nil
}
