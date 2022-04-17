package chat

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

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func sendMessageToChatApi(msg *discordgo.MessageCreate) error {

	if msg.ChannelID != coflChatId {
		log.Info().Msgf("message not in cofl chat, skipping")
		return nil
	}

	if msg.Author.ID == "888725077191974913" {
		log.Info().Msgf("message from cofl bot, skipping")
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

	uuid := GetUuidForPlayer(msg.Author.Username)
	if uuid != "" {
		log.Info().Msgf("found uuid for player %s", msg.Author.Username)
	} else {
		log.Info().Msgf("no uuid found for player %s", msg.Author.Username)
	}

	payload := &ChatApiPayload{
		Message:    msg.Content,
		Name:       msg.Author.Username,
		UUID:       uuid,
		ClientName: "coflnet-discord",
		Prefix:     "d",
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

	_ = &http.Client{
		Timeout: 10 * time.Second,
	}

	log.Warn().Msgf("not sending request, for test purpose")
	log.Info().Msgf("message: %s, name: %s, uuid: %s, clientName: %s, prefix: %s", payload.Message, payload.Name, payload.UUID, payload.ClientName, payload.Prefix)

	// response, err := client.Do(requet)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("error sending request, status: %s")
	// }

	// log.Info().Msgf("response code %s", response.Status)
	return nil
}

func GetUuidForPlayer(name string) string {
	apiBaseUrl := os.Getenv("API_BASE_URL")
	p := fmt.Sprintf("/api/search/player/%s", name)

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
