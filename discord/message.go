package discord

import (
	"context"
	"os"
	"time"

	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	log.Info().Msgf("storing a message")

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		metrics.ErrorOccured()
		log.Error().Err(err).Msgf("can not get channel for message %s and channel id %s", m.Content, m.ChannelID)
	}

	var attachmentUrls []string
	for _, attachemnt := range m.Attachments {
		attachmentUrls = append(attachmentUrls, attachemnt.ProxyURL)
	}
	msg := DiscordMessage{
		Content:     m.Content,
		Author:      m.Author.Username,
		Timestamp:   time.Now(),
		Channel:     channel.Name,
		Attachments: attachmentUrls,
	}

	database := client.Database("discord")
	messageCollection := database.Collection("messages")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = messageCollection.InsertOne(ctx, msg)

	if err != nil {
		log.Error().Err(err).Msgf("error when inserting message")
		metrics.ErrorOccured()
	}

	metrics.MessageProcessed()
}

type DiscordMessage struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Content     string             `bson:"content,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Channel     string             `bson:"channel"`
	Attachments []string           `bson:"attachments"`
	Timestamp   time.Time          `bson:"timestamp,omitempty"`
}
