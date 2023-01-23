package discord

import (
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func switchUsernameCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "switch-discord-username",
		Description: "Choose the preferred username, if you have multiple accounts",
	}
}

func switchUsernameHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	username := fmt.Sprintf("%s#%s", i.Member.User.Username, i.Member.User.Discriminator)
	log.Info().Msgf("searching for the user %s", username)
	user, err := mongo.GetUserByDiscordName(username)
	// if the user is not found, we can't do anything
	if err != nil {
		log.Error().Err(err).Msgf("error searching user %s", username)
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ your username is not registered",
			},
		})
		if err != nil {
			log.Error().Err(err).Msgf("Error sending message to user %s: %s", i.User.Username, err)
		}
		return
	}

	// show modal
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         "Choose your preferred username",
			Flags:           discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "switch_username_select",
							Placeholder: "Choose your preferred username",
							Options:     buildUsernameSelectOptions(user),
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

}

func switchUsernameSelected(s *discordgo.Session, i *discordgo.InteractionCreate) {

	username := fmt.Sprintf("%s#%s", i.Member.User.Username, i.Member.User.Discriminator)

  // get the user
  user, err := mongo.GetUserByDiscordName(username)
  if err != nil {
    log.Error().Err(err).Msgf("error searching user %s", username)
    err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData{
        Content: "❌ your username is not registered",
      },
    })

    if err != nil {
      log.Error().Err(err).Msgf("Error sending message to user %s: %s", i.User.Username, err)
    }
    return
  }

  user.PreferredUsername = i.MessageComponentData().Values[0]

  // mongo update
  err = mongo.UpdateUser(user)
  if err != nil {
    log.Error().Err(err).Msgf("error updating user %s", username)
    err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData{
        Content: "❌ error updating your username",
      },
    })

    if err != nil {
      log.Error().Err(err).Msgf("Error sending message to user %s: %s", i.User.Username, err)
    }

    return
  }

  s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
      Content: "You selected " + i.MessageComponentData().Values[0],
    },
  })
}

func buildUsernameSelectOptions(user *model.User) []discordgo.SelectMenuOption {
	options := []discordgo.SelectMenuOption{}
	for _, name := range user.DiscordNames {
		def := false
		if name == user.PreferredUsername {
			def = true
		}

		options = append(options, discordgo.SelectMenuOption{
			Label:   name,
			Value:   name,
			Default: def,
		})
	}
	return options
}
