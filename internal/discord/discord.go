package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sync"

	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/pkg/discord"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
    discordHandlerTracerName = "discord-handler"
)

var (
    instance *DiscordHandler
    initMutex sync.Mutex
)

type Discord interface {
    SendMessage(msg *DiscordMessage) error
}

type DiscordHandler struct {
    session *discordgo.Session
    commands []*discordgo.ApplicationCommand


    // internal data structure that receives messages from discord
    messagesReceived chan discordgo.Message

    tracer trace.Tracer
}

type DiscordMessage struct {
	Message string
	Channel string
	Webhook string
}

type WebhookRequest struct {
	Content             string          `json:"content"`
	Username            string          `json:"username"`
	AvatarUrl           string          `json:"avatar_url"`
	AllowedMentionsData AllowedMentions `json:"allowed_mentions"`
}

type AllowedMentions struct {
	Parse []string `json:"parse"`
}

func NewDiscordHandler() (*DiscordHandler, error) {

    if instance != nil {
        return instance, nil
    }

    initMutex.Lock()
    slog.Info("creating new discord handler")


    instance := &DiscordHandler{
        tracer: otel.Tracer(discordHandlerTracerName),
    }

    err := instance.initSession()
    if err != nil {
        return nil, err
    }

    go instance.RegisterCommands()

    return instance, nil
}

func (d *DiscordHandler) Close() {
    d.unregisterCommands()
    d.session.Close()
}

func (d *DiscordHandler) initSession() error {
	var err error
	d.session, err = discordgo.New("Bot " + utils.DiscordBotToken())
	if err != nil {
        return err
	}

	d.session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

    go func() {
        d.session.Open()
        defer d.session.Close()

        sig := make(chan os.Signal, 1)
        signal.Notify(sig, os.Interrupt)
        <-sig

        slog.Info("closing discord session")
    }()

    return nil
}


// SendMessage sends a discord message to the discord server
func (d *DiscordHandler) SendMessage(ctx context.Context, msg *DiscordMessage) error {
    _, span := d.tracer.Start(ctx, "send-discord-message")
    defer span.End()
    slog.Debug("sending a discord message")

    if msg.Message == "" {
		return fmt.Errorf("no message is set")
	}

	if msg.Webhook == "" {
		if msg.Channel != "" {
            var err error
            msg.Webhook, err = d.webhookForChannel(msg.Channel)
            if err != nil {
                slog.Error("error when getting webhook for channel", err)
                span.RecordError(err)
                return err
            }
		}
	}

	if msg.Webhook == "" {
		return fmt.Errorf("no webhook is set, webhook: %s, channel: %s", msg.Webhook, msg.Channel)
	}

	data := &WebhookRequest{
		Content:             msg.Message,
		Username:            "cofl-bot",
		AvatarUrl:           "https://cdn.discordapp.com/app-icons/888725077191974913/0c0e3b97e6865091ef14162083a54a42.png?size=256",
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	body, err := json.Marshal(data)
	if err != nil {
        slog.Error("error when marshalling webhook request", err)
        span.RecordError(err)
		return err
	}

	url := msg.Webhook
	_, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
        slog.Error("error when sending webhook request", err)
        span.RecordError(err)
		return err
	}

	return nil
}

// returns a channel which contains all messages received from discord, or error if such a channel already exists
func (d *DiscordHandler) DiscordMessagesSent(ctx context.Context) (<-chan discordgo.Message, error) {
    if d.messagesReceived == nil {
        d.messagesReceived = make(chan discordgo.Message, 100)
    } else {
        return nil, errors.New("messages received channel already exists")
    }

    slog.Info("adding a discord handler for received messages")
    d.session.AddHandler(d.messageReceived)

    return d.messagesReceived, nil
}

func (d *DiscordHandler) messageReceived(_ *discordgo.Session, m *discordgo.MessageCreate) {
    d.messagesReceived <- *m.Message
}

func (d *DiscordHandler) webhookForChannel(channel string) (string, error) {
	switch channel {
	case discord.WarningsChannel:
		return utils.WarningsWebhook(), nil
	case discord.MutesChannel:
		return utils.MutesWebhook(), nil
	case discord.CiSuccessChannel:
		return utils.CiSuccessWebhook(), nil
	case discord.CiFailureChannel:
		return utils.CiFailureWebhook(), nil
	case discord.FlipperRoleChannel:
		return utils.FlipperRoleWebhook(), nil
	case discord.FeedbackChannel:
		return utils.FeedbackWebhook(), nil
	}

	return "", errors.New("no webhook found for channel")
}

func (d *DiscordHandler) SendMessageToIngameChat(ctx context.Context, message *mongo.ChatMessage) error {
    _, span := d.tracer.Start(ctx, "send-discord-message-to-ingame-chat")
    defer span.End()

	if message.UUID == "" {
		return fmt.Errorf("no icon url is set")
	}

	iconUrl := fmt.Sprintf("https://crafatar.com/avatars/%s", message.UUID)
	url := os.Getenv("CHAT_WEBHOOK")

	msg := message.Message
	data := &WebhookRequest{
		Content:             d.sanitizeMessage(msg),
		Username:            message.Name,
		AvatarUrl:           iconUrl,
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	body, err := json.Marshal(data)
	if err != nil {
        slog.Error("error when marshalling webhook request", err)
        span.RecordError(err)
	}

	_, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
        slog.Error("error when sending webhook request", err)
        span.RecordError(err)
		return err
	}

	return nil
}

func (d *DiscordHandler) sanitizeMessage(message string) string {
	// if the strings has a § in it, remove all § characters and the following character

	slog.Debug("sanitizing message: %s", message)
	reg := regexp.MustCompile("§.")
	message = reg.ReplaceAllString(message, "")
    slog.Debug("sanitized message: %s", message)

	return message
}

func (d *DiscordHandler) AnswerDiscordMessage(msg string, originalMessage *discordgo.Message) (*discordgo.Message, error) {
    return d.session.ChannelMessageSendReply(originalMessage.ChannelID, msg, originalMessage.Reference())
}