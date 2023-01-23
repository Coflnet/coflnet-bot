package mongo

import (
	"context"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserNotFoundError struct {
	UserId int
}

func SearchByUserId(userId int) (*model.User, error) {

	filter := bson.D{{"user_id", userId}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	res := userCollection.FindOne(ctx, filter)

	var user model.User
	err := res.Decode(&user)
	if err != nil {
		if err == m.ErrNoDocuments {
			return nil, &model.UserNotFoundError{UserId: userId}
		}

		return nil, err
	}

	return &user, nil
}

func SearchByDiscordTag(discordTag string) (*model.User, error) {
	filter := bson.D{{"discord_names",
		bson.D{{"$elemMatch", bson.D{{"$eq", discordTag}}}},
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

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

func InsertUser(user *model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Error().Err(err).Msgf("error when inserting user %s", user.UserId)
		return err
	}

	return nil
}

func UpdateUser(user *model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := userCollection.ReplaceOne(ctx, bson.D{{"user_id", user.UserId}}, user)
	if err != nil {
		return err
	}

	log.Info().Msgf("modified %v user (%d)", result.ModifiedCount, user.UserId)

	return nil
}

func SaveUser(user *model.User) error {
	_, err := SearchByUserId(user.UserId)
	if err != nil {

		if _, ok := err.(*model.UserNotFoundError); ok {

			return InsertUser(user)
		}

		return err
	}

	return UpdateUser(user)
}

func SetFlipperRoleForUser(user *model.User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.D{{"user_id", user.UserId}}
	update := bson.D{{"$set", bson.D{{"has_flipper_role", user.HasFlipperRole}}}}

	r, err := userCollection.UpdateOne(ctx, filter, update, &options.UpdateOptions{})

	if err != nil {
		log.Error().Err(err).Msgf("there was an error when setting has_flipper_role for user %d", user.UserId)
		return err
	}

	log.Info().Msgf("set has_flipper_role %t for user %d, %d documents affected", user.HasFlipperRole, user.UserId, r.ModifiedCount)
	return nil
}

func GetUsersWithFlipperRole() (<-chan *model.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	userChan := make(chan *model.User)

	filter := bson.D{{"has_flipper_role", true}}

	cursor, err := userCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	go func() {

		for cursor.Next(ctx) {
			var users []*model.User
			err = cursor.Decode(&users)
			if err != nil {
				continue
			}

			for _, u := range users {
				userChan <- u
			}
		}

		close(userChan)
	}()

	return userChan, nil
}

func GetUserByDiscordName(name string) (*model.User, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  filter := bson.D{{"discord_names", name}}
  var user model.User
  res := userCollection.FindOne(ctx, filter)

  if err := res.Err(); err != nil {
    return nil, res.Err()
  }

  err := res.Decode(&user)
  if err != nil {
    return nil, err
  }

  return &user, nil
}

func GetUserByMinecraftUuid(uuid string) (*model.User, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  filter := bson.D{{"minecraft_uuids", uuid}}
  var user model.User
  res := userCollection.FindOne(ctx, filter)

  if err := res.Err(); err != nil {
    return nil, res.Err()
  }

  err := res.Decode(&user)
  if err != nil {
    return nil, err
  }

  return &user, nil
}
