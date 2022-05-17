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

			time.Sleep(time.Hour)
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
	offset := 0
	for {

		users, err := coflnet.GetUsersFromId(offset)
		if err != nil {
			log.Error().Err(err).Msg("error getting users")
			return
		}

		if len(users) == 0 {
			return
		}

		for _, user := range users {

			err := refreshUser(user.UserId)
			if err != nil {
				log.Error().Err(err).Msgf("error loading user %d", user.UserId)
				continue
			}

			offset++
		}

	}
}

func refreshUser(id int) error {

	// wait to not get into api limit
	defer time.Sleep(time.Second * 20)
	log.Info().Msgf("loading user with id %d", id)

	u, err := coflnet.LoadUserById(id)
	if err != nil {
		return err
	}

	return discord.SetFlipperRoleForUser(u)
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
