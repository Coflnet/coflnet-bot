package coflnet

import (
	"errors"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/model"
)

type MuteRequest struct {
	Uuid    string    `json:"uuid"`
	Muter   string    `json:"muter"`
	Reason  string    `json:"reason"`
	Expires time.Time `json:"expires"`
	Message string    `json:"message"`
}

type MuteResponse struct {
	UUID           string `json:"uuid"`
	Muter          string `json:"muter"`
	UnMuter        string `json:"unMuter"`
	Message        string `json:"message"`
	Reason         string `json:"reason"`
	ClientID       int    `json:"clientId"`
	UnMuteClientID int    `json:"unMuteClientId"`
	Timestamp      string `json:"timestamp"`
	Expires        string `json:"expires"`
	Status         int    `json:"status"`
}

// MutePlayer sends a mute request to the cofl chat service and saves the mute in the internal database
func MutePlayer(mute *model.Mute) (time.Time, error) {
    return time.Now(), errors.New("not implemented")
}

func sendMuteToAPI(muteRequest MuteRequest) (time.Time, error) {
    return time.Now(), errors.New("not implemented")
}
