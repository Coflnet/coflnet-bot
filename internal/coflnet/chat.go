package coflnet

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"

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
    instance := &ChatApi{}
    instance.apiClient, err = chat.NewClient(utils.ChatBaseUrl())
    instance.tracer = otel.Tracer(chatApiTracerName)

    return instance, err
}

func (r *ChatApi) SendMessage(ctx context.Context, msg *chat.ChatMessage) error {
    ctx, span := r.tracer.Start(ctx, "send-message-to-chat-api")
    defer span.End()

    if msg.Message.IsSet() {
        span.SetAttributes(attribute.String("message", msg.Message.Value))
    }
    if msg.Name.IsSet() {
        span.SetAttributes(attribute.String("sender", msg.Name.Value))
    }

    if r.apiClient == nil {
        return errors.New("chat api client not initialized")
    }

    response, err := r.apiClient.APIChatSendPost(ctx, msg)

    if err != nil {
        slog.Error("error sending message to chat api", err)
        span.RecordError(err)
        return err
    }

    switch response.(type) {
    case *chat.ChatMessage:
        slog.Info("message sent successfully sent to chat api")
        span.SetAttributes(attribute.Bool("success", true))
        return nil
    case *chat.APIChatSendPostApplicationJSONBadRequest:
        err := errors.New("bad request")
        slog.Error("error sending message to chat api, bad request", err)
        span.RecordError(err)
        return err
    case *chat.APIChatSendPostApplicationJSONInternalServerError:
        err := errors.New("internal server error")
        slog.Error("error sending message to chat api, internal server error", err)
        span.RecordError(err)
        return err
    default:
        err = errors.New("unknown response")
        slog.Error("error sending message to chat api, unknown response", err)
        span.RecordError(err)
        return err
    }
}

func (a *ChatApi) MuteUser(ctx context.Context, mute *chat.Mute) (*chat.Mute, error) {
    ctx, span := a.tracer.Start(ctx, "mute-user")
    defer span.End()

    if a.apiClient == nil {
        return nil, errors.New("chat api client not initialized")
    }

    response, err := a.apiClient.APIChatMutePost(ctx, mute)

    if err != nil {
        slog.Error("error muting user", err)
        span.RecordError(err)
        return nil, err
    }

    switch mute := response.(type) {
    case *chat.Mute:
        slog.Info("user muted successfully")
        span.SetAttributes(attribute.Bool("success", true))
        return mute, nil
    case *chat.APIChatMutePostApplicationJSONBadRequest:
        err := errors.New("bad request")
        slog.Error("error muting user, bad request", err)
        span.RecordError(err)
        return nil, err
    case *chat.APIChatMutePostApplicationJSONInternalServerError:
        err := errors.New("internal server error")
        slog.Error("error muting user, internal server error", err)
        span.RecordError(err)
        return nil, err
    default:
        err = errors.New("unknown response")
        slog.Error("error muting user, unknown response", err)
        span.RecordError(err)
        return nil, err
    }
}
