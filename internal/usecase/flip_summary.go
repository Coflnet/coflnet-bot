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
		somethingSent, err := processFlipSummary()
		if err != nil {
			log.Error().Err(err).Msg("error processing flip summary")
			metrics.FlipSummaryProcessingError()
		}

		if !somethingSent {
			continue
		}

		metrics.FlipSummarySend()
		time.Sleep(time.Millisecond * 1500)
	}
}

// returns true if somthing was sent
func processFlipSummary() (bool, error) {
	msg, err := kafka.ConsumeFlipSummary()

	if err != nil {
		return false, err
	}

	var flip model.Flip
	err = json.Unmarshal(msg.Value, &flip)

	if err != nil {
		return false, err
	}

	if flip.Profit < 500_000 {
		return false, nil
	}

	err = discord.FlipTracked(&flip)
	if err != nil {
		return false, err
	}

	return true, kafka.CommitFlipSummary(msg)
}
