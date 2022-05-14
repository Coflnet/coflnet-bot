package mongo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type ChatMessage struct {
	UUID       string `bson:"uuid,omitempty" json:"uuid"`
	Name       string `bson:"name,omitempty" json:"name"`
	Prefix     string `bson:"prefix,omitempty" json:"prefix"`
	Message    string `bson:"message,omitempty" json:"message"`
	ClientName string `bson:"client_name,omitempty" json:"clientName"`
}

func InsertChatMessage(message *ChatMessage) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	_, err := coflChatCollection.InsertOne(ctx, message)

	if err != nil {
		log.Error().Err(err).Msgf("could not save message %s", message.Message)
		return err
	}

	return err

}
