package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
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
			log.Info().Msgf("user %s had flipper role before, skip it", discordNameForUser(u))
			return
		}

		if u.HasFlipperRole {
			err := giveUserFlipperRole(u)
			if err != nil {
				log.Error().Err(err).Msgf("error when giving flipper role to user %d", u.UserId)
				return
			}

			err = SendMessageToFlipperRoleChannel(fmt.Sprintf("give user %s the flipper role, he has premium until %v", discordNameForUser(u), u.PremiumUntil))
			if err != nil {
				log.Error().Err(err).Msgf("can not send message to dev chat")
				return
			}
		}
	}(user, userHadFlipperRoleBefore)

	if user.DiscordNames == nil || len(user.DiscordNames) == 0 {
		log.Info().Msgf("user %d has no discord names set, skip flipper role check", user.UserId)
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

func giveUserFlipperRole(user *model.User) error {

	guild := guildId()
	role := flipperRole()

	discordName := user.DiscordNames[0]
	discordNameWithoutTag := strings.Split(discordName, "#")[0]

	log.Info().Msgf("search api for the player with string %s", discordNameWithoutTag)
	members, err := session.GuildMembersSearch(guild, discordNameWithoutTag, 100)
	if err != nil {
		log.Error().Err(err).Msgf("error when searching for user %s", discordNameWithoutTag)
		return err
	}

	log.Info().Msgf("searched discord api for the player with string %s, found %d players", discordNameWithoutTag, len(members))

	if len(members) > 1 {
		log.Error().Msgf("found more than one player with the name %s", discordNameWithoutTag)
		return fmt.Errorf("found more than one player with the name %s", discordNameWithoutTag)
	}

	if len(members) == 0 {
		log.Error().Msgf("found no player with the name %s", discordNameWithoutTag)
		return fmt.Errorf("found no player with the name %s", discordNameWithoutTag)
	}

	log.Info().Msgf("set flipper role for user %s", members[0])
	return session.GuildMemberRoleAdd(guild, members[0].User.ID, role)
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

func warnedRole() string {
	r := os.Getenv("WARNED_ROLE")
	if r == "" {
		log.Panic().Msg("WARNED_ROLE is not set")
	}
	return r
}

func flipperRole() string {
	r := os.Getenv("DISCORD_FLIPPER_ROLE")
	if r == "" {
		log.Panic().Msg("DISCORD_FLIPPER_ROLE is not set")
	}
	return r
}

func guildId() string {
	g := os.Getenv("DISCORD_GUILD")
	if g == "" {
		log.Panic().Msg("DISCORD_GUILD is not set")
	}
	return g
}
