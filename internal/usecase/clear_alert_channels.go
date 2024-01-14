package usecase

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"strings"
	"time"
)

type ClearAlertChannels struct {
	tracer  trace.Tracer
	session *discordgo.Session
}

func NewClearAlertChannels(session *discordgo.Session) *ClearAlertChannels {
	return &ClearAlertChannels{
		tracer:  otel.Tracer("clear-alert-channels"),
		session: session,
	}
}

func (c *ClearAlertChannels) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if opt, ok := optionMap["channel-option"]; ok {
		ch := opt.ChannelValue(nil)

		if !c.channelInClearAllowList(ch.ID) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Channel %s is not in the allow list", ch.Name),
				},
			})
			if err != nil {
				slog.Error("Cannot respond to interaction", "err", err)
			}
			return
		}

		go func() {
			ctx := context.Background()
			ctx, span := c.tracer.Start(ctx, "clear-messages-in-channel")
			defer span.End()

			err := c.clearMessagesInChannel(ctx, ch)
			if err != nil {
				slog.Error("Cannot clear messages in channel", "err", err)
			}
		}()

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Clearing messages in channel %s, this can take a while", ch.Name),
			},
		})

		if err != nil {
			slog.Error("Cannot respond to interaction", "err", err)
		}

		return
	}

	slog.Warn("No channel option found")
}

func (c *ClearAlertChannels) clearMessagesInChannel(ctx context.Context, channel *discordgo.Channel) error {
	ctx, span := c.tracer.Start(ctx, "clear-messages-in-channel")
	defer span.End()

	msgLimit := 100

	slog.Info(fmt.Sprintf("Clearing messages in channel %s", channel.Name))

	messages, err := c.session.ChannelMessages(channel.ID, msgLimit, "", "", "")
	if err != nil {
		slog.Error("Cannot get messages in channel", "err", err)
		span.RecordError(err)
		return err
	}

	messageIds := make([]string, len(messages))
	for i, message := range messages {
		messageIds[i] = message.ID
	}

	err = c.session.ChannelMessagesBulkDelete(channel.ID, messageIds)
	if err != nil {
		slog.Error("Cannot delete messages in channel", "err", err)
		span.RecordError(err)
		return err
	}
	slog.Info(fmt.Sprintf("deleted %d messages in channel %s", len(messages), channel.Name))

	if len(messages) == msgLimit {
		slog.Info(fmt.Sprintf("More than %d messages in channel %s, calling clearMessagesInChannel again", msgLimit, channel.Name))

		time.Sleep(1 * time.Minute)
		err = c.clearMessagesInChannel(ctx, channel)
		if err != nil {
			slog.Error("Cannot clear messages in channel", "err", err)
			span.RecordError(err)
			return err
		}
	}

	return nil
}

func (c *ClearAlertChannels) channelInClearAllowList(toCheck string) bool {
	channels := channelsThatAreAllowedToClear()
	for _, channel := range channels {
		if channel == toCheck {
			return true
		}
	}
	return false
}

func channelsThatAreAllowedToClear() []string {
	v, found := os.LookupEnv("NOTIFICATION_CHANNELS")
	if !found {
		panic("NOTIFICATION_CHANNELS not found")
	}

	return strings.Split(v, ",")
}
