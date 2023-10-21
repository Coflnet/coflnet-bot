package processor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/Coflnet/coflnet-bot/internal/usecase"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	redisgo "github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
	coflDiscordClientName = "cofl-discord"
)

func NewChatProcessor(u *usecase.UserHandler, r *usecase.RedisHandler, d *discord.DiscordHandler, c *coflnet.ChatApi) *ChatProcessor {
	return &ChatProcessor{
		userHandler:       u,
		redisHandler:      r,
		discordHandler:    d,
		coflnetChatClient: c,
		tracer:            otel.Tracer("chat-processor"),
	}
}

// ChatProcessor handles all the chat messages from redis pubsub and discord and transfers them to the other platform
type ChatProcessor struct {
	redisHandler      *usecase.RedisHandler
	discordHandler    *discord.DiscordHandler
	userHandler       *usecase.UserHandler
	coflnetChatClient *coflnet.ChatApi
	tracer            trace.Tracer
}

func (p *ChatProcessor) StartProcessing() error {
	slog.Info("starting chat processor")

	var errCh chan error
	ctx := context.Background()

	go func() {
		slog.Info("starting redis chat processor")
		err := p.startRedisChatProcessor(ctx)
		slog.Error("redis chat processor stopped", err)
		errCh <- err
	}()

	go func() {
		slog.Info("starting discord chat processor")
		err := p.startDiscordChatProcessor(ctx)
		slog.Error("discord chat processor stopped", err)
		errCh <- err
	}()

	return <-errCh
}

func (p *ChatProcessor) startRedisChatProcessor(ctx context.Context) error {
	msgs := p.redisHandler.ReceiveChatPubSubMessage(ctx)
	slog.Info("listening for redis chat messages")

	for msg := range msgs {

		go func(msg *redisgo.Message) {
			msgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			msgCtx, span := p.tracer.Start(msgCtx, "process-redis-chat-message")
			defer span.End()
			span.SetAttributes(attribute.String("message", msg.Payload))
			span.SetAttributes(attribute.String("channel", msg.Channel))

			err := p.processRedisMessage(msgCtx, msg)
			if err != nil {
				span.RecordError(err)
				slog.Warn("error processing a message from redis", "err", err)
			}
		}(msg)
	}

	return errors.New("redis chat pubsub channel closed")
}

func (p *ChatProcessor) startDiscordChatProcessor(ctx context.Context) error {
	msgs, err := p.discordHandler.DiscordMessagesSent(ctx)
	if err != nil {
		slog.Error("error getting discord messages", err)
		return err
	}

	for msg := range msgs {
		go func(msg discordgo.Message) {
			msgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			msgCtx, span := p.tracer.Start(msgCtx, "process-discord-chat-message")
			defer span.End()

			p.processDiscordMessage(ctx, &msg)
		}(msg)
	}

	return errors.New("discord chat pubsub channel closed")
}

// processDiscordMessage processes every incoming discord message
func (p *ChatProcessor) processDiscordMessage(ctx context.Context, msg *discordgo.Message) {
	ctx, span := p.tracer.Start(ctx, "process-discord-message")
	defer span.End()

	// create an error channel for potential errors in go routines
	wg := &sync.WaitGroup{}

	// save the message to the database
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("saving discord message to database")

		err := mongo.InsertDiscordMessage(ctx, msg)
		if err != nil {
			slog.Error("error inserting discord message to database", err)
			span.RecordError(err)
		}

		metrics.MessagesInsertedIntoDatabase.Inc()
	}()

	// send the message to redis chat channel if the message was sent in the in game channel
	if p.shouldMessageBeForwardedToChat(msg) {

		slog.Debug("sending discord message to redis chat channel", "user", msg.Author.ID, "content", msg.Content)

		err := p.sendDiscordMessageToChatAPI(ctx, msg)
		if err != nil {
			slog.Error("error sending discord message to redis chat channel", err)
			span.RecordError(err)

			// send a error message to the user
			_, err = p.discordHandler.AnswerDiscordMessage(fmt.Sprintf("an error occurred while sending the message to the chat, create a bug report with the report id: %s, error msg: %s", span.SpanContext().TraceID(), err.Error()), msg)
			if err != nil {
				slog.Error("error sending error message to user", "err", err, "user", msg.Author.ID)
				span.RecordError(err)
			}
		}

		metrics.MessagesForwardedToRedisChatChannel.Inc()
	}

	wg.Wait()
	slog.Debug("finished processing discord message from %s with content %s", msg.Author.Username, msg.Content)
	metrics.DiscordMessagesProcessed.Inc()
}

