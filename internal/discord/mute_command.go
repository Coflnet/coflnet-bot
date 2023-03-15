package discord

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/chat"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

type MuteCommand struct {
    baseCommand *baseCommand
    tracer trace.Tracer
    chatApi *coflnet.ChatApi
    clientApi *coflnet.ApiClient
}

func CreateMuteCommand() *MuteCommand {
    c := &MuteCommand{
        tracer: otel.Tracer("mute-command"),
    }

    c.Init()

    return c
}

func (m *MuteCommand) Name() string {
    return "in-game-mute"
}

func (m *MuteCommand) Init() {
	m.baseCommand = newBaseCommand()

    chatApi, err := coflnet.NewChatClient()
    if err != nil {
        slog.Error("failed to create chat client", err)
        panic(err)
    }

    apiClient, err := coflnet.NewApiClient()
    if err != nil {
        slog.Error("failed to create api client", err)
        panic(err)
    }

    m.chatApi = chatApi
    m.clientApi = apiClient
}

func (m *MuteCommand) CreateCommand() *discordgo.ApplicationCommand {
    m.Init()

    return &discordgo.ApplicationCommand{
        Name:        m.Name(),
        Description: "Mutes a user in the ingame chat",
        Options: []*discordgo.ApplicationCommandOption{
            {
                Name:        "user",
                Description: "The user to mute",
                Type:        discordgo.ApplicationCommandOptionString,
                Required:    true,
            },
            {
                Name:        "reason",
                Description: "Internal Reason",
                Type:        discordgo.ApplicationCommandOptionString,
                Required:    true,
            },
            {
                Name:        "message",
                Description: "Message for the user",
                Type:        discordgo.ApplicationCommandOptionString,
                Required:    true,
            },
        },
    }
}

func (m *MuteCommand) HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    ctx, span := m.tracer.Start(ctx, "handle-mute-command")
    defer span.End()

	// first response
    if err := m.baseCommand.requestReceived(ctx, s, i); err != nil {
        slog.Error("sending first response failed", err)
        span.RecordError(err)
        return
    }

	// respond to command
	msg, err := m.baseCommand.createFollowupMessage(ctx, "⏳ mute in progress", s, i)
	if err != nil {
		slog.Error("failed to create followup message", err)
		span.RecordError(err)
		return
	}


    // Access options in the order provided by the user.
	optionMap := m.baseCommand.parseResponseOptions(i)

    // get the reason
    reason := optionMap["reason"].StringValue()

    // get the message
    message := optionMap["message"].StringValue()

    // the the muter
    muter := i.Member.User.Username

    // get the user to mute
    user := optionMap["user"].StringValue()

    // check if the user is at least mod 
    if !utils.IsUserHelper(i.Member.Roles) && !utils.IsUserMod(i.Member.Roles) {
        err := errors.New(fmt.Sprintf("User %s is not a mod", i.Member.User.Username))
        slog.Warn("failed to mute user", err)
        span.RecordError(err)
        if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("❌ failed to mute user %s; you are not authorized; error: %s", user, span.SpanContext().TraceID()), msg.ID, s, i); err != nil {
            slog.Error("failed to edit followup message", err)
            span.RecordError(err)
        }
        return
    }

    // search the uuid for the mc username
    userUUID, err := m.clientApi.SearchUUIDForPlayer(ctx, user)
    if err != nil {
        slog.Error("mc uuid not found", err)
        span.RecordError(err)
        errMsg := fmt.Sprintf("❌ failed to mute %s, mc uuid for %s not found; error: %s", user, user, span.SpanContext().TraceID())
        if _, err := m.baseCommand.editFollowupMessage(ctx, errMsg, msg.ID, s, i); err != nil {
            slog.Error("failed to edit followup message", err)
            span.RecordError(err)
        }
        return
    }

    slog.Info(fmt.Sprintf("muting %s for %s; Muter: %s", user, reason, muter))
    mute, err := m.chatApi.MuteUser(ctx, &chat.APIChatMutePostTextJSON{
        Muter: chat.NewOptNilString(muter),
        Reason: chat.NewOptNilString(reason),
        Message: chat.NewOptNilString(message),
        UUID: chat.NewOptNilString(userUUID),
    })

    if err != nil {
        slog.Error("failed to mute user", err)
        span.RecordError(err)
        errMsg := fmt.Sprintf("❌ failed to mute %s; error: id: %s, text: %s", user, err.Error(), span.SpanContext().TraceID())
        if _, err := m.baseCommand.editFollowupMessage(ctx, errMsg, msg.ID, s, i); err != nil {
            slog.Error("failed to edit followup message", err)
            span.RecordError(err)
        }
        return
    }

	slog.Info("update the followup message")
    if _, err := m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("✅ muted %s until %s", user, mute.Expires.Value), msg.ID, s, i); err != nil {
        slog.Error("failed to edit follup message", err)
        span.RecordError(err)
    }
}


