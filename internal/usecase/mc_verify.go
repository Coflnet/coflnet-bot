package usecase

import (
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/kafka"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/rs/zerolog/log"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/vmihailenco/msgpack/v5"
	"strconv"
)

// when a user verifies their account, maybe he is qualified for the flipper role
// so consume every message from the mc_verify topic and update the user's flipper role

// StartMcVerifyConsumer this function is blocking
func StartMcVerifyConsumer() {
	for {
		msg, err := kafka.ConsumeMcVerify()
		if err != nil {
			log.Error().Err(err).Msg("failed to consume mc_verify topic")
			continue
		}

		if err := processMcVerifyKafkaMessage(msg); err != nil {
			log.Error().Err(err).Msgf("error happened when checking mc verify message")
			metrics.McVerifyMessageError()
			continue
		}

		metrics.McVerifyMessageProcessed()

		err = kafka.CommitMcVerify(msg)
		if err != nil {
			log.Error().Err(err).Msg("failed to commit mc_verify topic")
			continue
		}
	}
}

func processMcVerifyKafkaMessage(msg *kafkago.Message) error {
	var verificationMessage McVerifyMessageContent
	err := msgpack.Unmarshal(msg.Value, &verificationMessage)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal mc_verify message")
		return err
	}

	id, err := strconv.Atoi(verificationMessage.UserId)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert user id to int, sky mc verify consume")
		return err
	}

  log.Info().Msgf("loading user with id %d, because he verified his mc account", id)
	user, err := coflnet.LoadUserById(id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to load user with id %d", id)
		return err
	}

	return discord.SetFlipperRoleForUser(user)
}

type McVerifyMessageContent struct {
	UserId   string `msgpack:"userId"`
	UUID     string `msgpack:"uuid"`
	Existing int    `msgpack:"existing"`
}
