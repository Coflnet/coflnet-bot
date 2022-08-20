package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func userWarnings() *discordgo.ApplicationCommand {
	var writePerm int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "user-warnings",
		Description:              "List of Warnings for one user",
		DefaultMemberPermissions: &writePerm,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to warn",
				Required:    true,
			},
		},
	}
}

func userWarningsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
			},
		})
		return
	}

	warnings, err := mongo.WarningsByUser(user)
	if err != nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ there was a problem when searching warnings for user " + username(user),
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		err = SendMessageToWarningsChannel(fmt.Sprintf("❌ there was a problem when listing warnings for user %s", username(user)))
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
		}
		return
	}

	text := fmt.Sprintf("✅ Here is a list of warnings for %s\n%s", username(user), formattedWarnings(warnings))
	if len(warnings) == 0 {
		text = fmt.Sprintf("✅ There are no warnings for %s", username(user))
	}

	// send success message
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         text,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		},
	})
}

func formattedWarnings(warnings []*model.Warning) string {
	var result string
	for i, warning := range warnings {
		result += fmt.Sprintf("- %d. %s until %s\n", i+1, warning.Reason, fmt.Sprintf("<t:%d:f>", warning.Until.Unix()))
	}
	return result
}

func username(user *discordgo.Member) string {
	if user.Nick != "" {
		return user.Nick
	}
	return user.User.Username
}
