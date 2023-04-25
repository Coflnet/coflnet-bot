package discord

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/chat"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

func NewUnmuteCommand(chat *coflnet.ChatApi, api *coflnet.ApiClient) *UnmuteCommand {
	return &UnmuteCommand{
		tracer:      otel.Tracer("unmute-command"),
		chatApi:     chat,
		clientApi:   api,
		baseCommand: newBaseCommand(),
	}
}

type UnmuteCommand struct {
	baseCommand *baseCommand
	tracer      trace.Tracer
	chatApi     *coflnet.ChatApi
	clientApi   *coflnet.ApiClient
}

func (m *UnmuteCommand) Name() string {
	return "in-game-unmute"
}

func (m *UnmuteCommand) CreateCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        m.Name(),
		Description: "Unmutes a user in the ingame chat",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "The user to unmute",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "reason",
				Description: "Reason for the unmute",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}
}

func (m *UnmuteCommand) HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ctx, span := m.tracer.Start(ctx, "handle-unmute-command")
	defer span.End()

	// first fake response
	if err := m.baseCommand.requestReceived(ctx, s, i); err != nil {
		slog.Error("failed sending request received message", err)
		span.RecordError(err)
		return
	}

	// respond to command
	msg, err := m.baseCommand.createFollowupMessage(ctx, "⏳ unmute in progress", s, i)
	if err != nil {
		slog.Error("failed to create followup message", err)
		span.RecordError(err)
		return
	}

	// parse the strange discord map to a normal map
	optionMap := m.baseCommand.parseResponseOptions(i)

	// get the reason
	reason := optionMap["reason"].StringValue()
	span.SetAttributes(attribute.String("reason", reason))

	// get the user to mute
	user := optionMap["user"].StringValue()
	span.SetAttributes(attribute.String("user", user))

	// check if the user is at least mod
	span.SetAttributes(attribute.String("unmuter", i.Member.User.Username))
	if !utils.IsUserMod(i.Member.Roles) && !utils.IsUserHelper(i.Member.Roles) {
		err := errors.New(fmt.Sprintf("User %s is not a mod or a helper", i.Member.User.Username))
		slog.Warn("failed to unmute user", err)
		span.RecordError(err)
		if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("❌ failed to unmute user %s; you are not authorized; error: %s", user, span.SpanContext().TraceID()), msg.ID, s, i); err != nil {
			slog.Error("failed to edit followup message", err)
			span.RecordError(err)
		}

		span.SetAttributes(attribute.Bool("authorized", false))
		return
	}
	span.SetAttributes(attribute.Bool("authorized", true))

	// search the uuid for the mc username
	userUUIDs, err := m.clientApi.SearchUUIDForPlayer(ctx, user)
	if err != nil {
		slog.Error("mc uuid not found", err)
		span.RecordError(err)
		if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("❌ failed to unmute user %s; uuid for %s not found; error: %s", user, user, span.SpanContext().TraceID()), msg.ID, s, i); err != nil {
			slog.Error("failed to edit followup message", err)
			span.RecordError(err)
		}
		return
	}

	for _, userUUID := range userUUIDs {
		slog.Info(fmt.Sprintf("unmuting %s for %s; Muter: %s", user, reason, i.Member.User.Username))
		_, err = m.chatApi.UnmuteUser(ctx, &chat.APIChatMuteDeleteTextJSON{
			UUID: chat.OptNilString{
				Value: userUUID,
				Set:   true,
			},
			Reason: chat.OptNilString{
				Value: reason,
				Set:   true,
			},
		})

		if err != nil {
			slog.Error("failed to unmute user", err)
			span.RecordError(err)
			if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("❌ failed to unmute user %s; error: %s", user, span.SpanContext().TraceID()), msg.ID, s, i); err != nil {
				slog.Error("failed to edit followup message", err)
				span.RecordError(err)
			}
			return
		}

		slog.Info("update the followup message")
		if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("✅ %s unmuted", user), msg.ID, s, i); err != nil {
			slog.Error("failed to edit follup message", err)
			span.RecordError(err)
		}
	}

	span.SetAttributes(attribute.String("unmuted-uuids", strings.Join(userUUIDs, ", ")))
}
