package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"net/http"
	"os"
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
	FlipChannel        = "flip"
)

type DiscordMessage struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Webhook string `json:"webhook"`
}

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
	payload, err := message(msg, channel)
	if err != nil {
		return err
	}

	// make http POST request
	host, err := utils.CoflnetBotBaseUrl()
	if err != nil {
		return err
	}

	_, err = http.Post(fmt.Sprintf("%s/api/webhook/message", host), "application/json", bytes.NewBuffer(payload))
	return err
}

func message(msg, channel string) ([]byte, error) {
	v := DiscordMessageToSend{
		Message: msg,
		Channel: channel,
	}

	b, err := json.Marshal(v)

	if err != nil {
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
