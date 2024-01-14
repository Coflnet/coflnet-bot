package usecase

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"os"
)

type DiscordCommands struct {
	commands        []*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	session         *discordgo.Session

	clearAlertChannels *ClearAlertChannels
}

func NewDiscordCommands(session *discordgo.Session) *DiscordCommands {
	return &DiscordCommands{
		session:            session,
		clearAlertChannels: NewClearAlertChannels(session),
	}
}

func (d *DiscordCommands) defineCommands() {
	var defaultMemberPermissions int64 = discordgo.PermissionManageServer

	d.commands = []*discordgo.ApplicationCommand{
		{
			Name:                     "clear-alert-channel",
			Description:              "Clears a specific alert channel",
			DefaultMemberPermissions: &defaultMemberPermissions,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel-option",
					Description: "Channel option",
					ChannelTypes: []discordgo.ChannelType{
						discordgo.ChannelTypeGuildText,
						discordgo.ChannelTypeGuildVoice,
					},
					Required: true,
				},
			},
		},
	}

	d.commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"clear-alert-channel": d.clearAlertChannels.Execute,
	}
}

func (d *DiscordCommands) RegisterCommands() error {
	d.defineCommands()

	d.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := d.commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	slog.Info("Adding commands...")
	for _, v := range d.commands {
		cmd, err := d.session.ApplicationCommandCreate(d.session.State.User.ID, notificationServerGuildId(), v)
		if err != nil {
			slog.Error(fmt.Sprintf("Cannot create '%v' command: %v", v.Name, err))
			return err
		}

		slog.Info(fmt.Sprintf("Command '%s' added", cmd.Name))
	}

	return nil
}

func notificationServerGuildId() string {
	id, found := os.LookupEnv("NOTIFICATION_SERVER_GUILD_ID")
	if !found {
		panic("NOTIFICATION_SERVER_GUILD_ID not found")
	}

	slog.Debug("NOTIFICATION_SERVER_GUILD_ID found, returning")
	return id
}
