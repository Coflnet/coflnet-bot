package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
)

type RedisMessageService struct {
	rdb            *redis.Client
	messageChannel chan *RedisMessage
}

type RedisMessage struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Prefix     string `json:"prefix"`
	Message    string `json:"message"`
	ClientName string `json:"clientName"`
}

func NewRedisMessageService() *RedisMessageService {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
	})

	res := rdb.Ping(context.Background())
	if res.Err() != nil {
		panic(res.Err())
	}
	channel := make(chan *RedisMessage)

	return &RedisMessageService{
		rdb:            rdb,
		messageChannel: channel,
	}
}

func (r *RedisMessageService) Reader(ctx context.Context) (<-chan *RedisMessage, error) {
	go func() {
		defer close(r.messageChannel)

		slog.Info(fmt.Sprintf("start listening for chat messages on channel %s", r.chatChannel()))
		pubsub := r.rdb.Subscribe(ctx, r.chatChannel())

		defer func(pubsub *redis.PubSub) {
			err := pubsub.Close()
			if err != nil {
				slog.Error("failed to close pubsub", err)
				close(r.messageChannel)
			}
		}(pubsub)

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				slog.Error("failed to receive message", err)
				continue
			}

			redisMsg, err := r.convertRawMessageInRedisMessage(msg)
			if err != nil {
				slog.Error("failed to convert redis message", err)
				continue
			}

			slog.Debug(fmt.Sprintf("receive message: %s", msg.Payload))
			r.messageChannel <- redisMsg
		}
	}()

	return r.messageChannel, nil
}

func (r *RedisMessageService) convertRawMessageInRedisMessage(msg *redis.Message) (*RedisMessage, error) {
	var result RedisMessage
	err := json.Unmarshal([]byte(msg.Payload), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *RedisMessageService) chatChannel() string {
	ch, found := os.LookupEnv("REDIS_CHAT_CHANNEL")
	if !found {
		panic("REDIS_CHAT_CHANNEL not found")
	}
	return ch
}
