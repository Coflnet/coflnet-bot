package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func SendMessageToChatApi(msg *discordgo.MessageCreate) error {

  // if the message was not sent in cofl chat skip it
	if msg.ChannelID != coflChatId() {
		return nil
	}

	// if the message was sent by the bot skip it
	if msg.Author.ID == "888725077191974913" {
		return nil
	}

	if msg.Author.Bot {
		return nil
	}

	// webhook
	if msg.Author.ID == "975127829916286986" {
		return nil
	}


	log.Info().Msgf("sending message to chat api")

	apiKey := os.Getenv("COFL_CHAT_API_KEY")
	baseUrl := os.Getenv("COFL_CHAT_BASE_URL")
	p := "/api/chat/send"

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing base url")
		return err
	}

	u.Path = path.Join(u.Path, p)

	username := fmt.Sprintf("%s#%s", msg.Author.Username, msg.Author.Discriminator)
	log.Info().Msgf("searching username %s", msg.Author.Username)

  err = RefreshUserByPlayername(msg.Author.Username)
  if err != nil {
    // ignore error
  }

	user, err := mongo.SearchByDiscordTag(username)
	if err != nil {
		log.Warn().Err(err).Msgf("error when searching user %s in db", username)
    sendInvalidUUIDMessageToDiscord(msg.Message, fmt.Sprintf("user with the name %s not found, please check if your discord name is set correctly on your hypixel profile, maybe try the /refresh-user command", username))
		return err
	}

	if user == nil {
		log.Warn().Msgf("user %s not found in db", username)

		sendInvalidUUIDMessageToDiscord(msg.Message, fmt.Sprintf("user with the name %s not found, please check if your discord name is set correctly on your hypixel profile, maybe try the /refresh-user command", username))

		return nil
	}

	log.Info().Msgf("found uuid for player %s", msg.Author.Username)

	// TODO user minecraft uuids at 0 is potentially wrong
	payload := &ChatApiPayload{
		Message:    msg.Content,
		UUID:       user.MinecraftUuids[0],
		ClientName: "cofl-discord",
		Prefix:     "[cofl-dc] ",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msgf("error serializing payload")
		return err
	}

	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msgf("error creating request")
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", apiKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Info().Msgf("message: %s, name: %s, uuid: %s, clientName: %s, prefix: %s", payload.Message, payload.Name, payload.UUID, payload.ClientName, payload.Prefix)

	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msgf("error sending request, status: %s")
	}

	log.Info().Msgf("response code %s", response.Status)

	return nil
}

type PlayerSearchResult struct {
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
	HitCount int    `json:"hitCount"`
}

type ChatApiPayload struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Prefix     string `json:"prefix"`
	Message    string `json:"message"`
	ClientName string `json:"clientName"`
}
