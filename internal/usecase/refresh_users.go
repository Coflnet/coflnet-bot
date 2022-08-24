package usecase

import (
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
	"time"
)

func StartRefresh() {
	for {
		warnings, err := mongo.ExpiredWarnings()
		if err != nil {
			log.Error().Err(err).Msg("error getting expired warnings")
			metrics.ErrorOccurred()
		}

		for warning := range warnings {
			err = discord.RemoveUserWarnedRole(warning.User)
			if err != nil {
				log.Error().Err(err).Msgf("error removing warned role for user %d", warning.User)
				metrics.ErrorOccurred()
			}

			time.Sleep(time.Second * 10)
		}
	}
}
