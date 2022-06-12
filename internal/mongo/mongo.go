package mongo

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client             *mongo.Client
	database           *mongo.Database
	messageCollection  *mongo.Collection
	coflChatCollection *mongo.Collection
	userCollection     *mongo.Collection
	muteCollection     *mongo.Collection
)

func Init() error {

	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		return err
	}

	database = client.Database("discord")
	messageCollection = database.Collection("messages")
	coflChatCollection = database.Collection("cofl_chat")
	userCollection = database.Collection("users")
	muteCollection = database.Collection("mutes")

	return nil
}

func CloseConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Disconnect(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("error closing mongo connection")
	}
}
