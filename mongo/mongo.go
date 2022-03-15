package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var (
	client            *mongo.Client
	database          *mongo.Database
	messageCollection *mongo.Collection
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

	return nil
}

func CloseConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
}
