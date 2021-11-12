package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func Role(session *discordgo.Session, guildId string, roleId string) (*discordgo.Role, error) {

	defer session.Close()

	roles, err := session.GuildRoles(guildId)
	if err != nil {
		log.Error().Err(err).Msgf("error when getting guild roles in guild %s", guildId)
		return nil, err
	}

	var role *discordgo.Role
	for _, r := range roles {
		if r.ID == roleId {
			role = r
		}
	}
	if role == nil {
		log.Error().Msgf("Role with id %s not found", roleId)
		return nil, errors.New("role not found")
	}

	return role, nil
}
