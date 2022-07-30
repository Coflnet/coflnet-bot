package usecase

import (
	"encoding/json"
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/kafka"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/rs/zerolog/log"
	"time"
)

func StartFlipSummaryProcessing() {
	for {
		err := processFlipSummary()
		if err != nil {
			log.Error().Err(err).Msg("error processing flip summary")
			metrics.FlipSummaryProcessingError()
		}

		metrics.FlipSummarySend()
		time.Sleep(time.Millisecond * 1500)
	}
}

func processFlipSummary() error {
	msg, err := kafka.ConsumeFlipSummary()

	if err != nil {
		return err
	}

	var flip model.Flip
	err = json.Unmarshal(msg.Value, &flip)

	if err != nil {
		return err
	}

	err = discord.FlipTracked(&flip)
	if err != nil {
		return err
	}

	return kafka.CommitFlipSummary(msg)
}
