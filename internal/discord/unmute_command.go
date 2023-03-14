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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

type UnmuteCommand struct {
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

    // Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

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
                Content: fmt.Sprintf("❌ failed to unmute %s, error: %s", user, span.SpanContext().TraceID()),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })
        return
    }

    go func() {
        _, span := m.tracer.Start(ctx, "respond-to-unmute-command")
        defer span.End()

        // set start time as attribute
        span.SetAttributes(attribute.Int64("start-respond", time.Now().Unix()))

        slog.Debug("sending response to unmute command")
        err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: fmt.Sprintf("trying to unmute %s", user),
                AllowedMentions: &discordgo.MessageAllowedMentions{},
            },
        })

        if err != nil {
            slog.Error("failed to respond to unmute command", err)
            span.RecordError(err)
            return
        }

        slog.Info("successfully responded to unmute command")
        span.SetAttributes(attribute.Int64("end-respond", time.Now().Unix()))
    }()


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

    slog.Info(fmt.Sprintf("unmute was sent successfully, trace: %s", span.SpanContext().TraceID()))
    err = discord.SendMessageToMutesChannel(fmt.Sprintf("%s unmuted %s", i.User.Username, user))
    if err != nil {
        slog.Error("failed to send message to mutes channel", err)
        span.RecordError(err)
        return
    }

    slog.Info("update response message")
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseUpdateMessage,
        Data: &discordgo.InteractionResponseData{
            Content: fmt.Sprintf("✅ unmuted %s", user),
            AllowedMentions: &discordgo.MessageAllowedMentions{},
        },
    })
}


