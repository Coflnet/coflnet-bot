package main

import (
	"github.com/Coflnet/coflnet-bot/api"
	"github.com/Coflnet/coflnet-bot/mongo"
	"os"

	"github.com/Coflnet/coflnet-bot/discord"
	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var session *discordgo.Session

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Error loading .env file")
	} else {
		log.Info().Msg("loaded .env file")
	}

	log.Info().Msg("starting metrics server")
	go metrics.Init()

	err = mongo.Init()
	if err != nil {
		log.Error().Err(err).Msg("error connecting to database")
	}
	defer mongo.CloseConnection()

	go func() {
		err := api.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("fatal error from api server")
		}
	}()

	session = getSession()
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	errorCh := make(chan error)

	go discord.ObserveMessages(session, errorCh)
	err = session.Open()
	if err != nil {
		log.Fatal().Err(err).Msgf("discord open failed")
	}

	log.Info().Msg("discord started")

	err = <-errorCh
	log.Fatal().Err(err).Msgf("quit program due to error")

}

func getSession() *discordgo.Session {
	if session != nil {
		return session
	}
	var err error
	log.Info().Msgf("login: %s", "Bot "+os.Getenv("DISCORD_BOT_TOKEN"))
	session, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Error().Err(err).Msgf("error getting discord session")
	}
	return session
}
