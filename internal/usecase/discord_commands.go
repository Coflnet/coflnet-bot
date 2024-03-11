package usecase

import (
	chatgen "coflnet-bot/internal/gen/chat"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

type DiscordCommands struct {
	notificationChannelCommands []*discordgo.ApplicationCommand
	coflnetChannelCommands      []*discordgo.ApplicationCommand
	commandHandlers             map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	session                     *discordgo.Session

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
	var banPermissions int64 = discordgo.PermissionBanMembers

	d.notificationChannelCommands = []*discordgo.ApplicationCommand{
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

	d.coflnetChannelCommands = []*discordgo.ApplicationCommand{
		{
			Name:                     "mute",
			Description:              "Mute a user",
			DefaultMemberPermissions: &banPermissions,
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
					Description: "the message the user sent (why he got banned)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "either 'rule 1' or 'rule 2'",
					Required:    true,
				},
			},
		},
		{
			Name:                     "unmute",
			Description:              "Unmute a user",
			DefaultMemberPermissions: &banPermissions,
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
	for _, v := range d.notificationChannelCommands {
		cmd, err := d.session.ApplicationCommandCreate(d.session.State.User.ID, notificationServerGuildId(), v)
		if err != nil {
			slog.Error(fmt.Sprintf("Cannot create '%v' command: %v", v.Name, err))
			return err
		}

		slog.Info(fmt.Sprintf("Command '%s' added", cmd.Name))
	}

	for _, v := range d.coflnetChannelCommands {
		cmd, err := d.session.ApplicationCommandCreate(d.session.State.User.ID, coflnetServerGuildId(), v)
		if err != nil {
			slog.Error(fmt.Sprintf("Cannot create '%v' command: %v", v.Name, err))
			return err
		}

		slog.Info(fmt.Sprintf("Command '%s' added", cmd.Name))
	}

	return nil
}

func notificationServerGuildId() string {
	return mustEnv("NOTIFICATION_SERVER_GUILD_ID")
}

func coflnetServerGuildId() string {
	return mustEnv("GUILD_ID")
}
