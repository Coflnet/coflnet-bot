package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

func SetFlipperRoleForUser(user *model.User) error {

	userHadFlipperRoleBefore := user.HasFlipperRole

	// set discord role at the end of the function
	defer func(u *model.User, hadFlipperRoleBefore bool) {
		log.Info().Msgf("finished setting flipper role for user %s set role to %t", discordNameForUser(u), u.HasFlipperRole)
		err := mongo.SetFlipperRoleForUser(u)
		if err != nil {
			log.Error().Err(err).Msgf("error when saving user %d", u.UserId)
		}

		if hadFlipperRoleBefore {
			return
		}

		if u.HasFlipperRole {
			err := SendMsgToDevChat(fmt.Sprintf("give user %s the flipper role, he has premium until %v", discordNameForUser(u), u.PremiumUntil))
			if err != nil {
				log.Error().Err(err).Msgf("can not send message to dev chat")
				return
			}
		}
	}(user, userHadFlipperRoleBefore)

	if user.DiscordNames == nil || len(user.DiscordNames) == 0 {
		log.Info().Msgf("user %d has no discord names, skip flipper role check", user.UserId)
		user.HasFlipperRole = false
		return nil
	}

	discordName := discordNameForUser(user)
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

// TODO dont assume first discord name, that is not an empty string is the correct username
// search the discord username for the user
func discordNameForUser(u *model.User) string {
	var discordName = ""
	for _, d := range u.DiscordNames {
		if d != "" {
			discordName = d
			break
		}
	}

	return discordName
}
