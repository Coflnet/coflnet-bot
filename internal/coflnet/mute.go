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
	"time"
)

type MuteRequest struct {
	Uuid    string    `json:"uuid"`
	Muter   string    `json:"muter"`
	Reason  string    `json:"reason"`
	Expires time.Time `json:"expires"`
	Message string    `json:"message"`
}

// MutePlayer sends a mute request to the cofl chat service and saves the mute in the internal database
// TODO also sends the mute to the tfm server
func MutePlayer(mute *model.Mute) error {

	msg, err := mongo.FirstChatMessageOfUsername(mute.Username)
	if err != nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", mute.Username)
		return err
	}

	if msg == nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", mute.Username)
		return fmt.Errorf("could not find user %s in chat collection", mute.Username)
	}

	muteRequest := MuteRequest{
		Uuid:    msg.UUID,
		Muter:   mute.Muter,
		Reason:  mute.Reason,
		Expires: mute.MuteUntil,
		Message: mute.Message,
	}

	data, err := json.Marshal(muteRequest)
	if err != nil {
		log.Error().Err(err).Msgf("can not marshal mute request")
		return err
	}

	log.Info().Msgf("sending mute request: %s", string(data))

	apiKey := os.Getenv("COFL_CHAT_API_KEY")
	baseUrl := os.Getenv("COFL_CHAT_BASE_URL")
	p := "/api/chat/mute"
	url := fmt.Sprintf("%s%s", baseUrl, p)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		log.Error().Err(err).Msgf("can not create mute request")
		return err
	}

	req.Header.Set("authorization", apiKey)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msgf("can not send mute request")
		return err
	}

	if resp.StatusCode >= 400 {
		log.Error().Msgf("mute request failed with status code %d", resp.StatusCode)
		return fmt.Errorf("mute request failed with status code %d", resp.StatusCode)
	}

	return mongo.InsertMute(mute)
}
