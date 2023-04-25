package discord

import (
	"context"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/exp/slog"
)

func (d *DiscordHandler) RegisterCommands() error {
	time.Sleep(time.Second * 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	ctx, span := d.tracer.Start(ctx, "register-commands")
	defer span.End()

	customCommands := []DiscordCommand{
		d.muteCommand,
		d.unmuteCommand,
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
		_, span := d.tracer.Start(ctx, "register-command")
		defer span.End()
		span.SetAttributes(attribute.String("command-name", v.Name))

		slog.Info(fmt.Sprintf("registering command: %s", v.Name))
		cmd, err := d.session.ApplicationCommandCreate(d.session.State.User.ID, utils.DiscordGuildId(), v)
		if err != nil {
			span.RecordError(err)
			slog.Error("failed to register command", err)
			return err
		}

		registeredCommands[i] = cmd
		slog.Info(fmt.Sprintf("registered command: %s", cmd.Name))
		span.SetAttributes(attribute.String("command-id", cmd.ID))
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
