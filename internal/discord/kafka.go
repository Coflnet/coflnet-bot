package discord

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type TransactionMessage struct {
	Id           int64     `json:"Id"`
	UserId       string    `json:"UserId"`
	ProductSlug  string    `json:"ProductSlug"`
	ProductId    int64     `json:"ProductId"`
	OwnedSeconds int64     `json:"OwnedSeconds"`
	Amount       int64     `json:"Amount"`
	Timestamp    time.Time `json:"Timestamp"`
}

func StartTransactionConsume() error {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "transaction-discord-consumer",
		Topic:    os.Getenv("TOPIC_TRANSACTION"),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)

		if err != nil {
			log.Error().Err(err).Msgf("error occured when fetching from kafka")
			return err
		}

		log.Info().Msgf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		err = ProcessTransactionMessage(&m)
		if err != nil {
			log.Error().Err(err).Msgf("error processing kafka transaction message")
			continue
		}

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("failed to commit messages:")
		}

	}
}

func ProcessTransactionMessage(message *kafka.Message) error {

	t := TransactionMessage{}
	err := json.Unmarshal(message.Value, &t)
	if err != nil {
		log.Error().Err(err).Msgf("error deserializing kafka message: %v", string(message.Value))
		return err
	}

	log.Info().Msgf("got a %s message from user %d", t.ProductSlug, t.UserId)

	return nil
}
