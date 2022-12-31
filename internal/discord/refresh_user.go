package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
)

func refreshUser() *discordgo.ApplicationCommand {
	var writePerm int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "refresh-user",
		Description:              "Sync your coflnet user with the hypixel api",
		DefaultMemberPermissions: &writePerm,
		Options: []*discordgo.ApplicationCommandOption{
		{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user",
				Description: "The user to check",
				Required:    true,
			},
		},
	}
}

func refreshUserHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	userId := optionMap["user"].Value.(string)
  fmt.Println(userId)

  uuid, err := coflnet.PlayerUUIDByName(userId)
  if err != nil {
    log.Error().Err(err).Msg("Error getting uuid")
    log.Error().Err(err).Msg("Error getting user")
	  _ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		  Type: discordgo.InteractionResponseChannelMessageWithSource,
		  Data: &discordgo.InteractionResponseData{
		  	Content:         "Sync failed",
			  AllowedMentions: &discordgo.MessageAllowedMentions{},
			  Flags:           discordgo.MessageFlagsEphemeral,
		  },
    })
    return
  }

  _, err = coflnet.LoadUserByUUID(uuid)
  if err != nil {
    log.Error().Err(err).Msg("Error loading user")
    log.Error().Err(err).Msg("Error getting user")
	  _ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		  Type: discordgo.InteractionResponseChannelMessageWithSource,
		  Data: &discordgo.InteractionResponseData{
		  	Content:         "Sync failed",
			  AllowedMentions: &discordgo.MessageAllowedMentions{},
			  Flags:           discordgo.MessageFlagsEphemeral,
		  },
    })
    return
  }

	// send success message
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         "User synced",
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Flags:           discordgo.MessageFlagsEphemeral,
		},
	})
}

func RefreshUserByPlayername(name string) error {
  uuid, err := coflnet.PlayerUUIDByName(name)   
  if err != nil {
    log.Error().Err(err).Msgf("can not get uuid for %s", name)
    return err
  }

  _, err = coflnet.LoadUserByUUID(uuid)
  if err != nil {
    log.Error().Err(err).Msgf("can not load user for %s", name)
    return err
  }

  return nil
}

