package coflnet

import (
	"context"
	"errors"
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

type ApiClient struct {
	tracer    trace.Tracer
	apiClient *api.Client
}

func NewApiClient() (*ApiClient, error) {
	c := &ApiClient{
		tracer: otel.Tracer("api-client"),
	}

    err := c.Init()
    if err != nil {
        return nil, err
    }
    return c, nil
}

func (m *ApiClient) Init() error {

	apiClient, err := api.NewClient(utils.ApiBaseUrl())
	if err != nil {
		slog.Error("failed to create api client", err)
        return err
	}

	m.apiClient = apiClient
    return nil
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

    return result, nil
}
