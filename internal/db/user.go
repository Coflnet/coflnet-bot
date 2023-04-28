package db

import (
	"context"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		collection: db.client.Database(dbName).Collection("users"),
		tracer:     otel.Tracer("db-user-repo"),
	}
}

type UserRepo struct {
	collection *mongo.Collection
	tracer     trace.Tracer
}

func (u *UserRepo) SearchByDiscordTag(ctx context.Context, searchTerm string) (*model.User, error) {
	ctx, span := u.tracer.Start(ctx, "db-search-user-by-discord-tag")
	defer span.End()
	span.SetAttributes(attribute.String("discord_tag", searchTerm))

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
	ctx, span := u.tracer.Start(ctx, "db-search-user-by-discord-id")
	defer span.End()
	span.SetAttributes(attribute.String("discord_id", searchId))

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

func (u *UserRepo) SearchUserByUUID(ctx context.Context, uuid string) ([]*model.User, error) {
	ctx, span := u.tracer.Start(ctx, "db-search-user-by-uuid")
	defer span.End()
	span.SetAttributes(attribute.String("uuid", uuid))

	filter := bson.D{{Key: "minecraft_uuids", Value: uuid}}
	cur, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var user []*model.User
	err = cur.All(ctx, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) InsertEmptyModelUser(ctx context.Context, user *model.User) error {
	ctx, span := u.tracer.Start(ctx, "db-insert-empty-user")
	defer span.End()
	span.SetAttributes(attribute.Int("user_id", user.UserId))

	_, err := u.collection.InsertOne(ctx, user)
	return err
}

func (u *UserRepo) UnsetDiscordTag(ctx context.Context, userId int) error {
	ctx, span := u.tracer.Start(ctx, "db-unset-discord-tag")
	defer span.End()
	span.SetAttributes(attribute.Int("user_id", userId))

	var filter = bson.D{{Key: "user_id", Value: userId}}
	var update = bson.D{{Key: "$unset", Value: bson.D{{Key: "discord_names", Value: ""}}}}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	return err
}

func (u *UserRepo) SetDiscordIdForUser(ctx context.Context, userId int, discordId []string) error {
	ctx, span := u.tracer.Start(ctx, "db-set-discord-id-for-user")
	defer span.End()
	span.SetAttributes(attribute.Int("user_id", userId))
	if len(discordId) > 0 {
		span.SetAttributes(attribute.String("first_discord_id", discordId[0]))
	}

	var filter = bson.D{{Key: "user_id", Value: userId}}
	var update = bson.D{{Key: "$set", Value: bson.D{{Key: "discord_ids", Value: discordId}}}}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	return err
}
