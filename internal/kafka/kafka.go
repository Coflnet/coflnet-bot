package kafka

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type VerificationMessage struct {
	UserId string `json:"UserId"`
}

var (
	flipSummaryReader *kafka.Reader
	transactionReader *kafka.Reader
)

func Init() {
	flipSummaryReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "flip-summary-discord-consumer",
		Topic:    flipSummaryTopic(),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	transactionReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "verification-discord-consumer",
		Topic:    transactionTopic(),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func ReadTransactionMessage() (kafka.Message, error) {
	ctx := context.Background()
	return transactionReader.FetchMessage(ctx)
}

func CommitTransactionMessage(msg kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return transactionReader.CommitMessages(ctx, msg)
}

func transactionTopic() string {
	t := os.Getenv("TOPIC_TRANSACTION")
	if t == "" {
		log.Panic().Msg("TOPIC_TRANSACTION not set")
	}

	return t
}

func flipSummaryTopic() string {
	t := os.Getenv("TOPIC_FLIPSUMMARY")
	if t == "" {
		log.Panic().Msg("TOPIC_FLIPSUMMARY not set")
	}

	return t
}
