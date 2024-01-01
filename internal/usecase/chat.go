package usecase

import (
	chatgen "coflnet-bot/gen/chat"
	"coflnet-bot/internal/db"
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"sync"
	"time"
)

const (
	coflDiscordClientName = "cofl-discord"
)

type Chat struct {
	tracer          trace.Tracer
	discordMessages *DiscordMessageService
	redisMessages   *RedisMessageService
	chatClient      *chatgen.Client
	userService     *UserService
}

func NewChat(discordMessageService *DiscordMessageService, redisMessageService *RedisMessageService, chatClient *chatgen.Client, userService *UserService) *Chat {
	return &Chat{
		tracer:          otel.Tracer("chat-service"),
		discordMessages: discordMessageService,
		redisMessages:   redisMessageService,
		chatClient:      chatClient,
		userService:     userService,
	}
}

func (c *Chat) StartChatService(ctx context.Context) error {
	slog.Info("Starting chat service")

	discordMessageReader, err := c.discordMessages.Reader(ctx)
	if err != nil {
		return err
	}

	redisMessageReader, err := c.redisMessages.Reader(ctx)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		c.StartDiscordMessageListener(ctx, discordMessageReader)
	}()

	go func() {
		defer wg.Done()
		c.StartRedisMessageListener(ctx, redisMessageReader)
	}()

	return nil
}

func (c *Chat) StartDiscordMessageListener(ctx context.Context, reader <-chan *discordgo.Message) {

	slog.Info("Start listening for discord messages")
	for msg := range reader {
		go func(msg *discordgo.Message) {
			msgCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			c.processDiscordMessage(msgCtx, msg)
		}(msg)
	}
}

func (c *Chat) StartRedisMessageListener(ctx context.Context, reader <-chan *RedisMessage) {
	slog.Info("Start listening for redis messages")
	for msg := range reader {
		go func(msg *RedisMessage) {
			msgCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			c.processRedisMessage(msgCtx, msg)
		}(msg)
	}
}

func (c *Chat) processDiscordMessage(ctx context.Context, msg *discordgo.Message) {
	ctx, span := c.tracer.Start(ctx, "process-discord-message")
	defer span.End()
	span.SetAttributes(
		attribute.String("discord-user-id", msg.Author.ID),
	)

	// convert into chat message
	chatMessage := c.convertDiscordMessageInChatMessage(msg)

	// save the message
	err := db.SaveMessage(ctx, chatMessage)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to save message", "err", err)
	}

	// send the message to the chat api
	if !c.shouldMessageBeForwardedToChatAPI(msg) {
		span.SetAttributes(attribute.Bool("should-forward-to-chat-api", false))
		return
	}

	span.SetAttributes(attribute.Bool("should-forward-to-chat-api", true))

	// search the user by the discord id
	user, err := db.UserByDiscordId(ctx, msg.Author.ID)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to load user by discord id", "err", err)

		// TODO tell the user that he is not registered
		return
	}

	var uuid string
	uuid, err = user.PreferredUUID()
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to get preferred uuid", "err", err)
		// TODO tell the user that he is not registered
		return
	}

	err = c.sendMessageToChatAPI(ctx, chatMessage, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to send message to chat api", "err", err)

		// TODO tell the user there was a problem
		return
	}

	slog.Info("message forwarded to chat api")
	return
}

func (c *Chat) sendMessageToChatAPI(ctx context.Context, message *db.Message, uuid string) error {
	ctx, span := c.tracer.Start(ctx, "send-message-to-chat-api")
	defer span.End()

	body := chatgen.PostApiChatSendJSONRequestBody{
		ClientName: strPtr("cofl-discord"),
		Message:    strPtr(message.Content),
		Name:       strPtr(coflDiscordClientName),
		Prefix:     strPtr(""),
		Uuid:       strPtr(uuid),
	}
	headers := &chatgen.PostApiChatSendParams{
		Authorization: strPtr(c.authorization()),
	}

	response, err := c.chatClient.PostApiChatSend(ctx, headers, body)
	if err != nil {
		return err
	}

	if response.StatusCode > 299 {
		return errors.New(fmt.Sprintf("unable to send message to chat api, status code is %d", response.StatusCode))
	}

	slog.Debug(fmt.Sprintf("message sent to chat api, status code is %d", response.StatusCode))
	return nil
}

func (c *Chat) shouldMessageBeForwardedToChatAPI(message *discordgo.Message) bool {
	if message.Author.Bot {
		return false
	}

	chatChannelID := c.chatChannelID()
	if message.ChannelID != chatChannelID {
		return false
	}

	return true
}

func (c *Chat) processRedisMessage(ctx context.Context, msg *RedisMessage) {
	ctx, span := c.tracer.Start(ctx, "process-redis-message")
	defer span.End()

	uuid := msg.UUID
	user, err := c.userService.LoadUserByUUID(ctx, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to load user", "err", err)
		return
	}
	slog.Info(fmt.Sprintf("loaded user with uuid %s, external id %s", uuid, user.ExternalId))

	// send the message to the discord channel
	err = c.sendMessageToIngameChat(ctx, msg)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to send message to ingame chat", "err", err)
		return
	}
}

func (c *Chat) sendMessageToIngameChat(ctx context.Context, msg *RedisMessage) error {
	ctx, span := c.tracer.Start(ctx, "send-message-to-ingame-chat")
	defer span.End()

	err := SendMessageToIngameChat(ctx, msg.Message, msg.UUID, msg.Name)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to send message to ingame chat", "err", err)
		return err
	}

	return nil
}

func (c *Chat) convertDiscordMessageInChatMessage(msg *discordgo.Message) *db.Message {
	return &db.Message{
		Content:         msg.Content,
		Timestamp:       msg.Timestamp,
		DiscordUserId:   msg.Author.ID,
		DiscordUsername: msg.Author.Username,
		IsBot:           msg.Author.Bot,
		ChannelId:       msg.ChannelID,
		GuildId:         msg.GuildID,
	}
}

func (c *Chat) chatChannelID() string {
	return mustEnv("CHAT_CHANNEL_ID")
}

func (c *Chat) authorization() string {
	return mustEnv("CHAT_API_AUTHORIZATION")
}

func mustEnv(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Sprintf("%s not found", key))
	}
	return val
}

func strPtr(s string) *string {
	return &s
}
