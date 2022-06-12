package mongo

import (
	"context"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/rs/zerolog/log"
	"time"
)

func InsertMute(mute *model.Mute) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err := muteCollection.InsertOne(ctx, mute)

	if err != nil {
		log.Error().Err(err).Msgf("error inserting mute for user %s from %s", mute.Username, mute.Muter)
		return err
	}

	log.Info().Msgf("inserted mute %s for user %s from %s", res.InsertedID, mute.Username, mute.Muter)
	return nil
}
