package discord

import (
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

func SetFlipperRoleForUser(user *model.User) error {

	// set discord role at the end of the function
	defer func() {
		log.Info().Msgf("finished setting flipper role for user %d set role to %t", user.UserId, user.HasFlipperRole)
		err := mongo.SetFlipperRoleForUser(user)
		if err != nil {
			log.Error().Err(err).Msgf("error when saving user %d", user.UserId)
		}
	}()

	if user.DiscordNames == nil || len(user.DiscordNames) == 0 {
		log.Info().Msgf("user %d has no discord names, skip flipper role check", user.UserId)
		user.HasFlipperRole = false
		return nil
	}

	// TODO dont assume first discord name, that is not an empty string is the correct username
	// search the discord username for the user
	var discordName = ""
	for _, d := range user.DiscordNames {
		if d != "" {
			discordName = d
			break
		}
	}
	if discordName == "" {
		log.Info().Msgf("user %s has no discord name, skip flip role set", user.UserId)
		user.HasFlipperRole = false
		return nil
	}

	if !user.HasPremium() {
		log.Info().Msgf("remove role from user %s if he has the flipper role")
		user.HasFlipperRole = false

		return nil
	}

	log.Info().Msgf("set role for user %s", user.UserId)
	user.HasFlipperRole = true

	return nil
}
