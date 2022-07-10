package mongo

import (
	"context"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"time"
)

func InsertUnMute(unmute *model.Unmute) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := unMuteCollection.InsertOne(ctx, unmute)
	if err != nil {
		return err
	}

	return nil
}
