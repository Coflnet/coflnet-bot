package main

import (
	"github.com/Coflnet/coflnet-bot/internal/api"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/kafka"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

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
	go startRedisChatConsume()

	// start kafka transaction consume
	go func() {
		err := kafka.StartTransactionConsume()
		if err != nil {
			log.Panic().Err(err).Msgf("error consuming messages from kafka")
		}
	}()

	// start dev chat consumer
	go kafka.StartDiscordMessagesConsumer()
	go kafka.StartDiscordSpamMessagesConsumer()

	// start kafka verification consume
	go func() {
		err := kafka.StartVerificationConsume()
		if err != nil {
			log.Panic().Err(err).Msgf("error consuming messages from kafka")
		}
	}()

	// open discord session and wait for messages
	discord.InitDiscord()
	defer discord.StopDiscord()

	// start the refresh of the user table
	// go usecase.StartRefresh()

	// start api
	err = api.Start()
	if err != nil {
		log.Panic().Err(err).Msg("fatal error from api server")
	}
}

func startRedisChatConsume() {
	err := discord.StartConsume()
	if err != nil {
		log.Panic().Err(err).Msgf("error consuming messages from chat")
	}
}
