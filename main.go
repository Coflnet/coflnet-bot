package main

import (
	"os"

	"github.com/Coflnet/coflnet-bot/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var session *discordgo.Session

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// defer database.RedisClose()

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	session = getSession()
	session.Identify.Intents = discordgo.IntentsGuildMessages

	errorCh := make(chan error)

	// go kafka.StartConsumer(errorCh)
	// go kafka.StartProducer(errorCh)
	// go server.Start(errorCh)

	// checkAllUsers := os.Getenv("CHECK_ALL_USERS") == "true"
	// if checkAllUsers {
	// 	go kafka.CheckAllUsers()
	// }

	go discord.ObserveMessages(session, errorCh)
	err = session.Open()
	if err != nil {
		log.Fatal().Err(err).Msgf("discord open failed")
	}

	err = <-errorCh
	discord.CloseConnection()
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
