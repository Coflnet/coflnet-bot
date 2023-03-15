package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type DiscordCommand interface {
	Name() string
	CreateCommand() *discordgo.ApplicationCommand
	HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type baseCommand struct {
    tracer trace.Tracer
}

func newBaseCommand() *baseCommand {
    return &baseCommand{
        tracer: otel.Tracer("-command"),
    }
}

func (b *baseCommand) parseResponseOptions(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
    // Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

    return optionMap
}

func (b *baseCommand) createFollowupMessage(ctx context.Context, content string, s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.Message, error) {
    _, span := b.tracer.Start(ctx, "create-followup-message")
    defer span.End()

    msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: content,
	})

    if err != nil {
        span.RecordError(err)
    }

    return msg, err
}

func (b *baseCommand) editFollowupMessage(ctx context.Context, content, initialFollowupMessageId string, s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.Message, error) {
    return s.FollowupMessageEdit(i.Interaction, initialFollowupMessageId, &discordgo.WebhookEdit{
		Content: &content,
	})
}

func (b *baseCommand) requestReceived(ctx context.Context, s *discordgo.Session, i *discordgo.InteractionCreate) error {
    _, span := b.tracer.Start(ctx, "send-first-response")
    defer span.End()

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "⏳ request received",
		},
	})
}
