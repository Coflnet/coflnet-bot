package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func warnCommand() *discordgo.ApplicationCommand {
	var zero float64 = 0
	var writePerm int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "warn-user",
		Description:              "Warn a user",
		DefaultMemberPermissions: &writePerm,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to warn",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "days",
				Description: "The amount of days the user is warned for",
				Required:    true,
				MinValue:    &zero,
				MaxValue:    365,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the warning",
				Required:    true,
			},
		},
	}
}

func warnUserHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	userId := optionMap["user"].Value.(string)
	days := optionMap["days"].Value.(float64)
	reason := optionMap["reason"].Value.(string)

	user, err := s.GuildMember(i.GuildID, userId)
	if err != nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         fmt.Sprintf("❌ Did not found user %s in guild", userId),
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		return
	}

	if reason == "" {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ You must provide a reason for the warning",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		log.Warn().Msgf("reason is empty")
		return
	}

	if days < 0 || days > 365 {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ Days must be between 0 and 365",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		log.Warn().Msgf("days input %f is not between 0 and 365", days)
		return
	}

	if !isUserPermittedToWarnPlayers(i.Member) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ You don't have permission to warn users",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		log.Warn().Msgf("user %s is not permitted to warn players", i.Member.User.Username)
		return
	}

	warnings, err := warnUser(user, reason, i.Member, int(days))
	if err != nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "❌ there was a problem when warning user " + username(user),
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		log.Error().Err(err).Msgf("error warning user")
		err = SendMessageToDevLog(&DiscordMessageToSend{
			Message: "❌ there was a problem when warning user " + username(user),
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
			Content:         "✅ User " + username(user) + " has been warned, this is his " + fmt.Sprintf("%d", warnings) + " warning",
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		},
	})
	err = SendMessageToDevLog(&DiscordMessageToSend{
		Message: "✅ User " + username(user) + " has been warned, this is his " + fmt.Sprintf("%d", warnings) + " warning",
	})
	if err != nil {
		log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
	}
}

func warnUser(user *discordgo.Member, reason string, muter *discordgo.Member, days int) (int, error) {
	warning := &model.Warning{
		User:   user,
		Warner: muter,
		Until:  time.Now().Add(time.Duration(days) * 24 * time.Hour),
		Reason: reason,
	}

	err := mongo.InsertWarning(warning)
	if err != nil {
		return 0, err
	}

	warnings, err := mongo.WarningsByUser(user)
	if err != nil {
		return 0, err
	}

	return len(warnings), nil
}

func isUserPermittedToWarnPlayers(user *discordgo.Member) bool {
	for _, role := range user.Roles {
		if role == helperRole() || role == devRole() || role == modRole() {
			return true
		}
	}

	return false
}

func helperRole() string {
	r := os.Getenv("HELPER_ROLE")
	if r == "" {
		log.Panic().Msg("HELPER_ROLE is not set")
	}
	return r
}

func devRole() string {
	r := os.Getenv("DEV_ROLE")
	if r == "" {
		log.Panic().Msg("DEV_ROLE is not set")
	}
	return r
}

func modRole() string {
	r := os.Getenv("MOD_ROLE")
	if r == "" {
		log.Panic().Msg("MOD_ROLE is not set")
	}
	return r
}
