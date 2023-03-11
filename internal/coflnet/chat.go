package coflnet

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// "github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/chat"
)

const (
    chatApiTracerName = "chat-api-client"
)

var (
    instance *ChatApi
)

type ChatApi struct {
    apiClient *chat.Client
    tracer trace.Tracer
}

func NewChatClient() (*ChatApi, error) {

    if instance != nil {
        return instance, nil
    }

    var err error
    r := &ChatApi{}
    r.apiClient, err = chat.NewClient(utils.ChatBaseUrl())
    r.tracer = otel.Tracer(chatApiTracerName)

    return instance, err
}

func (r *ChatApi) SendMessage(ctx context.Context, msg *chat.APIChatSendPostReqApplicationJSON) error {
    ctx, span := r.tracer.Start(ctx, "send-message-to-chat-api")
    defer span.End()


    if r.apiClient == nil {
        return errors.New("chat api client not initialized")
    }

    response, err := r.apiClient.APIChatSendPost(ctx, msg)

    if err != nil {
        span.RecordError(err)
        return err
    }

    if response.UUID.IsSet() {
        span.SetAttributes(attribute.String("response-sender", response.UUID.Value))
    }
    return nil
}

func (a *ChatApi) MuteUser(ctx context.Context, mute *chat.APIChatMutePostReqApplicationJSON) (*chat.APIChatMutePostOKApplicationJSON, error) {
    ctx, span := a.tracer.Start(ctx, "mute-user")
    defer span.End()

    if a.apiClient == nil {
        return nil, errors.New("chat api client not initialized")
    }

    response, err := a.apiClient.APIChatMutePost(ctx, mute)

    if err != nil {
        span.RecordError(err)
        return nil, err
    }

    return response, err
}
