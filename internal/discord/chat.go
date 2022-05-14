package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	if msg.ChannelID != coflChatId {
		return nil
	}

	if msg.Author.ID == "888725077191974913" {
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

	user, err := mongo.SearchByDiscordTag(msg.Author.Username)
	if err != nil {
		log.Error().Err(err).Msgf("error when searching user %s in db", msg.Author.Username)
		return err
	}

	if user == nil {
		log.Warn().Msgf("user %s not found in db", msg.Author.Username)

		// TODO notify user
		// sendInvalidUUIDMessageToDiscord(msg.Message)

		return nil
	}

	log.Info().Msgf("found uuid for player %s", msg.Author.Username)

	// TODO user minecraft uuids at 0 is potentially wrong
	payload := &ChatApiPayload{
		Message:    msg.Content,
		Name:       msg.Author.Username,
		UUID:       user.MinecraftUuids[0],
		ClientName: "cofl-discord",
		Prefix:     "cofl-dc",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msgf("error serializing payload")
		return err
	}

	requet, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msgf("error creating request")
		return err
	}

	requet.Header.Set("Content-Type", "application/json")
	requet.Header.Set("Authorization", apiKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Info().Msgf("message: %s, name: %s, uuid: %s, clientName: %s, prefix: %s", payload.Message, payload.Name, payload.UUID, payload.ClientName, payload.Prefix)

	response, err := client.Do(requet)
	if err != nil {
		log.Error().Err(err).Msgf("error sending request, status: %s")
	}

	log.Info().Msgf("response code %s", response.Status)

	session.ChannelMessageDelete(msg.ChannelID, msg.Message.ID)

	return nil
}

func GetUuidForPlayerDeprecated(name string) string {

	log.Warn().Msg("this method should not be used to get the uuid of a player")

	apiBaseUrl := os.Getenv("COFL_API_BASE_URL")
	p := fmt.Sprintf("/search/player/%s", name)

	u, err := url.Parse(apiBaseUrl)
	if err != nil {
		log.Error().Err(err).Msgf("error parsing base url")
		return ""
	}

	u.Path = path.Join(u.Path, p)

	log.Info().Msgf("getting uuid for player %s, request: %s", name, u.String())

	response, err := http.Get(u.String())
	if err != nil {
		log.Error().Err(err).Msgf("error getting uuid for player %s", name)
		return ""
	}

	if response.StatusCode != 200 {
		log.Error().Msgf("error getting uuid for player %s, status: %s", name, response.Status)
		return ""
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error reading response body")
		return ""
	}

	var results []PlayerSearchResult
	err = json.Unmarshal(body, &results)

	for _, result := range results {
		if result.Name == name {
			return result.UUID
		}
	}

	return ""
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
