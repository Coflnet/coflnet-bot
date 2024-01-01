package usecase

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

type DiscordMessageService struct {
	messageChannel chan *discordgo.Message
}

func NewDiscordMessageService() *DiscordMessageService {
	return &DiscordMessageService{
		messageChannel: make(chan *discordgo.Message),
	}
}

func (d *DiscordMessageService) Reader(ctx context.Context) (<-chan *discordgo.Message, error) {
	session, err := DiscordSession(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("Registering discord message handler")
	session.AddHandler(d.messageCreate)

	return d.messageChannel, nil
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func (d *DiscordMessageService) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example, but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	slog.Debug("Received a new message", "msg", m.Message.Content)

	d.messageChannel <- m.Message
}
