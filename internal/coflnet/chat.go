package coflnet

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"

	"github.com/Coflnet/coflnet-bot/codegen/chat"
	"github.com/Coflnet/coflnet-bot/internal/utils"
)

const (
	chatApiTracerName = "chat-api-client"
)

func NewChatApi() *ChatApi {
	instance := &ChatApi{}
	instance.tracer = otel.Tracer(chatApiTracerName)

	chatBaseUrl, err := utils.ChatBaseUrl()
	if err != nil {
		slog.Error("error getting chat api url", err)
		return instance
	}

	instance.apiClient, err = chat.NewClientWithResponses(chatBaseUrl)
	if err != nil {
		slog.Error("error creating chat api client", err)
	}

	return instance
}

type ChatApi struct {
	apiClient *chat.ClientWithResponses
	tracer    trace.Tracer
}

func (r *ChatApi) SendMessage(ctx context.Context, msg string, playerUUID string, coflDiscordClientName string, playerName string) error {
	ctx, span := r.tracer.Start(ctx, "send-message-to-chat-api")
	defer span.End()

	if r.apiClient == nil {
		return errors.New("chat api client not initialized")
	}

	params := chat.PostApiChatSendParams{
		Authorization: utils.StrPtr(utils.ChatApiKey()),
	}

	body := chat.PostApiChatSendJSONRequestBody{
		Message:    utils.StrPtr(msg),
		Uuid:       utils.StrPtr(playerUUID),
		ClientName: utils.StrPtr(coflDiscordClientName),
		Name:       utils.StrPtr(playerName),
	}

	slog.Debug("sending message to chat api", "user", playerUUID, "message", msg)
	response, err := r.apiClient.PostApiChatSendWithResponse(ctx, &params, body)

	if err != nil {
		slog.Error("error sending message to chat api", err)
		span.RecordError(err)
		return err
	}

	switch response.StatusCode() {
	case 200:
		slog.Info("message sent successfully", "user", playerUUID, "message", msg)
		span.SetAttributes(attribute.Bool("success", true))
		return nil
	case 400:
		err = errors.New("bad request")
		slog.Error("error sending message to chat api, bad request", "err", err, "user", playerUUID, "message", msg)
		span.RecordError(err)
		return err
	case 500:
		err = errors.New("internal server error")
		slog.Error("error sending message to chat api, internal server error", "err", err, "user", playerUUID, "message", msg)
		span.RecordError(err)
		return err
	default:
		err = errors.New(fmt.Sprintf("unknown response: %v %v", response.StatusCode(), response))
		slog.Error("error sending message to chat api", "err", err, "user", playerUUID, "message", msg)
		span.RecordError(err)
		return err
	}
}

func (a *ChatApi) MuteUser(ctx context.Context, uuidTarget, muter, message, reason string) (*chat.Mute, error) {
	ctx, span := a.tracer.Start(ctx, "mute-user")
	defer span.End()

	if a.apiClient == nil {
		return nil, errors.New("chat api client not initialized")
	}

	param := chat.PostApiChatMuteParams{
		Authorization: utils.StrPtr(utils.ChatApiKey()),
	}
	body := chat.PostApiChatMuteJSONRequestBody{
		Uuid:    utils.StrPtr(uuidTarget),
		Reason:  utils.StrPtr(reason),
		Message: utils.StrPtr(message),
		Muter:   utils.StrPtr(muter),
	}

	slog.Debug("send mute request to chat api")
	response, err := a.apiClient.PostApiChatMuteWithResponse(ctx, &param, body)

	if err != nil {
		slog.Error("error muting user", err)
		span.RecordError(err)
		return nil, err
	}

	switch response.StatusCode() {
	case 200:
		mute := response.JSON200
		slog.Info("user muted successfully", "user", uuidTarget, "muter", muter, "reason", reason, "duration", mute.Expires)
		span.SetAttributes(attribute.Bool("success", true))
		return mute, nil
	case 400:
		err = errors.New("bad request")
		slog.Error("error muting user, bad request", "message", response.JSON400.Message, "slug", response.JSON400.Slug, "user", uuidTarget, "muter", muter, "reason", reason)
		span.RecordError(err)
		span.SetAttributes(attribute.String("message", *response.JSON400.Message))
		span.SetAttributes(attribute.String("slug", *response.JSON400.Slug))
		return nil, err
	case 500:
		err = errors.New("internal server error")
		slog.Error("error muting user, internal server error", "message", response.JSON500.Message, "slug", response.JSON500.Slug, "user", uuidTarget, "muter", muter, "reason", reason)
		span.RecordError(err)
		span.SetAttributes(attribute.String("message", *response.JSON400.Message))
		span.SetAttributes(attribute.String("slug", *response.JSON400.Slug))
		return nil, err
	default:
		err = errors.New(fmt.Sprintf("unknown response: %v %v", response.StatusCode(), response))
		slog.Error("error muting user, unknown status code", err, "user", uuidTarget, "muter", muter, "reason", reason, "code", response.StatusCode())
		span.RecordError(err)
		return nil, err
	}
}

func (a *ChatApi) UnmuteUser(ctx context.Context) (*chat.UnMute, error) {
	return nil, fmt.Errorf("unmute not implemented")
}
