package coflnet

import (
	"context"
	"errors"
	"fmt"

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
	tracer    trace.Tracer
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

func (r *ChatApi) SendMessage(ctx context.Context, msg *chat.APIChatSendPostTextJSON) (*chat.ChatMessage, error) {
	ctx, span := r.tracer.Start(ctx, "send-message-to-chat-api")
	defer span.End()

	if r.apiClient == nil {
		return nil, errors.New("chat api client not initialized")
	}

	response, err := r.apiClient.APIChatSendPost(ctx, msg)

	if err != nil {
		slog.Error("error sending message to chat api", err)
		span.RecordError(err)
		return nil, err
	}

	switch r := response.(type) {
	case *chat.APIChatSendPostApplicationJSONOK:
		msg := chat.ChatMessage(*r)
		slog.Info("message sent successfully sent to chat api")
		span.SetAttributes(attribute.Bool("success", true))
		return &msg, nil
	case *chat.APIChatSendPostTextJSONInternalServerError:
		err := errors.New("internal server error")
		slog.Error("error sending message to chat api, internal server error", err)
		span.RecordError(err)
		return nil, err
	case *chat.APIChatSendPostTextJSONBadRequest:
		err := errors.New("bad request")
		slog.Error(fmt.Sprintf("message: %s", r.Message.Value), err)
		span.RecordError(err)
		return nil, err
	default:
		err = errors.New("unknown response")
		slog.Error("error sending message to chat api, unknown response", err)
		span.RecordError(err)
		return nil, err
	}
}

func (a *ChatApi) MuteUser(ctx context.Context, mute *chat.APIChatMutePostTextJSON) (*chat.Mute, error) {
	ctx, span := a.tracer.Start(ctx, "mute-user")
	defer span.End()

	if a.apiClient == nil {
		return nil, errors.New("chat api client not initialized")
	}

	slog.Info(fmt.Sprintf("sending mute request to chat api for user %s", mute.UUID.Value))
	response, err := a.apiClient.APIChatMutePost(ctx, mute)

	if err != nil {
		slog.Error("error muting user", err)
		span.RecordError(err)
		return nil, err
	}
	slog.Info("got response from chat api")

	switch r := response.(type) {
	case *chat.APIChatMutePostApplicationJSONOK:
		mute := chat.Mute(*r)
		slog.Info("user muted successfully until %s", mute.Expires.Value)
		span.SetAttributes(attribute.Bool("success", true))
		return &mute, nil
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
		err = errors.New(fmt.Sprintf("unknown response: %v %v", r, response))
		slog.Error("error muting user", err)
		span.RecordError(err)
		return nil, err
	}
}

func (a *ChatApi) UnmuteUser(ctx context.Context, unmute *chat.APIChatMuteDeleteTextJSON) (*chat.UnMute, error) {
	ctx, span := a.tracer.Start(ctx, "unmute-user")
	defer span.End()

	if a.apiClient == nil {
		return nil, errors.New("chat api client not initialized")
	}

	slog.Info(fmt.Sprintf("sending unmute request to chat API for user %s", unmute.UUID.Value))
	response, err := a.apiClient.APIChatMuteDelete(ctx, unmute)

	if err != nil {
		slog.Error("error unmuting user", err)
		span.RecordError(err)
		return nil, err
	}

	switch r := response.(type) {
	case *chat.APIChatMuteDeleteApplicationJSONOK:
		unmute := chat.UnMute(*r)
		slog.Info("user unmuted successfully")
		span.SetAttributes(attribute.Bool("success", true))
		return &unmute, nil
	case *chat.APIChatMuteDeleteApplicationJSONBadRequest:
		err := errors.New("bad request")
		slog.Error("error unmuting user, bad request", err)
		span.RecordError(err)
		return nil, err
	case *chat.APIChatMuteDeleteApplicationJSONInternalServerError:
		err := errors.New("internal server error")
		slog.Error("error unmuting user, internal server error", err)
		span.RecordError(err)
		return nil, err
	default:
		err = errors.New(fmt.Sprintf("unknown response: %v %v", r, response))
		slog.Error("error unmuting user", err)
		span.RecordError(err)
		return nil, err
	}
}
