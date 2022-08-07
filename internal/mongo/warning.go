package mongo

import (
	"context"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

func InsertWarning(warning *model.Warning) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	warning.Id = primitive.NewObjectID()

	_, err := warningCollection.InsertOne(ctx, warning)
	return err
}

func WarningsByUser(user *discordgo.Member) ([]*model.Warning, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"$and": []bson.M{
		{"user.user.id": user.User.ID},
		{"until": bson.M{"$gt": time.Now()}},
	}}

	var warnings []*model.Warning
	res, err := warningCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	err = res.All(ctx, &warnings)
	if err != nil {
		return nil, err
	}

	return warnings, nil
}

func ExpiredWarnings() (<-chan model.Warning, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	filter := bson.M{"until": bson.M{"$lt": time.Now()}}

	warnings := make(chan model.Warning, 0)
	res, err := warningCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer func(c *mongo.Cursor, ct context.Context) {
		_ = c.Close(ct)
	}(res, ctx)

	go func() {
		wg := sync.WaitGroup{}
		var batch []model.Warning
		for res.Next(ctx) {
			if err := res.Decode(&batch); err != nil {
				log.Error().Err(err).Msg("failed to decode warnings")
				continue
			}

			for _, el := range batch {
				wg.Add(1)
				go func(w model.Warning) {
					defer wg.Done()
					userWarnings, err := WarningsByUser(w.User)
					if err != nil {
						log.Error().Err(err).Msg("failed to get user warnings")
						return
					}

					if len(userWarnings) == 0 {
						warnings <- w
					}
				}(el)
			}
		}

		wg.Wait()
		close(warnings)
	}()

	return warnings, nil
}
