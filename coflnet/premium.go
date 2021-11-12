package coflnet

import (
	"time"

	"github.com/Coflnet/coflnet-bot/database"
	"github.com/rs/zerolog/log"
)

func CheckIfUserHasPremium(user *database.User) (bool, error) {
	log.Info().Msgf("%v", user)

	premiumExpires, err := convertedDate(user.PremiumExpires)
	if err != nil {
		log.Error().Err(err).Msgf("error when parsing premiumexpires for user %s", user.Id)
		return false, err
	}

	if premiumExpires.After(time.Now()) {
		return true, nil
	}
	return false, nil
}

func convertedDate(val string) (time.Time, error) {
	layout := "0001-01-01T00:00:00Z"
	return time.Parse(layout, val)
}
