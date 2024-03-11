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

type UnmuteCommand struct {
	tracer      trace.Tracer
	session     *discordgo.Session
	userService *UserService
	chatClient  *chatgen.Client
}

func NewUnmuteCommand(session *discordgo.Session, userService *UserService, chatClient *chatgen.Client) *UnmuteCommand {
	return &UnmuteCommand{
		session:     session,
		userService: userService,
		chatClient:  chatClient,
		tracer:      otel.Tracer("unmute-command"),
	}
}

func (u *UnmuteCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ctx, span := u.tracer.Start(ctx, "unmute-command")
	defer span.End()

	slog.Info(fmt.Sprintf("unmute user, trace id: %s", span.SpanContext().TraceID()))

	options := i.ApplicationCommandData().Options

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

	username := opt.Value.(string)
	span.SetAttributes(attribute.String("username", username))
	reason := reasonOpt.Value.(string)
	span.SetAttributes(attribute.String("reason", reason))

	if !userAllowedToMute(i.Member) {
		span.AddEvent("User not allowed to unmute")
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are not allowed to unmute",
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

	uuid, err := u.userService.LoadUserUUIDByMinecraftName(ctx, username)
	if err != nil {
		span.RecordError(err)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Cannot find user with name %s", username),
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

	unmuter := i.Member.User.ID
	span.SetAttributes(attribute.String("unmuter", unmuter))

	err = u.unmuteUser(ctx, uuid, unmuter, reason)
	if err != nil {
		span.RecordError(err)
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Cannot unmute user %s", username),
			},
		})
		if err != nil {
			span.RecordError(err)
			slog.Warn("Cannot respond to interaction", "err", err)
		}
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("User %s has been unmuted", username),
		},
	})

	if err != nil {
		span.RecordError(err)
		slog.Warn("Cannot respond to interaction", "err", err)
	}

	span.AddEvent("User has been unmuted")
}

func (u *UnmuteCommand) unmuteUser(ctx context.Context, uuid, unmuter, reason string) error {
	ctx, span := u.tracer.Start(ctx, "unmute-user")
	defer span.End()

	span.SetAttributes(attribute.String("uuid", uuid))
	span.SetAttributes(attribute.String("unmuter", unmuter))
	span.SetAttributes(attribute.String("reason", reason))

	body := chatgen.DeleteApiChatMuteJSONRequestBody{
		Reason:  strPtr(reason),
		UnMuter: strPtr(unmuter),
		Uuid:    strPtr(uuid),
	}

	params := chatgen.DeleteApiChatMuteParams{
		Authorization: strPtr(authorization()),
	}

	response, err := u.chatClient.DeleteApiChatMuteWithApplicationWildcardPlusJSONBody(ctx, &params, body)
	if err != nil {
		span.RecordError(err)
		return err
	}

	span.SetAttributes(attribute.Int("response", response.StatusCode))

	if response.StatusCode > 299 {
		err := fmt.Errorf("cannot unmute user, status code: %d", response.StatusCode)
		span.RecordError(err)
		return err
	}

	span.AddEvent("User has been unmuted")
	return nil
}
