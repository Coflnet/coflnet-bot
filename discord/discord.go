package discord

import (
	"errors"
	"os"

	"github.com/Coflnet/coflnet-bot/coflnet"
	"github.com/Coflnet/coflnet-bot/database"
	"github.com/Coflnet/coflnet-bot/hypixel"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var userRolesChanged = 0
var dryRunSkips = 0

func UserRole(session *discordgo.Session, user *database.User) (*discordgo.Role, error) {

	roleId := os.Getenv("DISCORD_ROLE")
	guildId := os.Getenv("DISCORD_GUILD")

	// load discord name from hypixel
	err := hypixel.User(user)
	if err != nil {
		log.Error().Err(err).Msgf("error getting hypixel data for user %d", user.Id)
		return nil, err
	}
	if user.DiscordName == "" {
		log.Warn().Msgf("can not update user role, no discord name is given")
		return nil, errors.New("no discord name given")
	}

	// load discord member to check roles
	userName := user.DiscordName
	member, err := DiscordMember(session, userName, guildId)
	if err != nil {
		log.Error().Err(err).Msgf("could not get discord member %s in guild %s", userName, guildId)
		return nil, err
	}

	// load role
	role, err := Role(session, guildId, roleId)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error when loading the role %s in guild %s", guildId, roleId)
		return nil, err
	}

	// add role if user has it
	hasRole := hasMemberRole(member, role)
	if hasRole {
		user.Role = role.Name
	}

	return nil, nil
}

func UpdateUserRole(userId int, session *discordgo.Session) error {
	roleId := os.Getenv("DISCORD_ROLE")
	guildId := os.Getenv("DISCORD_GUILD")

	// getUser
	user, err := database.UserById(userId)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user %d", userId)
		return err
	}

	// check if user has premium
	hasPremium, err := coflnet.CheckIfUserHasPremium(&user)
	if err != nil {
		log.Error().Err(err).Msgf("error when checking if user has premium %v", user)
		return err
	}
	if !hasPremium {
		log.Warn().Msgf("the user %v has no premium, so no role will be given", user)
	}

	// get role
	role, err := Role(session, guildId, roleId)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error, when loading the role %s in guild %s", roleId, guildId)
		return err
	}

	// load discord name from hypixel
	err = hypixel.User(&user)
	if err != nil {
		log.Error().Err(err).Msgf("error getting hypixel data for user %d", userId)
		return err
	}
	if user.DiscordName == "" {
		log.Warn().Msgf("can not update user role, no discord name is given")
		return nil
	}

	// load user name
	userName := user.DiscordName
	member, err := DiscordMember(session, userName, guildId)
	if err != nil {
		log.Error().Err(err).Msgf("could not get discord member %s in guild %s", userName, guildId)
		return err
	}

	// check if dry run is enabled
	if os.Getenv("DRY_RUN") == "true" {
		dryRunSkips++
		log.Warn().Msgf("did not add role %s to member %s, because dry run is enabled, skip %d", role.ID, user.Id, dryRunSkips)
		return nil
	}

	// add role to user
	err = session.GuildMemberRoleAdd(guildId, member.User.ID, role.ID)
	if err != nil {
		log.Error().Err(err).Msgf("error when addding role %s to user %s in guild %s", role.Name, member.User.Username, guildId)
		return err
	}
	userRolesChanged++
	log.Info().Msgf("gave the user %s the role %s in guild %s, users changed: %d", member.User.Username, role.Name, guildId, userRolesChanged)

	return nil
}

func hasMemberRole(member *discordgo.Member, role *discordgo.Role) bool {
	for _, r := range member.Roles {
		if r == role.ID {
			return true
		}
	}
	return false
}
