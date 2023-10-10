package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/Coflnet/coflnet-bot/internal/utils"
)

type DiscordMessageToSend struct {
	Message string
	Channel DiscordChannel
	Webhook string
}

const (
	WarningsChannel          DiscordChannel = "warnings"
	MutesChannel             DiscordChannel = "mutes"
	FlipperRoleChannel       DiscordChannel = "flipper-role"
	CiSuccessChannel         DiscordChannel = "ci-success"
	CiFailureChannel         DiscordChannel = "ci-failure"
	FeedbackChannel          DiscordChannel = "feedback"
	FlipChannel              DiscordChannel = "flip"
	SongvoterFeedbackChannel DiscordChannel = "songvoter-feedback"
)

type DiscordChannel string

type Message struct {
	Message         string          `json:"message"`
	Channel         DiscordChannel  `json:"channel"`
	Webhook         string          `json:"webhook"`
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

func SendMessageToSongvoterFeedbackChannel(msg string) error {
	return sendMessageToChannel(msg, SongvoterFeedbackChannel)
}

func sendMessageToChannel(msg string, channel DiscordChannel) error {
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

func message(msg string, channel DiscordChannel) ([]byte, error) {
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

func SendMessageToChannel(content string, channel DiscordChannel) error {

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
