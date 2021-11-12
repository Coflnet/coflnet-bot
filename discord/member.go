package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func DiscordMember(session *discordgo.Session, discordName string, guildId string) (*discordgo.Member, error) {

	defer session.Close()

	lastReturn := ""
	members := []*discordgo.Member{}
	for {
		list, err := session.GuildMembers(guildId, lastReturn, 1000)
		if err != nil {
			log.Error().Err(err).Msgf("error when loading users in guild %s", guildId)
			return nil, err
		}

		members = append(members, list...)
		if len(list) < 1000 {
			break
		}

		lastReturn = members[len(members)-1].User.ID
	}

	for _, member := range members {
		if member.User.Username == discordName {
			return member, nil
		}
	}

	return nil, fmt.Errorf("no user found with username %s", discordName)
}
