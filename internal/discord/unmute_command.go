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

type UnmuteCommand struct {
    baseCommand *baseCommand
    tracer trace.Tracer
    chatApi *coflnet.ChatApi
    clientApi *coflnet.ApiClient
}

func CreateUnmuteCommand() *UnmuteCommand {
    c := &UnmuteCommand{
        tracer: otel.Tracer("unmute-command"),
    }

    c.Init()

    return c
}

func (m *UnmuteCommand) Name() string {
    return "in-game-unmute"
}

func (m *UnmuteCommand) Init() {

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

func (m *UnmuteCommand) CreateCommand() *discordgo.ApplicationCommand {
    m.Init()

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
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    ctx, span := m.tracer.Start(ctx, "handle-unmute-command")
    defer span.End()

    // check if the user is at least mod 
    if !utils.IsUserMod(i.Member.Roles) && !utils.IsUserHelper(i.Member.Roles) {
        err := errors.New(fmt.Sprintf("User %s is not a mod or a helper", i.Member.User.Username))
        slog.Warn("failed to unmute user", err)
        span.RecordError(err)

        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ you do not have permissions to unmute players, this incident will be reported",
                AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
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

    // get the user to mute
    user := optionMap["user"].StringValue()

    // search the uuid for the mc username
    userUUID, err := m.clientApi.SearchUUIDForPlayer(ctx, user)
    if err != nil {
        slog.Error("mc uuid not found", err)
        span.RecordError(err)
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: fmt.Sprintf("❌ no uuid found for the user %s, check spelling. error: %s", user, span.SpanContext().TraceID()),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })
        return
    }


    slog.Info(fmt.Sprintf("unmuting %s for %s; Muter: %s", user, reason, i.Member.User.Username))
    _, err = m.chatApi.UnmuteUser(ctx, &chat.APIChatMuteDeleteTextJSON{
        UUID: chat.OptNilString{
            Value: userUUID,
            Set: true,
        },
        Reason: chat.OptNilString{
            Value: reason,
            Set: true,
        }, 
    })

    if err != nil {
        slog.Error("failed to mute user", err)
        span.RecordError(err)
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseDeferredMessageUpdate,
            Data: &discordgo.InteractionResponseData{
                Content: fmt.Sprintf("❌ failed to unmute %s; error: id: %s, text: %s", user, err.Error(), span.SpanContext().TraceID()),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })
        return
    }

    slog.Info("update the followup message")    
    m.baseCommand.editFollowupMessage(ctx, fmt.Sprintf("✅ %s unmuted", user), msg.ID, s, i)
}

