package coflnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
	"io"
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

type MuteResponse struct {
	UUID           string    `json:"uuid"`
	Muter          string    `json:"muter"`
	UnMuter        string    `json:"unMuter"`
	Message        string    `json:"message"`
	Reason         string    `json:"reason"`
	ClientID       int       `json:"clientId"`
	UnMuteClientID int       `json:"unMuteClientId"`
	Timestamp      time.Time `json:"timestamp"`
	Expires        time.Time `json:"expires"`
	Status         int       `json:"status"`
}

// MutePlayer sends a mute request to the cofl chat service and saves the mute in the internal database
func MutePlayer(mute *model.Mute) (time.Time, error) {

	msg, err := mongo.FirstChatMessageOfUsername(mute.Username)
	if err != nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", mute.Username)
		return time.Time{}, err
	}

	if msg == nil {
		log.Error().Err(err).Msgf("could not find user %s in chat collection", mute.Username)
		return time.Now(), fmt.Errorf("could not find user %s in chat collection", mute.Username)
	}

	muteRequest := MuteRequest{
		Uuid:    msg.UUID,
		Muter:   mute.Muter,
		Reason:  mute.Reason,
		Expires: mute.MuteUntil,
		Message: mute.Message,
	}

	result, err := sendMuteToAPI(muteRequest)
	if err != nil {
		return time.Time{}, err
	}

	// use the muteuntil from the api response
	mute.MuteUntil = result

	if err = mongo.InsertMute(mute); err != nil {
		log.Error().Err(err).Msgf("can not save mute in db")
		return time.Time{}, err
	}

	return result, nil
}

func sendMuteToAPI(muteRequest MuteRequest) (time.Time, error) {

	data, err := json.Marshal(muteRequest)
	if err != nil {
		log.Error().Err(err).Msgf("can not marshal mute request")
		return time.Time{}, err
	}

	log.Info().Msgf("sending mute request: %s", string(data))

	apiKey := os.Getenv("COFL_CHAT_API_KEY")
	baseUrl := os.Getenv("COFL_CHAT_BASE_URL")
	p := "/api/chat/mute"
	url := fmt.Sprintf("%s%s", baseUrl, p)

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		log.Error().Err(err).Msgf("can not create mute request")
		return time.Time{}, err
	}

	req.Header.Set("authorization", apiKey)
	req.Header.Set("content-type", "application/json")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msgf("can not send mute request")
		return time.Time{}, err
	}

	if resp.StatusCode >= 400 {
		log.Error().Msgf("mute request failed with status code %d", resp.StatusCode)
		return time.Time{}, fmt.Errorf("mute request failed with status code %d", resp.StatusCode)
	}

	var b []byte
	b, err = io.ReadAll(resp.Body)

	var muteResponse MuteResponse
	err = json.Unmarshal(b, &muteResponse)
	if err != nil {
		log.Error().Err(err).Msgf("can not decode mute response; response: " + string(b))
		return time.Time{}, err
	}

	log.Info().Msgf("mute response from api: %v", muteResponse)
	return muteResponse.Expires, nil
}
