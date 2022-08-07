package kafka

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type TransactionMessage struct {
	Id           int64     `json:"Id"`
	UserId       string    `json:"UserId"`
	ProductSlug  string    `json:"ProductSlug"`
	ProductId    int64     `json:"ProductId"`
	OwnedSeconds int64     `json:"OwnedSeconds"`
	Amount       float64   `json:"Amount"`
	Timestamp    time.Time `json:"Timestamp"`
}

type VerificationMessage struct {
	UserId int64 `msgpack:"alias:UserId"`
}

var flipSummaryReader *kafka.Reader

func Init() {
	flipSummaryReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "flip-summary-discord-consumer",
		Topic:    flipSummaryTopic(),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func StartVerificationConsume() error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "verification-discord-consumer",
		Topic:    transactionTopic(),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)

		if err != nil {
			log.Error().Err(err).Msgf("error occured when fetching from kafka, verification topic")
			return err
		}

		err = ProcessVerificationMessage(&m)
		if err != nil {
			log.Warn().Err(err).Msgf("error processing verification message, ignoring")
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

	// only check if amount is negative, recommended by ekwav
	if t.Amount < 0 {
		log.Info().Msgf("got a transaction message with negative amount, ignoring")
		return nil
	}

	log.Info().Msgf("got a %s message from user %d", t.ProductSlug, t.UserId)
	id, err := strconv.Atoi(t.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("could not parse userId from kafka %s", t.UserId)
		return err
	}

	user, err := coflnet.UserById(id)

	if err != nil {
		log.Error().Err(err).Msgf("loading user with id failed %d, cannot update role", t.UserId)
		return err
	}

	return discord.SetFlipperRoleForUser(user)
}

func ProcessVerificationMessage(message *kafka.Message) error {
	content := message.Value

	var msg VerificationMessage
	err := json.Unmarshal(content, &msg)
	if err != nil {
		return err
	}

	log.Info().Msgf("deserialized verification message for userId %d", msg.UserId)
	log.Warn().Msg("processing of verifications is not implemented yet")
	return nil
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
