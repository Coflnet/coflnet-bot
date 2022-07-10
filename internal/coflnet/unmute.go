package coflnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type UnMuteRequest struct {
	Unmuter string `json:"unMuter"`
	Reason  string `json:"reason"`
	Uuid    string `json:"uuid"`
}

func UnmutePlayer(unmute *model.Unmute) error {

	msg, err := mongo.FirstChatMessageOfUsername(unmute.Username)
	if err != nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", unmute.Username)
		return err
	}

	if msg == nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", unmute.Username)
		return fmt.Errorf("could not find user %s in chat collection", unmute.Username)
	}

	muteRequest := UnMuteRequest{
		Uuid:    msg.UUID,
		Unmuter: unmute.Unmuter,
		Reason:  unmute.Reason,
	}

	data, err := json.Marshal(muteRequest)
	if err != nil {
		log.Error().Err(err).Msgf("can not marshal unmute request")
		return err
	}

	log.Info().Msgf("sending unmute request: %s", string(data))

	apiKey := os.Getenv("COFL_CHAT_API_KEY")
	baseUrl := os.Getenv("COFL_CHAT_BASE_URL")
	p := "/api/chat/mute"
	url := fmt.Sprintf("%s%s", baseUrl, p)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(data))

	if err != nil {
		log.Error().Err(err).Msgf("can not create unmute request")
		return err
	}

	req.Header.Set("authorization", apiKey)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msgf("can not send unmute request")
		return err
	}

	if resp.StatusCode >= 400 {
		log.Error().Msgf("unmute request failed with status code %d", resp.StatusCode)
		return fmt.Errorf("unmute request failed with status code %d", resp.StatusCode)
	}

	return mongo.InsertUnMute(unmute)
}
