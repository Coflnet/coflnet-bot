package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	redisHandlerTracerName = "redis-handler"
)

type Redis interface {
}

type RedisHandler struct {
	rdb    *redis.Client
	tracer trace.Tracer
}

func NewRedisHandler() *RedisHandler {
	r := &RedisHandler{
		tracer: otel.Tracer(redisHandlerTracerName),
	}

	redisUrl, err := utils.RedisHost()
	if err != nil {
		slog.Error("error getting redis url", err)
		return r
	}

	r.rdb = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	slog.Info(fmt.Sprintf("use redis host %s", redisUrl))
	return r
}

func (r *RedisHandler) ReceiveChatPubSubMessage(ctx context.Context) <-chan *redis.Message {
	ch := make(chan *redis.Message, 100)

	if r.rdb == nil {
		slog.Error("redis client not initialized")
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)

		slog.Info(fmt.Sprintf("start listening for chat messages on channel %s", utils.RedisChatPubSubChannel()))
		pubsub := r.rdb.Subscribe(ctx, utils.RedisChatPubSubChannel())

		defer func(pubsub *redis.PubSub) {
			err := pubsub.Close()
			if err != nil {
				slog.Error("failed to close pubsub", err)
				close(ch)
			}
		}(pubsub)

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				slog.Error("failed to receive message", err)
				continue
			}
			slog.Debug(fmt.Sprintf("receive message: %s", msg.Payload))
			ch <- msg
		}
	}()

	return ch
}
