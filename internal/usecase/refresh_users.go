package usecase

import (
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

func StartRefresh() {
	for {
		warnings, err := mongo.ExpiredWarnings()
		if err != nil {
			log.Error().Err(err).Msg("error getting expired warnings")
			metrics.ErrorOccurred()
		}

		log.Info().Msgf("found %d expired warnings, remove role from users", len(warnings))

		for warning := range warnings {
			err = discord.RemoveUserWarnedRole(warning.User)
			if err != nil {
				log.Error().Err(err).Msgf("error removing warned role for user %d", warning.User)
				metrics.ErrorOccurred()
			}

			time.Sleep(time.Minute * 5)
		}

		time.Sleep(time.Hour * 1)
	}
}

func CheckIfUsersStillShouldHaveFlipperRole() {
	users, err := mongo.GetUsersWithFlipperRole()

	if err != nil {
		log.Error().Err(err).Msg("error getting users with flipper role")
		return
	}

	for user := range users {
		err = discord.SetFlipperRoleForUser(user)

		if err != nil {
			log.Error().Err(err).Msgf("there was an error when updating flipper role for user %d", user)
		}

		time.Sleep(time.Second * 5)
	}
}
