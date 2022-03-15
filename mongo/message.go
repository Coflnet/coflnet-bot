package mongo

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func InsertMessage(message *discordgo.Message) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := messageCollection.InsertOne(ctx, message)

	return err
}
func CountMessagesOfPlayer(name string) (int32, error) {
	matchStage := bson.D{{"$match", bson.D{{"author", name}}}}
	countStage := bson.D{{"$count", "author"}}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cursor, err := messageCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		countStage,
	})

	if err != nil {
		log.Error().Err(err).Msgf("error when counting messages for player")
		return 0, err
	}

	defer cursor.Close(ctx)

	var res []bson.M

	err = cursor.All(ctx, &res)
	if err != nil {
		log.Error().Err(err).Msgf("error retrieving result from cursor")
		return 0, err
	}

	if len(res) == 0 {
		return 0, fmt.Errorf("no player with name %s found", name)
	}

	r := res[0]["author"]

	return r.(int32), nil
}
