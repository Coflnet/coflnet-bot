package chat

import (
	"fmt"
	"os"

	"github.com/Coflnet/coflnet-bot/metrics"
	"github.com/Coflnet/coflnet-bot/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var (
	session    *discordgo.Session
	coflChatId string
)

func InitDiscord() {
	session = getSession()
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	coflChatId = os.Getenv("DISCORD_COFLCHAT_ID")

	go ObserveMessages()
	err := session.Open()

	if err != nil {
		log.Error().Err(err).Msgf("error in discord session")
	}
}

func getSession() *discordgo.Session {
	if session != nil {
		return session
	}
	var err error
	log.Info().Msgf("login: %s", "Bot "+os.Getenv("DISCORD_BOT_TOKEN"))
	session, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Error().Err(err).Msgf("error getting discord session")
	}
	return session
}

func ObserveMessages() {
	log.Info().Msgf("adding message handler")
	session.AddHandler(messageCreate)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info().Msgf("received discord message: %s", m.Content)

	err := mongo.InsertMessage(m.Message)

	if err != nil {
		log.Error().Err(err).Msgf("error when inserting message")
		metrics.ErrorOccured()
	}

	err = sendMessageToChatApi(m)
	if err != nil {
		log.Error().Err(err).Msgf("error when sending message to chat api")
		metrics.ErrorOccured()
	}

	metrics.MessageProcessed()
}

func SendMessageToDiscordChat(message *mongo.ChatMessage) error {

	if message.UUID == "" {
		return fmt.Errorf("no icon url is set")
	}

	iconUrl := fmt.Sprintf("https://crafatar.com/avatars/%s", message.UUID)

	_, err := session.ChannelMessageSendComplex(coflChatId, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: message.Message,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    message.Name,
					IconURL: iconUrl,
				},
				Footer: &discordgo.MessageEmbedFooter{
					IconURL: "https://avatars.githubusercontent.com/u/42452044?s=280&v=4",
					Text:    message.ClientName,
				},
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error sending message to discord chat")
		return err
	}

	return nil
}

func sendInvalidUUIDMessageToDiscord(message *discordgo.Message) {
	session.ChannelMessageSendReply(message.ChannelID, "no uuid found for player "+message.Author.Username, &discordgo.MessageReference{
		MessageID: message.ID,
		ChannelID: message.ChannelID,
		GuildID:   message.GuildID,
	})
}
