package mongo

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
)

type ChatMessage struct {
	UUID       string    `bson:"uuid,omitempty" json:"uuid"`
	Name       string    `bson:"name,omitempty" json:"name"`
	Prefix     string    `bson:"prefix,omitempty" json:"prefix"`
	Message    string    `bson:"message,omitempty" json:"message"`
	ClientName string    `bson:"client_name,omitempty" json:"clientName"`
	Timestamp  time.Time `bson:"timestamp,omitempty" json:"timestamp"`
}

func InsertDiscordMessage(ctx context.Context, message *discordgo.Message) error {
    // create a span
    _, span := tracer.Start(ctx, "insert-discord-message")
    defer span.End()


	message.Timestamp = time.Now()
	_, err := messageCollection.InsertOne(ctx, message)


	if err != nil {
		return err
	}

	return err

}
