package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func RedisClient() *redis.Client {

	if client != nil {
		return client
	}

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func RedisClose() {
	if client != nil {
		client.Close()
		client = nil
	}
}
