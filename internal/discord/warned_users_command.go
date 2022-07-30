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

	user := optionMap["user"].Value.(*discordgo.User)

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
				Content:         "❌ there was a problem when searching warnings for user " + user.Username,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		err = SendMessageToDevLog(&DiscordMessageToSend{
			Message: "❌ there was a problem when listing warnings for user " + user.Username,
		})
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
		}
		return
	}

	// send success message
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("✅ Here is a list of warnings for %s\n%s", user.Username, formattedWarnings(warnings)),
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		},
	})
	err = SendMessageToDevLog(&DiscordMessageToSend{
		Message: "✅ User " + user.Username + " has been warned, this is his " + fmt.Sprintf("%d", len(warnings)) + " warning",
	})
	if err != nil {
		log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
	}
}

func formattedWarnings(warnings []*model.Warning) string {
	var result string
	for i, warning := range warnings {
		result += fmt.Sprintf("- %d. %s until %s\n", i+1, warning.Reason, warning.Until.Format("2006-01-02 15:04:05"))
	}
	return result
}
