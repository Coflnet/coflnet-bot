package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/rs/zerolog/log"
)

type ChatMessage struct {
	UUID       string    `bson:"uuid,omitempty" json:"uuid"`
	Name       string    `bson:"name,omitempty" json:"name"`
	Prefix     string    `bson:"prefix,omitempty" json:"prefix"`
	Message    string    `bson:"message,omitempty" json:"message"`
	ClientName string    `bson:"client_name,omitempty" json:"clientName"`
	Timestamp  time.Time `bson:"timestamp,omitempty" json:"timestamp"`
}

func InsertChatMessage(message *ChatMessage) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	message.Timestamp = time.Now()
	_, err := coflChatCollection.InsertOne(ctx, message)

	if err != nil {
		log.Error().Err(err).Msgf("could not save message %s", message.Message)
		return err
	}

	return err

}

func FirstChatMessageOfUsername(username string) (*ChatMessage, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var msg ChatMessage
	res := coflChatCollection.FindOne(ctx, bson.D{{"name", username}})

	if res.Err() != nil {
		log.Error().Err(res.Err()).Msgf("could not find message for user %s", username)
		return nil, res.Err()
	}

	err := res.Decode(&msg)
	if err != nil {
		log.Error().Err(err).Msgf("could not decode message for user %s", username)
		return nil, err
	}

	if err != nil {
		log.Error().Err(err).Msgf("could not find user %s", username)
		return nil, err
	}

	return &msg, nil
}
