package redis

import (
	"context"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
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

	r.rdb = redis.NewClient(&redis.Options{
		Addr:    utils.RedisHost(), 
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return r
}

func (r *RedisHandler) ReceiveChatPubSubMessage(ctx context.Context) <-chan redis.Message {
	ch := make(chan redis.Message)

	go func() {
		defer close(ch)

		pubsub := r.rdb.Subscribe(ctx, utils.RedisChatPubSubChannel())
		defer pubsub.Close()

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Error().Err(err).Msg("failed to receive message")
				continue
			}

			log.Debug().Msgf("receive message: %s", msg.Payload)

			ch <- *msg
		}
	}()

	return ch
}
