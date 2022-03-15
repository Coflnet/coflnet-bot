package discord

import (
	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/Coflnet/coflnet-bot/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func ObserveMessages(session *discordgo.Session, errorCh chan<- error) {
	log.Info().Msgf("adding message handler")
	session.AddHandler(messageCreate)

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info().Msgf("received message: %s", m.Content)

	err := mongo.InsertMessage(m.Message)

	if err != nil {
		log.Error().Err(err).Msgf("error when inserting message")
		metrics.ErrorOccured()
	}

	metrics.MessageProcessed()
}