// processRedisMessage processes every incoming redis message
func (p *ChatProcessor) processRedisMessage(ctx context.Context, msg *redisgo.Message) error {
	ctx, span := p.tracer.Start(ctx, "process-redis-message")
	defer span.End()
	slog.Debug("processing redis message with content %s", msg.Payload)

	// send the message to discord
	message := &mongo.ChatMessage{}
	err := json.Unmarshal([]byte(msg.Payload), message)

	if err != nil {
		slog.Error("error unmarshalling redis message", err)
		span.RecordError(err)
		return err
	}

	if message.ClientName == coflDiscordClientName {
		slog.Debug("message was sent by discord, ignoring message")
		return nil
	}

	err = p.discordHandler.SendMessageToIngameChat(ctx, message)
	if err != nil {
		slog.Error(fmt.Sprintf("could not send message to discord, message: %s", msg.Payload), err)
		return err
	}

	go func(uuid string) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		ctx, span := p.tracer.Start(ctx, "trigger-refresh-user")
		defer span.End()
		err := p.userHandler.RefreshUserByUUID(ctx, uuid)
		if err != nil {
			slog.Error("error refreshing user", err)
			span.RecordError(err)
		}
	}(message.UUID)

	metrics.MessagesForwardedToCoflnetDiscordChatChannel.Inc()
	return nil
}

func (p *ChatProcessor) sendDiscordMessageToChatAPI(ctx context.Context, msg *discordgo.Message) error {
	ctx, span := p.tracer.Start(ctx, "send-discord-message-to-chat-api")
	defer span.End()

	users, err := mongo.UsersByDiscordAccount(ctx, msg.Author)
	if err != nil {
		slog.Error("error getting user from database", err)
		span.RecordError(err)
		return err
	}

	if len(users) == 0 {
		err := errors.New(fmt.Sprintf("no user found for discord account %s, can not forward the message\n To resolve this try the following:\n1. Make sure your hypixel account is connected to your discord account\n2. Try to write a chat message within minecraft with the `/fc <msg>, this is necessary to connect your minecraft account with your discord account.\n3. If this does not work, please create a bug report with the id: %s", msg.Author.Username, span.SpanContext().TraceID()))
		go func() {
			_, e := p.discordHandler.AnswerDiscordMessage(err.Error(), msg)
			if e != nil {
				slog.Error("error sending error message to user", "err", e, "user", msg.Author.ID)
				span.RecordError(e)
			}
		}()
		slog.Warn("error searching user for chat message", err)
		span.RecordError(err)
		return err
	}

	user := utils.FilterUsersForPreferredUsers(msg.Author.ID, users)

	span.SetAttributes(attribute.String("uuid", user.UUID()))
	span.SetAttributes(attribute.Int("user-id", user.UserId))

	if len(users) == 0 {
		err := errors.New(fmt.Sprintf("no user found for discord account %s, can not forward the message\n To resolve this try the following:\n1. Make sure your hypixel account is connected to your discord account\n2. Try to write a chat message within minecraft with the `/fc <msg>, this is necessary to connect your minecraft account with your discord account.\n3. If this does not work, please create a bug report with the id: %s", msg.Author.Username, span.SpanContext().TraceID()))
		p.discordHandler.AnswerDiscordMessage(err.Error(), msg)
		slog.Warn("error searching user for chat message", err)
		span.RecordError(err)
		return err
	}

	if user.UUID() == "" {
		err := errors.New(fmt.Sprintf("no user found for discord account %s, can not forward the message\n To resolve this try the following:\n1. Make sure your hypixel account is connected to your discord account\n2. Try to write a chat message within minecraft with the `/fc <msg>, this is necessary to connect your minecraft account with your discord account.\n3. If this does not work, please create a bug report with the id: %s", msg.Author.Username, span.SpanContext().TraceID()))
		slog.Warn(fmt.Sprintf("no uuid found for user"), err)
		span.RecordError(err)
		return err
	}

	slog.Info(fmt.Sprintf("sending discord message from %s(%s) to chat api, message: %s", msg.Author, user.UUID(), msg.Content))

	// TODO refactor
	playerName, err := coflnet.PlayerName(user.UUID())
	if err != nil {
		slog.Error("error getting player name from coflnet", err)
		span.RecordError(err)
		return err
	}

	slog.Info(fmt.Sprintf("sending discord message from %s(%s) to chat api, message: %s", msg.Author, user.UUID(), msg.Content))
	err = p.coflnetChatClient.SendMessage(ctx, msg.Content, user.UUID(), coflDiscordClientName, playerName)

	if err != nil {
		slog.Error("error sending discord message to chat api", err)
		span.RecordError(err)
		return err
	}

	slog.Info("a message was sent to chat api")
	span.SetAttributes(attribute.Bool("message_sent", true))
	return nil
}

func (p *ChatProcessor) shouldMessageBeForwardedToChat(msg *discordgo.Message) bool {
	if msg.ChannelID != utils.DiscordInGameChannelId() {
		slog.Debug("message was not sent in in game channel, skipping")
		return false
	}

	if msg.Author.Bot {
		slog.Debug("message was sent by a bot, skipping")
		return false
	}

	// if the message was sent by the bot skip it
	if msg.Author.ID == "888725077191974913" {
		slog.Debug("message was sent by the bot, skipping")
		return false
	}

	slog.Debug("message should be forwarded to chat")
	return true
}
