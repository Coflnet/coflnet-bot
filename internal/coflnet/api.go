package coflnet

import (
	"context"
	"errors"
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/api"
	"go.opentelemetry.io/otel"
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

func (m *ApiClient) SearchUUIDForPlayer(ctx context.Context, username string) (string, error) {
    ctx, span := m.tracer.Start(ctx, "get-server-info")
    defer span.End()

    playerResults, err := m.apiClient.APISearchPlayerPlayerNameGet(ctx, api.APISearchPlayerPlayerNameGetParams{
        PlayerName: username,
    })

    if err != nil {
        slog.Error("failed to get player info", err)
        span.RecordError(err)
        return "", err
    }

    if len(playerResults) == 0 {
        slog.Warn("no player found")
        return "", errors.New(fmt.Sprintf("no player with name %s found"))
    }

    if len(playerResults) > 1 {
        slog.Warn("multiple players found")
        return "", errors.New(fmt.Sprintf("multiple players with name %s found"))
    }

    if playerResults[0].UUID.IsNull() {
        slog.Warn("player has no UUID")
        return "", errors.New(fmt.Sprintf("player %s has no UUID", username))
    }

    return playerResults[0].UUID.Value, nil
}
