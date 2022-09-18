package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

type DiscordMessageToSend struct {
	Message string
	Channel string
	Webhook string
}

const (
	WarningsChannel    = "warnings"
	MutesChannel       = "mutes"
	FlipperRoleChannel = "flipper-role"
	CiSuccessChannel   = "ci-success"
	CiFailureChannel   = "ci-failure"
	FeedbackChannel    = "feedback"
)

func SendMessageToFeedback(msg string) error {
	return sendMessageToChannel(msg, FeedbackChannel)
}

func SendMessageToCiSuccess(msg string) error {
	return sendMessageToChannel(msg, CiSuccessChannel)
}

func SendMessageToCiFailureChannel(msg string) error {
	return sendMessageToChannel(msg, CiFailureChannel)
}

func SendMessageToWarningsChannel(msg string) error {
	return sendMessageToChannel(msg, WarningsChannel)
}

func SendMessageToMutesChannel(msg string) error {
	return sendMessageToChannel(msg, MutesChannel)
}

func SendMessageToRoleChannel(msg string) error {
	return sendMessageToChannel(msg, FlipperRoleChannel)
}

func sendMessageToChannel(msg, channel string) error {
	w, err := writer()
	if err != nil {
		return err
	}

	defer func(w *kafka.Writer) {
		err := w.Close()
		if err != nil {
			log.Error().Err(err).Msgf("there was an error when closing the writer")
		}
	}(&w)

	b, err := message(msg, channel)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(fmt.Sprintf("%d", time.Now().Unix())),
		Value: b,
	})
}

func writer() (kafka.Writer, error) {
	h, err := kafkaHost()
	if err != nil {
		return kafka.Writer{}, err
	}

	t, err := discordMessageTopic()
	if err != nil {
		return kafka.Writer{}, err
	}

	return kafka.Writer{
		Addr:  kafka.TCP(h),
		Topic: t,
	}, nil
}

func message(msg, channel string) ([]byte, error) {
	v := DiscordMessageToSend{
		Message: msg,
		Channel: channel,
	}

	b, err := json.Marshal(v)

	if err != nil {
		log.Error().Err(err).Msgf("there was an error when marshaling the message")
		return nil, err
	}

	return b, nil
}

func kafkaHost() (string, error) {
	return getEnv("KAFKA_HOST")
}

func discordMessageTopic() (string, error) {
	return getEnv("DISCORD_MESSAGES_TOPIC")
}

func getEnv(k string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		return "", fmt.Errorf("%s not set", k)
	}
	return v, nil
}
