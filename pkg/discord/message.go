package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"log/slog"
	"net/http"
	"os"
	"strings"
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

type Message struct {
	Message string `json:"message"`
	Channel string `json:"channel"`
	Webhook string `json:"webhook"`
    AllowedMentions AllowedMentions `json:"allowed_mentions"`
}

type AllowedMentions struct {
    Parse []string `json:"parse"`
    Users []string `json:"users"`
    Roles []string `json:"roles"`
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

func SendMessageToChannel(content string, channel string) error {

	// prepare a http post request
	msg := Message{
		Channel: channel,
		Message: content,
        AllowedMentions: AllowedMentions{
            Parse: []string{"users", "roles"},
            Users: []string{},
            Roles: []string{},
        },
	}

	// send the message
	base := os.Getenv("DISCORD_BOT_BASE_URL")
	if base == "" {
		return fmt.Errorf("DISCORD_BOT_BASE_URL is not set")
	}

	url := fmt.Sprintf("%s/api/webhook/message", base)

	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			slog.Error("could not close body", err)
		}
	}()

	slog.Debug("discord webhook response status code", "code", res.StatusCode)
	return nil
}
