package usecase

import (
	"github.com/rs/zerolog/log"
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
)

func StartRefresh() {

  log.Info().Msgf("starting refresh all users")

  amount := 20 
  offset := 0

  for {
    users, err := coflnet.GetUsersFromId(amount, offset)
    
    if err != nil {
      log.Error().Err(err).Msgf("error getting users")
      return
    }

    if len(users) == 0 {
      break
    }

    for _, user := range users {
      log.Info().Msgf("refreshing user with id %d", user.UserId)

      _, err = coflnet.LoadUserById(user.UserId)

      if err != nil {
        log.Error().Err(err).Msgf("error loading user with id %d", user.UserId)
      }

    }

    offset += amount
  }

  log.Info().Msgf("finished refreshing all users")
}

