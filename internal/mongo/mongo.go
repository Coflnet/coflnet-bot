package mongo

import (
	"context"
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

	return nil
}

func CloseConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
}
