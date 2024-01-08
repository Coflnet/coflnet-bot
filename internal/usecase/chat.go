package usecase

import (
	"coflnet-bot/internal/db"
	chatgen "coflnet-bot/internal/gen/chat"
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"sync"
	"time"
)

const (
	coflDiscordClientName = "§1dc§7:§f"
)

type Chat struct {
	tracer          trace.Tracer
	discordMessages *DiscordMessageService
	redisMessages   *RedisMessageService
	chatClient      *chatgen.Client
	userService     *UserService

	redisMessageProcessed         metric.Int64Counter
	redisMessageProcessedFailed   metric.Int64Counter
	discordMessageProcessed       metric.Int64Counter
	discordMessageProcessedFailed metric.Int64Counter
}

func NewChat(discordMessageService *DiscordMessageService, redisMessageService *RedisMessageService, chatClient *chatgen.Client, userService *UserService) (*Chat, error) {
	meter := otel.Meter("chat-service")
	redisMessageForwarded, err := meter.Int64Counter("coflnet_bot_redis_message_processed", metric.WithDescription("Number of redis messages forwarded to chat api"), metric.WithUnit("1"))
	if err != nil {
		return nil, err
	}
	redisMessageForwardedFailed, err := meter.Int64Counter("coflnet_bot_redis_message_processed_failed", metric.WithDescription("Number of redis messages forwarded to chat api failed"), metric.WithUnit("1"))
	if err != nil {
		return nil, err
	}
	discordMessageForwarded, err := meter.Int64Counter("coflnet_bot_discord_message_processed", metric.WithDescription("Number of discord messages forwarded to chat api"), metric.WithUnit("1"))
	if err != nil {
		return nil, err
	}
	discordMessageForwardedFailed, err := meter.Int64Counter("coflnet_bot_discord_message_processed_failed", metric.WithDescription("Number of discord messages forwarded to chat api failed"), metric.WithUnit("1"))
	if err != nil {
		return nil, err
	}

	return &Chat{
		tracer:          otel.Tracer("chat-service"),
		discordMessages: discordMessageService,
		redisMessages:   redisMessageService,
		chatClient:      chatClient,
		userService:     userService,

		redisMessageProcessed:         redisMessageForwarded,
		redisMessageProcessedFailed:   redisMessageForwardedFailed,
		discordMessageProcessed:       discordMessageForwarded,
		discordMessageProcessedFailed: discordMessageForwardedFailed,
	}, nil
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
			err := c.processDiscordMessage(msgCtx, msg)
			if err != nil {
				slog.Error("unable to process discord message", "err", err)
				c.discordMessageProcessedFailed.Add(msgCtx, 1)
				return
			}
			c.discordMessageProcessed.Add(msgCtx, 1)
		}(msg)
	}
}

func (c *Chat) StartRedisMessageListener(ctx context.Context, reader <-chan *RedisMessage) {
	slog.Info("Start listening for redis messages")
	for msg := range reader {
		go func(msg *RedisMessage) {
			msgCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			err := c.processRedisMessage(msgCtx, msg)
			if err != nil {
				slog.Error("unable to process redis message", "err", err)
				c.redisMessageProcessedFailed.Add(msgCtx, 1)
				return
			}
			c.redisMessageProcessed.Add(msgCtx, 1)
		}(msg)
	}
}

func (c *Chat) processDiscordMessage(ctx context.Context, msg *discordgo.Message) error {
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
		return err
	}
	c.discordMessageProcessedFailed.Add(ctx, 1)

	// send the message to the chat api
	if !c.shouldMessageBeForwardedToChatAPI(msg) {
		span.SetAttributes(attribute.Bool("should-forward-to-chat-api", false))
		return nil
	}

	span.SetAttributes(attribute.Bool("should-forward-to-chat-api", true))

	// search the user by the discord id
	user, err := db.UserByDiscordId(ctx, msg.Author.ID)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to load user by discord id", "err", err)
		innerErr := AnswerDiscordMessage(ctx, msg, "no username for your discord account found\nYou can try to write a message in minecraft with your connected minecraft account\nIt is important that your hypixel account has a link to your discord account\nIf this does not work ask in <@669976123714699284>")
		if innerErr != nil {
			span.RecordError(innerErr)
			slog.Error("unable to answer discord message", "err", innerErr)
		}
		return err
	}

	var uuid string
	uuid, err = user.PreferredUUID()
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to get preferred uuid", "err", err)
		innerErr := AnswerDiscordMessage(ctx, msg, "no username for your discord account found\nYou can try to write a message in minecraft with your connected minecraft account\nIt is important that your hypixel account has a link to your discord account\nIf this does not work ask in <@669976123714699284>")
		if innerErr != nil {
			span.RecordError(innerErr)
			slog.Error("unable to answer discord message", "err", innerErr)
		}
		return err
	}

	err = c.sendMessageToChatAPI(ctx, chatMessage, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to send message to chat api", "err", err)
		innerErr := AnswerDiscordMessage(ctx, msg, "There was a internal issue when forwarding your request, if this continues to happen please open a issue in <@884002032392998942>")
		if innerErr != nil {
			span.RecordError(innerErr)
			slog.Error("unable to answer discord message", "err", innerErr)
		}
		return err
	}

	slog.Info("message forwarded to chat api")
	return nil
}

func (c *Chat) sendMessageToChatAPI(ctx context.Context, message *db.Message, uuid string) error {
	ctx, span := c.tracer.Start(ctx, "send-message-to-chat-api")
	defer span.End()

	body := chatgen.PostApiChatSendJSONRequestBody{
		ClientName: strPtr("cofl-discord"),
		Message:    strPtr(message.Content),
		Prefix:     strPtr(coflDiscordClientName),
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

func (c *Chat) processRedisMessage(ctx context.Context, msg *RedisMessage) error {
	ctx, span := c.tracer.Start(ctx, "process-redis-message")
	defer span.End()

	uuid := msg.UUID
	user, err := c.userService.LoadUserByUUID(ctx, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to load user", "err", err)
		return err
	}
	slog.Info(fmt.Sprintf("loaded user with uuid %s, external id %s", uuid, user.ExternalId))

	if msg.Prefix == coflDiscordClientName {
		slog.Info("message is already from discord, skip")
		return nil
	}

	// send the message to the discord channel
	err = c.sendMessageToIngameChat(ctx, msg)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to send message to ingame chat", "err", err)
		return err
	}
	return nil
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
	res := &db.Message{
		Content:         msg.Content,
		Timestamp:       msg.Timestamp,
		DiscordUserId:   msg.Author.ID,
		DiscordUsername: msg.Author.Username,
		IsBot:           msg.Author.Bot,
		ChannelId:       msg.ChannelID,
		GuildId:         msg.GuildID,
	}

	if msg.Thread != nil {
		res.ThreadId = &msg.Thread.ID
	}

	return res
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
