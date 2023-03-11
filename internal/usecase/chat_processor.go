package usecase

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
	"github.com/Coflnet/coflnet-bot/internal/redis"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/Coflnet/coflnet-bot/schemas/chat"
	"github.com/bwmarrin/discordgo"
	redisgo "github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
    chatProcessor = "chat-processor"
    coflDiscordClientName = "cofl-discord"
)

type ChatProcessor struct {
    redisHandler *redis.RedisHandler
    discordHandler *discord.DiscordHandler
    coflnetChatClient *coflnet.ChatApi
    tracer trace.Tracer
}

func (r *ChatProcessor) init() error {
    var err error
    r.tracer = otel.Tracer(flipProcessorTracerName)
    r.redisHandler = redis.NewRedisHandler()

    r.discordHandler, err = discord.NewDiscordHandler()
    if err != nil {
        return err
    }

    r.coflnetChatClient, err = coflnet.NewChatClient()
    return err
}

func (r *ChatProcessor) StartProcessing() error {

    if err := r.init(); err != nil {
        return err
    }

    var oneDone = make(chan error)

    go func() {
        err := r.StartRedisChatProcessor()
        slog.Error("redis chat processor stopped", err)
        oneDone <- err
    }()

    go func() {
        err := r.StartDiscordChatProcessor()
        slog.Error("discord chat processor stopped", err)
        oneDone <- err
    }()

    err := <-oneDone
    panic(err)
}

func (r *ChatProcessor) StartRedisChatProcessor() error {
	ctx := context.Background()
    msgs := r.redisHandler.ReceiveChatPubSubMessage(ctx)

    for msg := range msgs {

        go func(msg redisgo.Message) {
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            defer cancel()

            ctx, span := r.tracer.Start(ctx, "process-redis-chat-message")
            defer span.End()

            err := r.processRedisMessage(ctx, &msg)
            if err != nil {
                span.RecordError(err)
                log.Error().Err(err).Msg("error processing message")
            }
        }(msg)
    }

    return errors.New("redis chat pubsub channel closed")
}

func (r *ChatProcessor) StartDiscordChatProcessor() error {

    ctx := context.Background()
    slog.Info("starting discord chat processor")

    msgs, err := r.discordHandler.DiscordMessagesSent(ctx)
    if err != nil {
        slog.Error("error getting discord messages", err)
        return err
    }

    for msg := range msgs {
        go func (msg discordgo.Message) {
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            defer cancel()

            ctx, span := r.tracer.Start(ctx, "process-discord-chat-message")
            defer span.End()

            err := r.processDiscordMessage(ctx, &msg)
            if err != nil {
                span.RecordError(err)
                log.Error().Err(err).Msg("error processing message")
            }
        }(msg)
    }

    return errors.New("discord chat pubsub channel closed")
}


// processDiscordMessage processes every incoming discord message
func (p *ChatProcessor) processDiscordMessage(ctx context.Context, msg *discordgo.Message) error {

    ctx, span := p.tracer.Start(ctx, "process-discord-message")
    defer span.End()
    defer metrics.DiscordMessagesProcessed.Inc()
    slog.Debug("processing discord message from %s with content %s", msg.Author.Username, msg.Content)


    // create a error channel for potential errors in go routines
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

        slog.Debug("message was sent in in game channel, forwarding message to redis chat channel")
        wg.Add(1)
        go func() {
            defer wg.Done()
            slog.Debug("sending discord message to redis chat channel")

            err := p.sendDiscordMessageToChatAPI(ctx, msg)
            if err != nil {
                slog.Error("error sending discord message to redis chat channel", err)
                span.RecordError(err)
            }
            
            metrics.MessagesForwardedToRedisChatChannel.Inc()
        }()
    }

    wg.Wait()
    slog.Debug("finished processing discord message from %s with content %s", msg.Author.Username, msg.Content)
    return nil
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

	err = p.discordHandler.SendMessageToIngameChat(message)
	if err != nil {
		log.Error().Err(err).Msgf("could not send message to discord, message: %s", msg.Payload)
		return err
	}

    metrics.MessagesForwardedToCoflnetDiscordChatChannel.Inc()
    return nil
}


func (p *ChatProcessor) sendDiscordMessageToChatAPI(ctx context.Context, msg *discordgo.Message) error {
    ctx, span := p.tracer.Start(ctx, "send-discord-message-to-redis-chat-channel")
    defer span.End()

    users, err := mongo.UsersByDiscordAccount(ctx, msg.Author)
    if err != nil {
        slog.Error("error getting user from database", err)
        span.RecordError(err)
        return err
    }

    if len(users) >= 2  {
        users = utils.FilterUsersForPreferredUsers(msg.Author.ID, users)
    }

    if len(users) >= 2 {
        err := errors.New("more than one user found for discord account even after filtering for preferred users")
        slog.Warn("error searching user for chat message", err)
        span.RecordError(err)
        return err
    }

    if len(users) == 0 {
        err := errors.New(fmt.Sprintf("no user found for discord account %s, can not forward the message, error: %s", msg.Author.Username, span.SpanContext().TraceID()))
        p.discordHandler.AnswerDiscordMessage(err.Error(), msg)
        slog.Warn("error searching user for chat message", err)
        span.RecordError(err)
        return err
    }

    user := users[0]

    err = p.coflnetChatClient.SendMessage(ctx, &chat.APIChatSendPostReqApplicationJSON{
        Message: chat.OptNilString{
            Null: false,
            Value: msg.Content,
        },
        UUID: chat.OptNilString{
            Null: false,
            Value: user.UUID(),
        },
        ClientName: chat.OptNilString{
            Value: coflDiscordClientName,
            Null: false,
        },
    })

    if err != nil {
        slog.Error("error sending discord message to chat api", err)
        span.RecordError(err)
        return err
    }

    return nil
}

func (p *ChatProcessor) shouldMessageBeForwardedToChat(msg *discordgo.Message) bool {
    if msg.ChannelID != utils.DiscordInGameChannelId() {
        slog.Debug("message was not sent in in game channel, skipping")
        return false
    }

    if msg.Author.Bot  {
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
