package usecase

import (
	chatgen "coflnet-bot/internal/gen/chat"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log/slog"
	"strings"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
	opt, ok := optionMap["username"]
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
	reasonOpt, ok := optionMap["reason"]
	if !ok {
		span.AddEvent("Reason is required")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Reason is required",
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

	span.AddEvent("User allowed to mute")

	username := opt.StringValue()
	span.SetAttributes(attribute.String("username", username))
	message := messageOpt.StringValue()
	span.SetAttributes(attribute.String("message", message))
	reason := reasonOpt.StringValue()
	span.SetAttributes(attribute.String("reason", reason))

	if reason != "rule 1" && reason != "rule 2" {
		span.AddEvent("Invalid reason")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Invalid reason",
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

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
	span.SetAttributes(attribute.String("muter", muter))

	// Mute user
	err = m.muteUser(ctx, uuid, message, reason, muter)
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

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("User %s muted", username),
		}})

	if err != nil {
		span.RecordError(err)
		slog.Warn("Cannot respond to interaction", "err", err)
	}

	span.AddEvent("User muted")
}

func (m *MuteCommand) muteUser(ctx context.Context, uuid, message, reason, muter string) error {
	ctx, span := m.tracer.Start(ctx, "mute-user")
	defer span.End()

	span.SetAttributes(attribute.String("uuid", uuid))
	span.SetAttributes(attribute.String("message", message))
	span.SetAttributes(attribute.String("muter", muter))

	body := chatgen.PostApiChatMuteJSONRequestBody{
		Message: strPtr(message),
		Muter:   strPtr(muter),
		Reason:  strPtr(reason),
		Uuid:    strPtr(uuid),
	}
	params := chatgen.PostApiChatMuteParams{
		Authorization: strPtr(authorization()),
	}

	response, err := m.chatClient.PostApiChatMuteWithApplicationWildcardPlusJSONBody(ctx, &params, body)
	if err != nil {
		span.RecordError(err)
		slog.Error("Cannot mute user", "err", err)
		return err
	}

	span.SetAttributes(attribute.Int("status_code", response.StatusCode))
	content, err := io.ReadAll(response.Body)
	if err != nil {
		span.RecordError(err)
		return err
	}

	span.SetAttributes(attribute.String("response_body", string(content)))
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		span.AddEvent("User muted")
		return nil
	}

	span.AddEvent("Cannot mute user")
	return fmt.Errorf("cannot mute user, status code: %d", response.StatusCode)
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
	modRole := mustEnv("MOD_ROLES")

	roles := strings.Split(modRole, ",")

	for _, r := range roles {
		if strings.TrimSpace(r) == role {
			return true
		}
	}

	return false
}
