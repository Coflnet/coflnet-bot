package usecase

import (
	chatgen "coflnet-bot/internal/gen/chat"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type MuteCommand struct {
	tracer      trace.Tracer
	session     *discordgo.Session
	userService *UserService
	chatClient  *chatgen.Client
}

func NewMuteCommand(session *discordgo.Session, userService *UserService, chatClient *chatgen.Client) *MuteCommand {
	return &MuteCommand{
		session:     session,
		userService: userService,
		chatClient:  chatClient,
		tracer:      otel.Tracer("mute-command"),
	}
}

func (m *MuteCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ctx, span := m.tracer.Start(ctx, "mute-command")
	defer span.End()

	slog.Info(fmt.Sprintf("mute user with trace id %s", span.SpanContext().TraceID()))

	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	opt, ok := optionMap["user"]
	if !ok {
		span.AddEvent("User is required")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "User is required",
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}
	messageOpt, ok := optionMap["message"]
	if !ok {
		span.AddEvent("Message is required")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Message is required",
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return

	}

	if !userAllowedToMute(i.Member) {
		span.AddEvent("User not allowed to mute")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are not allowed to mute",
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

	username := opt.StringValue()
	span.SetAttributes(attribute.String("username", username))
	message := messageOpt.StringValue()
	span.SetAttributes(attribute.String("message", message))

	uuid, err := m.userService.LoadUserUUIDByMinecraftName(ctx, username)
	if err != nil {
		span.RecordError(err)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Cannot find user %s", username),
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

	muter := i.Member.User.Username

	// Mute user
	err = m.muteUser(ctx, uuid, message, muter)
	if err != nil {
		span.RecordError(err)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Cannot mute user, error: %s", span.SpanContext().TraceID()),
			},
		})

		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
	}
}

func (m *MuteCommand) muteUser(ctx context.Context, uuid, message, muter string) error {
	ctx, span := m.tracer.Start(ctx, "mute-user")
	defer span.End()

	body := chatgen.PostApiChatMuteJSONRequestBody{
		Message: strPtr(message),
		Muter:   strPtr(muter),
		Reason:  strPtr(message),
		Uuid:    strPtr(uuid),
	}
	params := chatgen.PostApiChatMuteParams{
		Authorization: strPtr(authorization()),
	}

	_, err := m.chatClient.PostApiChatMuteWithApplicationWildcardPlusJSONBody(ctx, &params, body)
	if err != nil {
		span.RecordError(err)
		slog.Error("Cannot mute user", "err", err)
		return err
	}

	return nil
}

func userAllowedToMute(member *discordgo.Member) bool {
	roles := member.Roles

	for _, role := range roles {
		if roleAllowedToMute(role) {
			return true
		}

	}

	return false
}

func roleAllowedToMute(role string) bool {
	modRole := mustEnv("MOD_ROLE")

	if role == modRole {
		return true
	}

	return false
}
