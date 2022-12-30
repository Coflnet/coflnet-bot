package usecase

import (
	"encoding/json"
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/kafka"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/rs/zerolog/log"
	kafkago "github.com/segmentio/kafka-go"
	"strconv"
	"time"
)

func StartTransactionConsumer() {
	for {
		msg, err := kafka.ReadTransactionMessage()
		if err != nil {
			log.Error().Err(err).Msg("error reading kafka message")
			metrics.ErrorOccurred()
			continue
		}

		err = processTransactionMessage(msg)
		if err != nil {
			log.Error().Err(err).Msg("error processing kafka message")
			metrics.ErrorOccurred()
			continue
		}

		err = kafka.CommitTransactionMessage(msg)
		if err != nil {
			log.Error().Err(err).Msg("error committing kafka message")
			metrics.ErrorOccurred()
			continue
		}
	}
}

func processTransactionMessage(message kafkago.Message) error {
	t := TransactionMessage{}
	err := json.Unmarshal(message.Value, &t)
	if err != nil {
		log.Error().Err(err).Msgf("error deserializing kafka message: %v", string(message.Value))
		return err
	}

	// only check if amount is negative, recommended by ekwav
	if t.Amount < 0 {
		log.Info().Msgf("got a transaction message with negative amount, ignoring")
		return nil
	}

	id, err := strconv.Atoi(t.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("could not parse userId from kafka %s", t.UserId)
		return err
	}

	user, err := coflnet.LoadUserById(id)
	if err != nil {
		log.Error().Err(err).Msgf("loading user with id failed %d, cannot update role", t.UserId)
		return err
	}

	return UpdateUserFlipperRole(user)
}

func UpdateUserFlipperRole(user *model.User) error {
	return discord.SetFlipperRoleForUser(user)
}

type TransactionMessage struct {
	Id           int64     `json:"Id"`
	UserId       string    `json:"UserId"`
	ProductSlug  string    `json:"ProductSlug"`
	ProductId    int64     `json:"ProductId"`
	OwnedSeconds int64     `json:"OwnedSeconds"`
	Amount       float64   `json:"Amount"`
	Timestamp    time.Time `json:"Timestamp"`
}
