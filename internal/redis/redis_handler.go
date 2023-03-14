package redis

import (
	"context"
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
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

	r.rdb = redis.NewClient(&redis.Options{
		Addr:    utils.RedisHost(), 
		Password: "", // no password set
		DB:       0,  // use default DB
	})

    slog.Info(fmt.Sprintf("use redis host %s", utils.RedisHost()))

	return r
}

func (r *RedisHandler) ReceiveChatPubSubMessage(ctx context.Context) <-chan *redis.Message {
	ch := make(chan *redis.Message, 100)

	go func() {
        defer panic("redis receive stopped")
		defer close(ch)

        slog.Info(fmt.Sprintf("start listening for chat messages on channel %s", utils.RedisChatPubSubChannel()))
		pubsub := r.rdb.Subscribe(ctx, utils.RedisChatPubSubChannel())
		defer pubsub.Close()

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
