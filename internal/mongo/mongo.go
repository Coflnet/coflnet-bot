package mongo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	client             *mongo.Client
	database           *mongo.Database

	messageCollection  *mongo.Collection
	coflChatCollection *mongo.Collection
	userCollection     *mongo.Collection
	muteCollection     *mongo.Collection
	unMuteCollection   *mongo.Collection
	warningCollection  *mongo.Collection

    tracer trace.Tracer
)

func Init() error {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoHost()))
	if err != nil {
		return err
	}

	database = client.Database("discord")
	messageCollection = database.Collection("messages")
	coflChatCollection = database.Collection("cofl_chat")
	userCollection = database.Collection("users")
	muteCollection = database.Collection("mutes")
	unMuteCollection = database.Collection("unmutes")
	warningCollection = database.Collection("warnings")

    tracer = otel.Tracer("mongo")

	return nil
}

func CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("error closing mongo connection")
	}
}

func mongoHost() string {
  return utils.Env("MONGO_HOST")
}
