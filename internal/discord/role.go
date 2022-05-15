package discord

import (
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

func SetFlipperRoleForUser(user *mongo.User) error {

	if user.DiscordNames == nil || len(user.DiscordNames) == 0 {
		log.Info().Msgf("user %d has no discord names, skip flipper role check", user.UserId)
		return nil
	}

	if user.DiscordNames[0] == "" {

		log.Info().Msgf("user %s has no discord name, skip flip role set", user.UserId)
		return nil
	}

	if user.HasPremium() {

		log.Info().Msgf("remove role from user %s if he has the flipper role")
		user.HasFlipperRole = false
		mongo.SetFlipperRoleForUser(user)

		return nil
	}

	log.Info().Msgf("set role for user %s", user.UserId)
	user.HasFlipperRole = true
	mongo.SetFlipperRoleForUser(user)

	return nil
}
