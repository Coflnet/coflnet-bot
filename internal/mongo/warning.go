package mongo

import (
	"context"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func InsertWarning(warning *model.Warning) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := warningCollection.InsertOne(ctx, warning)
	return err
}

func WarningsByUser(user *discordgo.User) ([]*model.Warning, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"$and": []bson.M{
		{"user.id": user.ID},
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
