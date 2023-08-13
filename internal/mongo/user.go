package mongo

import (
	"context"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
)

type UserNotFoundError struct {
	UserId int
}

func SearchByDiscordTag(discordTag string) (*model.User, error) {
	filter := bson.D{{"discord_names", discordTag}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := userCollection.FindOne(ctx, filter)

	var user model.User
	err := res.Decode(&user)
	if err != nil {
		if err == m.ErrNoDocuments {
			return nil, &model.UserNotFoundError{}
		}

		return nil, err
	}

	return &user, nil
}
