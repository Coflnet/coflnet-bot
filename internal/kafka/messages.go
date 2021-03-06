package kafka

import (
	"context"
	"encoding/json"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"os"
)

func StartDiscordMessagesConsumer() {

	log.Info().Msgf("consuming topic %s", os.Getenv("TOPIC_DEV_CHAT"))

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "coflnet-bot-cg",
		Topic:    os.Getenv("TOPIC_DEV_CHAT"),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	defer func() {
		if err := r.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close kafka reader")
		}
	}()

	for {
		err := consumeDiscordMessages(r)
		if err != nil {
			log.Error().Err(err).Msg("error consuming discord messages")
		}
	}
}

func StartDiscordSpamMessagesConsumer() {

	log.Info().Msgf("start consuming topic %s", os.Getenv("TOPIC_DEV_SPAM_CHAT"))

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  "coflnet-bot-cg",
		Topic:    os.Getenv("TOPIC_DEV_SPAM_CHAT"),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	defer func() {
		if err := r.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close kafka reader")
		}
	}()

	for {
		err := consumeDiscordSpamMessages(r)
		if err != nil {
			log.Error().Err(err).Msg("error consuming discord spam messages")
		}
	}
}

func consumeDiscordMessages(r *kafka.Reader) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m, err := r.FetchMessage(ctx)
	if err != nil {
		return err
	}

	var msg *discord.DiscordMessageToSend
	err = json.Unmarshal(m.Value, &msg)
	if err != nil {
		return err
	}

	err = discord.SendMessageToDevLog(msg)
	if err != nil {
		return err
	}

	if err := r.CommitMessages(ctx, m); err != nil {
		return err
	}

	return nil
}

func consumeDiscordSpamMessages(r *kafka.Reader) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m, err := r.FetchMessage(ctx)
	if err != nil {
		return err
	}

	var msg *discord.DiscordMessageToSend
	err = json.Unmarshal(m.Value, &msg)
	if err != nil {
		return err
	}

	err = discord.SendMessageToDevSpamLog(msg)
	if err != nil {
		return err
	}

	if err := r.CommitMessages(ctx, m); err != nil {
		return err
	}

	return nil
}
