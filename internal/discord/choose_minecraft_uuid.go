package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
)

func switchMinecraftAccountCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "switch-minecraft-account",
		Description: "Choose the preferred minecraft account, if you have multiple accounts",
	}
}

func switchMinecraftUUIDHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	username := fmt.Sprintf("%s#%s", i.Member.User.Username, i.Member.User.Discriminator)
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
			Content: "Choose your preferred minecraft account",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "switch_minecraft_uuid_select",
							Placeholder: "Choose your preferred username",
							Options:     buildMinecraftUuidSelectOptions(user),
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

func switchMinecraftSelected(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

	user.PreferredUUID = i.MessageComponentData().Values[0]

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

	uuid := i.MessageComponentData().Values[0]
	name, err := coflnet.PlayerName(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("error getting name for uuid %s", uuid)

		// send as ephemeral
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				AllowedMentions: &discordgo.MessageAllowedMentions{},
				Flags:           discordgo.MessageFlagsEphemeral,
				Content:         "❌ error updating your username",
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         "You selected " + name,
			Flags:           discordgo.MessageFlagsEphemeral,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		},
	})
}

func buildMinecraftUuidSelectOptions(user *model.User) []discordgo.SelectMenuOption {
	uuidUsernameMap := make(map[string]string)
	for _, uuid := range user.MinecraftUuids {

		name, err := coflnet.PlayerName(uuid)
		if err != nil {
			log.Error().Err(err).Msgf("error getting name for uuid %s", uuid)
			continue
		}

		uuidUsernameMap[uuid] = name
	}

	options := []discordgo.SelectMenuOption{}
	for uuid, name := range uuidUsernameMap {

		def := false
		if name == user.PreferredUsername {
			def = true
		}

		options = append(options, discordgo.SelectMenuOption{
			Label:   name,
			Value:   uuid,
			Default: def,
		})
	}
	return options
}
