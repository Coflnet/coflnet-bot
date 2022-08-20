package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"time"
)

func ingameMuteCommand() *discordgo.ApplicationCommand {

	var zero float64 = 0
	var modPerms int64 = discordgo.PermissionModerateMembers

	return &discordgo.ApplicationCommand{
		Name:                     "in-game-mute",
		Description:              "Mute a user in in-game-chat",
		DefaultMemberPermissions: &modPerms,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "username",
				Description: "The username of the user to mute",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the mute",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Message for the muted user",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "weeks",
				Description: "Mute for how many weeks",
				MinValue:    &zero,
				MaxValue:    52,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "days",
				Description: "Mute for how many days",
				MinValue:    &zero,
				MaxValue:    52,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "hours",
				Description: "Mute for how many hours",
				MinValue:    &zero,
				MaxValue:    24,
				Required:    false,
			},
		},
	}
}
func ingameMuteCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	metrics.InGameMuteTriggered()

	// check permissions
	isPermitted := isUserPermittedToMutePlayers(i.Member)
	if !isPermitted {
		log.Warn().Msgf("User %s is not permitted to mute players", i.Member.Nick)

		// the incident will not be reported
		// just the same message as
		// nonsudo is not in the sudoers file.  This incident will be reported.
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ you do not have permissions to mute players, this incident will be reported",
			},
		})
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to user %s: %s", i.User.Username, err)
		}

		err = SendMessageToMutesChannel(
			fmt.Sprintf("user %s tried to mute someone but did not have permissions", i.User.Username),
		)
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
			return
		}
		return
	}

	// get the parameters
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	username := fmt.Sprintf("%v", optionMap["username"].Value)
	muter := i.Member.User.Username
	reason := fmt.Sprintf("%v", optionMap["reason"].Value)

	message := ""
	if v, ok := optionMap["message"]; ok {
		message = fmt.Sprintf("%v", v.Value)
	}

	hours := 0
	if v, ok := optionMap["hours"]; ok {
		fHours := v.Value.(float64)
		hours = int(fHours)
	}

	days := 0
	if v, ok := optionMap["days"]; ok {
		fDays := v.Value.(float64)
		if fDays != 0 {
			days = int(fDays)
		}
	}

	weeks := 0
	if v, ok := optionMap["weeks"]; ok {
		fWeeks := v.Value.(float64)
		if fWeeks != 0 {
			weeks = int(fWeeks)
		}
	}

	// mute the user
	mute, err := muteCommand(username, muter, reason, message, hours, days, weeks)

	if err != nil {
		log.Error().Err(err).Msgf("failed to mute user")

		if e, ok := err.(*InvalidMuteUntilDurationError); ok {
			respErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content:         e.Error(),
					AllowedMentions: &discordgo.MessageAllowedMentions{},
					Flags:           discordgo.MessageFlagsEphemeral,
				},
			})

			if respErr != nil {
				log.Error().Err(respErr).Msgf("Error sending message to user %s: %s", i.User.Username, respErr)
			}
			return
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         fmt.Sprintf("❌ there was an error when muting %s please contact <@!256771988352139264>", username),
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
			},
		})

		if err != nil {
			log.Error().Err(err).Msgf("failed to respond to interaction")
		}

		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("🔇 user %s was muted until %v", username, mute.MuteUntil),
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Flags:           discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error in in-game-mute command")
	}

	err = SendMessageToMutesChannel(
		fmt.Sprintf("🔇 user %s was muted by %s for %s until %s", username, muter, reason, DiscordFormattedTime(mute.MuteUntil)),
	)
	if err != nil {
		log.Error().Err(err).Msgf("error in in-game-mute command")
	}
}

func muteCommand(username, muter, reason, message string, hours, days, weeks int) (*model.Mute, error) {

	muteUntil := time.Now()
	muteUntil = muteUntil.Add(time.Hour * time.Duration(hours))
	muteUntil = muteUntil.Add(time.Hour * 24 * time.Duration(days))
	muteUntil = muteUntil.Add(time.Hour * 24 * 7 * time.Duration(weeks))

	// check if time until is at least 50 minutes in the future
	if muteUntil.Before(time.Now().Add(time.Minute * 50)) {
		return nil, &InvalidMuteUntilDurationError{}
	}

	mute := model.Mute{
		Username:  username,
		Muter:     muter,
		MuteUntil: muteUntil,
		Reason:    reason,
	}

	// add message only if it is not empty
	if message != "" {
		mute.Message = message
	}

	err := coflnet.MutePlayer(&mute)
	if err != nil {
		return nil, err
	}

	return &mute, nil
}

func isUserPermittedToMutePlayers(u *discordgo.Member) bool {
	for _, role := range u.Roles {
		log.Info().Msgf("role: %v", role)
		if role == muteRole() {
			return true
		}
	}
	return false
}

type InvalidMuteUntilDurationError struct{}

func (err *InvalidMuteUntilDurationError) Error() string {
	return "mute time until must be at least 1 hour in the future"
}
