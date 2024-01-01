package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

var (
	client        *discordgo.Session
	discordTracer = otel.Tracer("discord")
)

func DiscordSession(ctx context.Context) (*discordgo.Session, error) {
	if client != nil {
		return client, nil
	}
	return StartSession(ctx)
}

func StartSession(ctx context.Context) (*discordgo.Session, error) {
	ctx, span := discordTracer.Start(ctx, "start-session")
	defer span.End()
	slog.Info("Starting a new discord session")

	token, err := BotToken()
	if err != nil {
		return nil, err
	}

	client, err = discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}

	// In this example, we only care about receiving message events.
	client.Identify.Intents = discordgo.IntentsGuildMessages

	err = client.Open()
	if err != nil {
		return nil, err
	}
	go startStopListener()

	return client, nil
}

func startStopListener() {

	slog.Info("Bot is now running, until the application stops")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	slog.Info("Shutting down discord session")
}

func BotToken() (string, error) {
	token, found := os.LookupEnv("BOT_TOKEN")
	if !found {
		slog.Warn("BOT_TOKEN not found")
		return "", errors.New("BOT_TOKEN not found")
	}

	slog.Debug("BOT_TOKEN found, returning")
	return token, nil
}

func guildId() string {
	id, found := os.LookupEnv("GUILD_ID")
	if !found {
		panic("GUILD_ID not found")
	}

	slog.Debug("GUILD_ID found, returning")
	return id
}

func SearchDiscordUser(ctx context.Context, name string) (*discordgo.Member, error) {
	ctx, span := discordTracer.Start(ctx, "search-discord-user")
	defer span.End()
	span.SetAttributes(attribute.String("search-term", name))

	members, err := client.GuildMembersSearch(guildId(), name, 1)
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return nil, &DiscordMemberNotFoundError{SearchTerm: name}
	}
	if len(members) > 1 {
		slog.Warn("found more than one discord member, choosing first one")
	}

	return members[0], nil
}

func SendMessageToIngameChat(ctx context.Context, content, uuid, username string) error {
	ctx, span := discordTracer.Start(ctx, "send-discord-message-to-ingame-chat")
	defer span.End()

	iconUrl := fmt.Sprintf("https://mc-heads.net/avatar/%s", uuid)
	url := chatWebhook()
	span.SetAttributes(attribute.String("iconUr", iconUrl))
	span.SetAttributes(attribute.String("url", url))

	data := &WebhookRequest{
		Content:             sanitizeMessage(content),
		Username:            username,
		AvatarUrl:           iconUrl,
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	slog.Debug(fmt.Sprintf("send message to discord webhook with icon url %s", iconUrl))
	body, err := json.Marshal(data)
	if err != nil {
		slog.Error("error when marshalling webhook request", err)
		span.RecordError(err)
		return err
	}

	response, err := http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))
	span.SetAttributes(attribute.Int("status", response.StatusCode))

	if err != nil {
		slog.Error("error when sending webhook request", err)
		span.RecordError(err)
		return err
	}

	return nil
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

func sanitizeMessage(message string) string {
	reg := regexp.MustCompile("§.")
	message = reg.ReplaceAllString(message, "")
	return message
}

func chatWebhook() string {
	val, found := os.LookupEnv("CHAT_WEBHOOK")
	if !found {
		panic("CHAT_WEBHOOK not found")
	}
	return val
}

type DiscordMemberNotFoundError struct {
	SearchTerm string
}

func (e *DiscordMemberNotFoundError) Error() string {
	return fmt.Sprintf("discord member with search term %s not found", e.SearchTerm)
}
