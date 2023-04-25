package db

import (
	"context"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "discord"
)

func NewDB() *DB {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.MongoHost()))
	if err != nil {
		panic(err)
	}

	return &DB{
		client: client,
	}
}

type DB struct {
	client *mongo.Client
}
