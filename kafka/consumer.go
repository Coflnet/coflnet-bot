package kafka

import (
	"os"
	"strconv"

	"github.com/Coflnet/coflnet-bot/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(session *discordgo.Session, errorCh chan<- error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaHost(),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{KafkaTopic()}, nil)

	log.Info().Msgf("cosumer subscribed to topic %s", KafkaTopic())

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Err(err).Msgf("Consumer error: %v (%v)\n", err, msg)
			errorCh <- err
			return
		}

		go func() {
			id, err := strconv.Atoi(string(msg.Value))

			if err != nil {
				log.Error().Err(err).Msgf("got a kafka message, that is not an integer %s", msg.Value)
			}

			log.Info().Msgf("got a message for the user %s", id)

			discord.UpdateUserRole(id, session)
		}()
	}
}
