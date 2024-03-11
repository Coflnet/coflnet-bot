package usecase

import (
	chatgen "coflnet-bot/internal/gen/chat"
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
	muteCommand        *MuteCommand
	unmuteCommand      *UnmuteCommand
}

func NewDiscordCommands(session *discordgo.Session, userService *UserService, chatClient *chatgen.Client) *DiscordCommands {
	return &DiscordCommands{
		session:            session,
		clearAlertChannels: NewClearAlertChannels(session),
		muteCommand:        NewMuteCommand(session, userService, chatClient),
		unmuteCommand:      NewUnmuteCommand(session, userService, chatClient),
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
		{
			Name:        "mute",
			Description: "Mute a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "User to mute",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "message",
					Required:    true,
				},
			},
		},
		{
			Name:        "unmute",
			Description: "Unmute a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "User to unmute",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "reason",
					Required:    true,
				},
			},
		},
	}

	d.commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"clear-alert-channel": d.clearAlertChannels.Execute,
		"mute":                d.muteCommand.Execute,
		"unmute":              d.unmuteCommand.Execute,
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
