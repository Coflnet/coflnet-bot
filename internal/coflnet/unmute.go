package coflnet

import (
    "errors"
	"github.com/Coflnet/coflnet-bot/internal/model"
)

type UnMuteRequest struct {
	Unmuter string `json:"unMuter"`
	Reason  string `json:"reason"`
	Uuid    string `json:"uuid"`
}

func UnmutePlayer(unmute *model.Unmute) error {
    return errors.New("not implemented")
}
