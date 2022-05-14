package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var (
	rdb          *redis.Client
	redisChannel string
)

func Init() error {

	host := os.Getenv("REDIS_HOST")

	if host == "" {
		return fmt.Errorf("missing redis host")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}

func StartConsume() error {

	err := Init()
	if err != nil {
		return err
	}

	if rdb == nil {
		return fmt.Errorf("redis not initialized")
	}

	redisChannel = os.Getenv("REDIS_CHAT_CHANNEL")

	if redisChannel == "" {
		return fmt.Errorf("missing redis channel")
	}

	ctx := context.Background()

	log.Info().Msgf("start redis chat consume..")
	pubsub := rdb.Subscribe(ctx, redisChannel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}

		log.Info().Msgf("received redis message: %s", msg.Payload)
		processMessage(msg)
	}
}

func processMessage(msg *redis.Message) error {

	message := &mongo.ChatMessage{}

	err := json.Unmarshal([]byte(msg.Payload), message)

	if err != nil {
		log.Error().Err(err).Msgf("could not unmarshal message, message: %s", msg.Payload)
		return err
	}

	if message.ClientName == "coflnet-discord" {
		return nil
	}

	err = mongo.InsertChatMessage(message)
	if err != nil {
		log.Error().Err(err).Msgf("could not insert message, message: %s", msg.Payload)
		return err
	}

	err = SendMessageToDiscordChat(message)
	if err != nil {
		log.Error().Err(err).Msgf("could not send message to discord, message: %s", msg.Payload)
		return err
	}

	return nil
}
