package main

import (
	"github.com/Coflnet/coflnet-bot/api"
	"github.com/Coflnet/coflnet-bot/chat"
	"github.com/Coflnet/coflnet-bot/mongo"

	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var session *discordgo.Session

func main() {

	// env vars
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Error loading .env file")
	} else {
		log.Info().Msg("loaded .env file")
	}

	// metrics
	log.Info().Msg("starting metrics server")
	go metrics.Init()

	// mongo
	err = mongo.Init()
	if err != nil {
		log.Error().Err(err).Msg("error connecting to database")
	}
	defer mongo.CloseConnection()

	// redis
	startRedisChatConsume()

	// open discord session and wait for messages
	chat.InitDiscord()

	// start api
	err = api.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("fatal error from api server")
	}
}

func startRedisChatConsume() {
	go func() {
		err := chat.StartConsume()
		if err != nil {
			log.Fatal().Err(err).Msgf("error consuming messages from chat")
		}
	}()
}
