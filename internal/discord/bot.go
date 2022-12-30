package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/pkg/discord"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var (
	session *discordgo.Session
)

func InitDiscord() {
	session = getSession()
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	go ObserveMessages()
	err := session.Open()

	if err != nil {
		log.Error().Err(err).Msgf("error in discord session")
	}

	err = registerCommands()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to register commands")
	}
}

func StopDiscord() {
	err := unregisterCommands()
	if err != nil {
		return
	}

	err = session.Close()
	if err != nil {
		log.Error().Err(err).Msgf("error closing discord session")
		return
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

func messageCreate(_ *discordgo.Session, m *discordgo.MessageCreate) {
	err := mongo.InsertMessage(m.Message)

	if err != nil {
		log.Error().Err(err).Msgf("error when inserting message")
		metrics.ErrorOccurred()
	}

	err = SendMessageToChatApi(m)
	if err != nil {
		log.Error().Err(err).Msgf("error when sending message to chat api")
		metrics.ErrorOccurred()
	}

	metrics.MessageProcessed()
}

type AllowedMentions struct {
	Parse []string `json:"parse"`
}

type WebhookRequest struct {
	Content             string          `json:"content"`
	Username            string          `json:"username"`
	AvatarUrl           string          `json:"avatar_url"`
	AllowedMentionsData AllowedMentions `json:"allowed_mentions"`
}

func GiveUserWarnedRole(user *discordgo.Member, w *model.Warning) error {
	// give role
	err := session.GuildMemberRoleAdd(guildId(), user.User.ID, warnedRole())
	if err != nil {
		return err
	}

	// send message to player
	return SendMessageToUser(fmt.Sprintf("⚠️ you have been warned by %s until %s", username(user), DiscordFormattedTime(w.Until)), user)
}

func RemoveUserWarnedRole(user *discordgo.Member) error {

	user, err := session.GuildMember(guildId(), user.User.ID)
	if err != nil {
		log.Error().Err(err).Msgf("error refreshing user")
		return err
	}

	roleFound := false
	for _, role := range user.Roles {
		if role == warnedRole() {
			roleFound = true
		}
	}

	if !roleFound {
		// user does not have the warned role skip it
		return nil
	}

	err = session.GuildMemberRoleRemove(guildId(), user.User.ID, warnedRole())
	if err != nil {
		return err
	}

	return SendMessageToWarningsChannel(fmt.Sprintf("⚠️ warned role has been removed from user %s", username(user)))
}

func SendMessageToUser(message string, member *discordgo.Member) error {
	channel, err := session.UserChannelCreate(member.User.ID)
	if err != nil {
		return err
	}

	_, err = session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}

	return nil
}

func SendMessageToDiscordChat(message *mongo.ChatMessage) error {

	if message.UUID == "" {
		return fmt.Errorf("no icon url is set")
	}

	iconUrl := fmt.Sprintf("https://crafatar.com/avatars/%s", message.UUID)
	url := os.Getenv("CHAT_WEBHOOK")

	msg := message.Message
	data := &WebhookRequest{
		Content:             sanitizeMessage(msg),
		Username:            message.Name,
		AvatarUrl:           iconUrl,
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msgf("can not marshal webhook request")
	}

	_, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Error().Err(err).Msgf("error sending message to discord chat")
		return err
	}

	return nil
}

func sendInvalidUUIDMessageToDiscord(message *discordgo.Message, msg string) {
	_, err := session.ChannelMessageSendReply(message.ChannelID, msg, &discordgo.MessageReference{
		MessageID: message.ID,
		ChannelID: message.ChannelID,
		GuildID:   message.GuildID,
	})

	if err != nil {
		log.Error().Err(err).Msgf("there was an error when sending the message to discord")
	}
}

func sanitizeMessage(message string) string {
	// if the strings has a § in it, remove all § characters and the following character

	log.Info().Msgf("sanitizing message: %s", message)
	reg := regexp.MustCompile("§.")
	message = reg.ReplaceAllString(message, "")
	log.Info().Msgf("sanitized message: %s", message)

	return message
}

func SendMessageToCiSuccess(msg string) error {
	return SendMessageToNotificationServer(discord.DiscordMessageToSend{
		Message: msg,
		Webhook: ciSuccessWebhook(),
	})
}

func SendMessageToCiFailureChannel(msg string) error {
	return SendMessageToNotificationServer(discord.DiscordMessageToSend{
		Message: msg,
		Webhook: ciFailureWebhook(),
	})
}

func SendMessageToWarningsChannel(msg string) error {
	return SendMessageToNotificationServer(discord.DiscordMessageToSend{
		Message: msg,
		Webhook: warningsWebhook(),
	})
}

func SendMessageToMutesChannel(msg string) error {
	return SendMessageToNotificationServer(discord.DiscordMessageToSend{
		Message: msg,
		Webhook: mutesWebhook(),
	})
}

func SendMessageToFlipperRoleChannel(msg string) error {
	return SendMessageToNotificationServer(discord.DiscordMessageToSend{
		Message: msg,
		Webhook: flipperRoleWebhook(),
	})
}

func SendMessageToNotificationServer(msg discord.DiscordMessageToSend) error {

	if msg.Message == "" {
		return fmt.Errorf("no message is set")
	}

	if msg.Webhook == "" {
		if msg.Channel != "" {
			msg.Webhook = webhookForChannel(msg.Channel)
		}
	}

	if msg.Webhook == "" {
		return fmt.Errorf("no webhook is set, webhook: %s, channel: %s", msg.Webhook, msg.Channel)
	}

	data := &WebhookRequest{
		Content:             msg.Message,
		Username:            "cofl-bot",
		AvatarUrl:           "https://cdn.discordapp.com/app-icons/888725077191974913/0c0e3b97e6865091ef14162083a54a42.png?size=256",
		AllowedMentionsData: AllowedMentions{Parse: make([]string, 0)},
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msgf("can not marshal webhook request")
		return err
	}

	url := msg.Webhook
	_, err = http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Error().Err(err).Msgf("error when sending webhook request")
		return err
	}

	return nil
}

func webhookForChannel(channel string) string {
	switch channel {
	case discord.WarningsChannel:
		return warningsWebhook()
	case discord.MutesChannel:
		return mutesWebhook()
	case discord.CiSuccessChannel:
		return ciSuccessWebhook()
	case discord.CiFailureChannel:
		return ciFailureWebhook()
	case discord.FlipperRoleChannel:
		return flipperRoleWebhook()
	case discord.FeedbackChannel:
		return feedbackWebhook()
	}

	log.Error().Msg("no webhook found for channel")
	return ""
}

func DiscordFormattedTime(t time.Time) string {
	return fmt.Sprintf("<t:%d:f>", t.Unix())
}
