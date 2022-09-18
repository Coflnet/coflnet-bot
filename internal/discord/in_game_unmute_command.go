package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func ingameUnmuteCommand() *discordgo.ApplicationCommand {

	var writePerm int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "in-game-unmute",
		Description:              "Unmute a user in in-game-chat",
		DefaultMemberPermissions: &writePerm,
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
				Name:        "additional-information",
				Description: "Additional information for the mute, e.g. link to the message",
				Required:    false,
			},
		},
	}
}
func ingameUnmuteCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	metrics.InGameUnmuteTriggered()

	// check permissions
	isPermitted := isUserPermittedToMutePlayers(i.Member)
	if !isPermitted {
		log.Warn().Msgf("User %s is not permitted to unmute players", i.Member.Nick)

		// the incident will not be reported
		// just the same message as
		// nonsudo is not in the sudoers file.  This incident will be reported.
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "you do not have permissions to unmute players, this incident will be reported",
				Flags:           discordgo.MessageFlagsEphemeral,
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to user %s: %s", i.User.Username, err)
			metrics.ErrorOccurred()
		}

		err = SendMessageToMutesChannel(
			fmt.Sprintf("user %s tried to unmute someone but did not have permissions", i.User.Username),
		)
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to dev log: %s", err)
			metrics.ErrorOccurred()
		}
	}

	// get the parameters
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	username := fmt.Sprintf("%v", optionMap["username"].Value)
	unmuter := i.Member.User.Username
	reason := fmt.Sprintf("%v", optionMap["reason"].Value)
	additionalInformation := fmt.Sprintf("%v", optionMap["additionalInformation"].Value)

	// mute the user
	_, err := unmuteCommand(username, unmuter, reason, additionalInformation)

	if err != nil {
		log.Error().Err(err).Msgf("❌ failed to unmute user")
		metrics.ErrorOccurred()

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
				log.Error().Err(respErr).Msgf("❌ Error sending message to user %s: %s", i.User.Username, respErr)
			}
			return
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("❌there was an error when unmuting %s please contact <@!256771988352139264>", username),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})

		if err != nil {
			log.Error().Err(err).Msgf("❌ failed to respond to interaction")
		}

		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         fmt.Sprintf("🔈 user %s was unmuted by %s", username, unmuter),
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Flags:           discordgo.MessageFlagsEphemeral,
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error in in-game-unmute command")
	}

	err = SendMessageToMutesChannel(
		fmt.Sprintf("🔈 user %s was unmuted by %s for %s", username, unmuter, reason),
	)

	if err != nil {
		log.Error().Err(err).Msgf("error in in-game-unmute command")
	}
}

func unmuteCommand(username, unmuter, reason, additionalInformation string) (*model.Unmute, error) {

	unmute := model.Unmute{
		Username:              username,
		Reason:                reason,
		Unmuter:               unmuter,
		AdditionalInformation: additionalInformation,
	}

	err := coflnet.UnmutePlayer(&unmute)
	if err != nil {
		return nil, err
	}

	return &unmute, nil
}
