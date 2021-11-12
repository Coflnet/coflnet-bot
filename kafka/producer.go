package kafka

import (
	"context"
	"time"

	"github.com/Coflnet/coflnet-bot/database"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

var messages = make(chan string, 1000)

func StartProducer(errorCh chan<- error) {
	log.Info().Msgf("connecting to host %s", KafkaHost())
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": KafkaHost()})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Error().Msgf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := KafkaTopic()
	log.Info().Msgf("producing in topic %s", topic)
	for {

		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(<-messages),
		}, nil)

		if err != nil {
			errorCh <- err
		}

		// Wait for message deliveries before shutting down
		// timeout of 15 seconds
		p.Flush(15 * 1000)
	}
}

func StartMocking() string {
	for {
		messages <- "1"
	}
}

func CheckAllUsers() {
	// rdb := database.RedisClient()

	users, err := database.AllUsers()

	if err != nil {
		log.Error().Err(err).Msgf("stopping checkAllUsers error occured")
	}

	for i, user := range users {
		log.Info().Msgf("checking %d. user %s", i, user.Id)

		checked, err := checkedUserChecked(user.Id)
		if err != nil {
			log.Error().Err(err).Msgf("error checking user %s in cache", user.Id)
			continue
		}
		if checked {
			log.Info().Msgf("user %s already checked, skipping", user.Id)
			time.Sleep(time.Second * 3)
			continue
		}

		messages <- user.Id
		log.Info().Msgf("put %s into the messages channel", user.Id)

		err = saveUserChecked(user.Id)
		if err != nil {
			log.Error().Err(err).Msgf("error when saving user %s to cache")
		}

		time.Sleep(time.Minute * 1)
	}
}

func checkedUserChecked(userId string) (bool, error) {
	userId = "checked_" + userId
	ctx := context.Background()
	rdb := database.RedisClient()

	result, err := rdb.Get(ctx, userId).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	if result == "true" {
		return true, nil
	} else {
		return false, nil
	}
}

func saveUserChecked(userId string) error {
	userId = "checked_" + userId
	rdb := database.RedisClient()
	ctx := context.Background()

	return rdb.Set(ctx, userId, "true", time.Hour*24).Err()
}
