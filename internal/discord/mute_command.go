package discord

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/pkg/discord"
	"github.com/Coflnet/coflnet-bot/schemas/chat"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

type MuteCommand struct {
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
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    ctx, span := m.tracer.Start(ctx, "handle-mute-command")
    defer span.End()

    // check if the user is at least mod 
    if !utils.IsUserMod(i.Member.Roles) {
        err := errors.New(fmt.Sprintf("User %s is not a mod", i.Member.User.Username))
        slog.Warn("failed to mute user", err)
        span.RecordError(err)

        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ you do not have permissions to mute players, this incident will be reported",
                AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
        return
    }

    // Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

    // get the reason
    reason := optionMap["reason"].StringValue()

    // get the message
    message := optionMap["message"].StringValue()

    // the the muter
    muter := i.Member.User.Username

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
                Content: fmt.Sprintf("❌ failed to mute %s, error: %s", user, span.SpanContext().TraceID()),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })
        return
    }


    slog.Info(fmt.Sprintf("Muting %s for %s; Muter: %s", user, reason, muter))
    mute, err := m.chatApi.MuteUser(ctx, &chat.Mute{
        Muter: chat.OptNilString{
            Value: muter,
            Null: false,
        },
        Reason: chat.OptNilString{
            Value: reason,
            Null: false,
        },
        Message: chat.OptNilString{
            Value: message,
            Null: false,
        },
        UUID: chat.OptNilString{
            Value: userUUID,
            Null: false,
        },
    })

    if err != nil {
        slog.Error("failed to mute user", err)
        span.RecordError(err)
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: fmt.Sprintf("❌ failed to mute %s; error: id: %s, text: %s", user, err.Error(), span.SpanContext().TraceID()),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })
        return
    }

    err = discord.SendMessageToMutesChannel(fmt.Sprintf("%s muted %s", muter, user))
    if err != nil {
        slog.Error("failed to send message to mutes channel", err)
        span.RecordError(err)
        return
    }

    slog.Debug("sending response to mute command")
    err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: fmt.Sprintf("%s muted %s until %s", mute.Muter.Value, mute.UUID.Value, mute.Expires.Value),
            AllowedMentions: &discordgo.MessageAllowedMentions{},
        },
    })

    if err != nil {
        slog.Error("failed to respond to mute command", err)
        span.RecordError(err)
        return
    }
}


