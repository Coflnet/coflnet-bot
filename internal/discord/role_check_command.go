package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func checkFlipperRole() *discordgo.ApplicationCommand {
	var writePerm int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "check-flipper-role",
		Description:              "Check if the user should have the flipper role",
		DefaultMemberPermissions: &writePerm,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to check",
				Required:    true,
			},
		},
	}
}

func checkFlipperRoleHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	userId := optionMap["user"].Value.(string)
	user, err := s.GuildMember(i.GuildID, userId)
	if err != nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌did not found user with id " + userId,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	if user == nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ no user provided",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	if !isUserPermittedToWarnPlayers(i.Member) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ You don't have permission to list warnings of an user",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	tag := user.User.Username + "#" + user.User.Discriminator
	dbUser, err := mongo.SearchByDiscordTag(tag)

	if err != nil || dbUser == nil {
		log.Error().Err(err).Msgf("searched user with tag %s but failed", tag)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ failed to search user with tag " + tag,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	beforeUpdate := dbUser.HasFlipperRole
	err = SetFlipperRoleForUser(dbUser)
	if err != nil {
		log.Error().Err(err).Msgf("failed to set flipper role for user with tag %s", tag)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ failed to update flipper role for user with tag " + tag,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	if beforeUpdate == dbUser.HasFlipperRole {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "⚠️ user flipper role has not changed",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	text := fmt.Sprintf("✅ user with tag %s has been updated, he has now the flipper role", tag)
	if !dbUser.HasFlipperRole {
		text = fmt.Sprintf("✅ user with tag %s has been updated, he has not the flipper role now", tag)
	}

	// send success message
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         text,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Flags:           discordgo.MessageFlagsEphemeral,
		},
	})
}
