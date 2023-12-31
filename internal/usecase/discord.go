package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	client        *discordgo.Session
	discordTracer = otel.Tracer("discord")
)

func DiscordSession(ctx context.Context) (*discordgo.Session, error) {
	if client != nil {
		return client, nil
	}
	return StartSession(ctx)
}

func StartSession(ctx context.Context) (*discordgo.Session, error) {
	ctx, span := discordTracer.Start(ctx, "start-session")
	defer span.End()
	slog.Info("Starting a new discord session")

	token, err := BotToken()
	if err != nil {
		return nil, err
	}

	client, err = discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}

	// In this example, we only care about receiving message events.
	client.Identify.Intents = discordgo.IntentsGuildMessages

	err = client.Open()
	if err != nil {
		return nil, err
	}
	go startStopListener()

	return client, nil
}

func startStopListener() {

	slog.Info("Bot is now running, until the application stops")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	slog.Info("Shutting down discord session")
}

func BotToken() (string, error) {
	token, found := os.LookupEnv("BOT_TOKEN")
	if !found {
		slog.Warn("BOT_TOKEN not found")
		return "", errors.New("BOT_TOKEN not found")
	}

	slog.Debug("BOT_TOKEN found, returning")
	return token, nil
}

func guildId() string {
	id, found := os.LookupEnv("GUILD_ID")
	if !found {
		panic("GUILD_ID not found")
	}

	slog.Debug("GUILD_ID found, returning")
	return id
}

func SearchDiscordUser(ctx context.Context, name string) (*discordgo.Member, error) {
	ctx, span := discordTracer.Start(ctx, "search-discord-user")
	defer span.End()
	span.SetAttributes(attribute.String("search-term", name))

	members, err := client.GuildMembersSearch(guildId(), name, 1)
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return nil, &DiscordMemberNotFoundError{SearchTerm: name}
	}
	if len(members) > 1 {
		slog.Warn("found more than one discord member, choosing first one")
	}

	return members[0], nil
}

type DiscordMemberNotFoundError struct {
	SearchTerm string
}

func (e *DiscordMemberNotFoundError) Error() string {
	return fmt.Sprintf("discord member with search term %s not found", e.SearchTerm)
}
