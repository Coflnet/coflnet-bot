package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var registeredCommands []*discordgo.ApplicationCommand

func registerCommands() error {
	commands := []*discordgo.ApplicationCommand{
		ingameMuteCommand(),
		ingameUnmuteCommand(),
		warnCommand(),
		userWarnings(),
		checkFlipperRole(),
		refreshUser(),
		switchUsernameCommand(),
		switchMinecraftAccountCommand(),
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"in-game-mute":                 ingameMuteCommandHandler,
		"in-game-unmute":               ingameUnmuteCommandHandler,
		"user-warnings":                userWarningsHandler,
		"warn-user":                    warnUserHandler,
		"check-flipper-role":           checkFlipperRoleHandler,
		"refresh-user":                 refreshUserHandler,
		"switch-discord-username":      switchUsernameHandler,
		"switch_username_select":       switchUsernameSelected,
		"switch-minecraft-account":     switchMinecraftUUIDHandler,
		"switch_minecraft_uuid_select": switchMinecraftSelected,
	}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := commandHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
	for _, v := range commands {
		registeredCommand, err := session.ApplicationCommandCreate(session.State.User.ID, os.Getenv("DISCORD_GUILD"), v)
		if err != nil {
			log.Error().Err(err).Msgf("Cannot create '%v' command: %v", v.Name, err)
			return err
		}

		log.Info().Msgf("Created command '%v'", v.Name)
		registeredCommands = append(registeredCommands, registeredCommand)
	}

	return nil
}

func unregisterCommands() error {
	log.Info().Msgf("unregistering %d commands", len(registeredCommands))
	for _, c := range registeredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, discordGuildId(), c.ID)
		if err != nil {
			log.Error().Err(err).Msgf("Cannot delete '%v' command: %v", c.Name, err)
			return err
		}
	}

	log.Info().Msgf("unregistered %d commands", len(registeredCommands))
	return nil
}
