package db

import (
	"context"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		collection: db.client.Database(dbName).Collection("users"),
	}
}

type UserRepo struct {
	collection *mongo.Collection
}

func (u *UserRepo) SearchByDiscordTag(ctx context.Context, searchTerm string) (*model.User, error) {
	filter := bson.D{{Key: "discord_names", Value: searchTerm}}

	res := u.collection.FindOne(ctx, filter)

	var user model.User
	err := res.Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &model.UserNotFoundError{}
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) SearchUserByDiscordId(ctx context.Context, searchId string) (*model.User, error) {
	filter := bson.D{{Key: "discord_ids", Value: searchId}}

	res := u.collection.FindOne(ctx, filter)

	var user model.User
	err := res.Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &model.UserNotFoundError{
				DiscordId: searchId,
			}
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) InsertEmptyModelUser(ctx context.Context, user *model.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u *UserRepo) UnsetDiscordTag(ctx context.Context, userId int) error {
	var filter = bson.D{{Key: "user_id", Value: userId}}
	var update = bson.D{{Key: "$unset", Value: bson.D{{Key: "discord_names", Value: ""}}}}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	return err
}

func (u *UserRepo) SetDiscordIdForUser(ctx context.Context, userId int, discordId []string) error {
	var filter = bson.D{{Key: "user_id", Value: userId}}
	var update = bson.D{{Key: "$set", Value: bson.D{{Key: "discord_id", Value: discordId}}}}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	return err
}
