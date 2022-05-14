package mongo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UserId         int       `json:"userId" bson:"user_id"`
	PremiumUntil   time.Time `json:"premiumUntil" bson:"premium_until"`
	DiscordNames   []string  `json:"discordNames" bson:"discord_names"`
	MinecraftUuids []string  `json:"minecraftUuids" bson:"minecraft_uuids"`
	LastRefresh    time.Time `json:"lastRefresh" bson:"last_refresh"`
}

func SearchByUserId(userId int) (*User, error) {

	filter := bson.D{{"user_id", userId}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	res := userCollection.FindOne(ctx, filter)

	var user User
	err := res.Decode(&user)
	if err != nil {
		if err == m.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func SearchByDiscordTag(discordTag string) (*User, error) {
	filter := bson.D{{"discord_names",
		bson.D{{"$elemMatch", bson.D{{"$eq", discordTag}}}},
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	res := userCollection.FindOne(ctx, filter)

	var user User
	err := res.Decode(&user)
	if err != nil {
		if err == m.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func InsertUser(user *User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	log.Info().Msgf("inserted user %v", result.InsertedID)

	return nil
}

func UpdateUser(user *User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	result, err := userCollection.ReplaceOne(ctx, bson.D{{"user_id", user.UserId}}, user)
	if err != nil {
		return err
	}

	log.Info().Msgf("modified %v user (%d)", result.ModifiedCount, user.UserId)

	return nil
}

func SaveUser(user *User) error {
	oldUser, err := SearchByUserId(user.UserId)
	if err != nil {
		return err
	}

	if oldUser == nil {
		return InsertUser(user)
	}

	return UpdateUser(user)
}

func DeleteUser(user *User) error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_, err := userCollection.DeleteOne(ctx, bson.D{{"user_id", user.UserId}})
	if err != nil {
		return err
	}

	return nil
}

func (u *User) IsPremium() bool {
	return u.PremiumUntil.After(time.Now())
}
