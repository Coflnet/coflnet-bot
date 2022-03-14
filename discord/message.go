package discord

import (
	"context"
	"os"
	"time"

	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ObserveMessages(session *discordgo.Session, errorCh chan<- error) {
	log.Info().Msgf("adding message handler")
	session.AddHandler(messageCreate)

	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		errorCh <- err
	}
}

func CloseConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info().Msgf("received message: %s", m.Content)

	database := client.Database("discord")
	messageCollection := database.Collection("messages")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := messageCollection.InsertOne(ctx, m.Message)

	if err != nil {
		log.Error().Err(err).Msgf("error when inserting message")
		metrics.ErrorOccured()
	}

	metrics.MessageProcessed()
}
