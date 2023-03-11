package discord

import (
	"fmt"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slog"
)

type DiscordCommand interface {
	Name() string
	CreateCommand() *discordgo.ApplicationCommand
	HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (d *DiscordHandler) RegisterCommands() error {

	customCommands := []DiscordCommand{
		CreateMuteCommand(),
	}

	// convert the custom commands into discordgo commands
	commands := make([]*discordgo.ApplicationCommand, 0)
	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	for _, v := range customCommands {
		commands = append(commands, v.CreateCommand())
		commandHandlers[v.Name()] = v.HandleCommand
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := d.session.ApplicationCommandCreate(d.session.State.User.ID, utils.DiscordGuildId(), v)
		if err != nil {
			slog.Error("failed to register command", err)
			panic(err)
		}
		registeredCommands[i] = cmd
		slog.Info(fmt.Sprintf("registered command %s", cmd.Name))
	}

	d.commands = registeredCommands

	// adding stuff to actually handle the commands
	d.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		slog.Info(fmt.Sprintf("got interaction: %s", i.ApplicationCommandData().Name))
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	return nil
}

func (d *DiscordHandler) unregisterCommands() {
	for _, v := range d.commands {
		err := d.session.ApplicationCommandDelete(d.session.State.User.ID, utils.DiscordGuildId(), v.ID)
		if err != nil {
			slog.Error("failed to unregister command", err)
		}
		slog.Info(fmt.Sprintf("unregistered command %s", v.Name))
	}
}
