package usecase

import (
	"time"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

func StartRefresh() {
	go func() {
		for {
			refreshUsers()
		}
	}()

	go func() {
		for {
			log.Info().Msgf("start check of all users with flipper role")
			CheckIfUsersStillShouldHaveFlipperRole()
			log.Info().Msgf("finished check of all users with flipper role")

			time.Sleep(time.Hour * 2)
		}
	}()
}

func refreshUsers() {
	id := 1
	errorCounter := 0

	// refresh users, when 100 errors in a row happen return
	for errorCounter < 100 {

		err := refreshUser(id)
		if err != nil {
			errorCounter++
		} else {
			errorCounter = 0
		}

		if id%100 == 0 {
			log.Info().Msgf("refreshed %d users", id)
		}

		id++

		if id%10 == 0 {
			time.Sleep(time.Second * 15)
		}
	}
}

func refreshUser(id int) error {
	u, err := coflnet.LoadUserById(id)
	if err != nil {
		return err
	}

	discord.SetFlipperRoleForUser(u)

	return nil
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
			log.Error().Err(err).Msgf("there was an error when updating flipper role for user %d", user.UserId)
		}

		time.Sleep(time.Second * 5)
	}
}
