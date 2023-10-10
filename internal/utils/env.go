package utils

import (
	"fmt"
	"os"
	"strings"
)

func Env(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		err := fmt.Errorf("missing %s environment variable", key)
		panic(err)
	}
	return value
}

func KafkaHost() (string, error) {
	return optionalEnvVar("KAFKA_BROKERS")
}

func KafkaCA() string {
	return Env("KAFKA_TLS_CA_LOCATION")
}

func KafkaCert() string {
	return Env("KAFKA_TLS_CERTIFICATE_LOCATION")
}

func KafkaKey() string {
	return Env("KAFKA_TLS_KEY_LOCATION")
}

func KafkaUser() string {
	return Env("KAFKA_USERNAME")
}

func KafkaPassword() string {
	return Env("KAFKA_PASSWORD")
}

func McVerifyKafkaTopic() string {
	return Env("TOPIC_MC_VERIFY")
}

func McVerifyKafkaConsumerGroup() string {
	return Env("TOPIC_MC_VERIFY_CONSUMER_GROUP")
}

func DevChatKafkaTopic() string {
	return Env("TOPIC_DEV_CHAT")
}

func TransactionKafkaTopic() string {
	return Env("TOPIC_TRANSACTION")
}

func FlipSummaryTopic() string {
	return Env("TOPIC_FLIPSUMMARY")
}

func FlipSummaryConsumerGrup() string {
	return Env("TOPIC_FLIPSUMMARY_CONSUMER_GROUP")
}

func DebugEnabled() bool {
	return Env("DEBUG") == "true"
}

func RedisChatPubSubChannel() string {
	return Env("REDIS_CHAT_PUBSUB_CHANNEL")
}

func RedisHost() (string, error) {
	return optionalEnvVar("REDIS_HOST")
}

func coflChatId() string {
	return getEnv("DISCORD_COFLCHAT_ID")
}

func muteRoles() []string {
	return strings.Split(getEnv("DISCORD_MUTE_ROLE"), ",")
}

func HelperRole() string {
	return getEnv("HELPER_ROLE")
}

func devRole() string {
	return getEnv("DEV_ROLE")
}

func ModRole() string {
	return getEnv("MOD_ROLE")
}

func MutesWebhook() string {
	return getEnv("DISCORD_MUTES_WEBHOOK")
}

func WarningsWebhook() string {
	return getEnv("DISCORD_WARNINGS_WEBHOOK")
}

func CiSuccessWebhook() string {
	return getEnv("DISCORD_CI_SUCCESS_WEBHOOK")
}

func CiFailureWebhook() string {
	return getEnv("DISCORD_CI_FAILURE_WEBHOOK")
}

func FlipperRoleWebhook() string {
	return getEnv("DISCORD_FLIPPER_ROLE_WEBHOOK")
}

func FeedbackWebhook() string {
	return getEnv("DISCORD_FEEDBACK_WEBHOOK")
}

func FlipWebhook() string {
	return getEnv("FLIP_WEBHOOK")
}

func SongvoterFeedbackWebhook() string {
	return getEnv("SONGVOTER_FEEDBACK_WEBHOOK")
}

func optionalEnvVar(e string) (string, error) {
	v := os.Getenv(e)
	if v == "" {
		err := fmt.Errorf("missing %s environment variable", e)
		return "", err
	}
	return v, nil
}

func getEnv(e string) string {
	v, err := optionalEnvVar(e)
	if err != nil {
		panic(err)
	}
	return v
}

func DiscordBotToken() string {
	return getEnv("DISCORD_BOT_TOKEN")
}

func DiscordGuildId() string {
	return getEnv("DISCORD_GUILD_ID")
}

func ChatBaseUrl() (string, error) {
	return optionalEnvVar("CHAT_BASE_URL")
}

func ChatChannelId() string {
	return getEnv("CHAT_CHANNEL_ID")
}

func DiscordInGameChannelId() string {
	return getEnv("DISCORD_IN_GAME_CHANNEL_ID")
}

func McConnectBaseUrl() (string, error) {
	return optionalEnvVar("MC_CONNECT_BASE_URL")
}

func PaymentBaseUrl() (string, error) {
	return optionalEnvVar("PAYMENT_BASE_URL")
}

func ApiBaseUrl() (string, error) {
	return optionalEnvVar("API_BASE_URL")
}

func DiscordMessageConsumerGroup() string {
	return getEnv("DISCORD_MESSAGE_CONSUMER_GROUP")
}

func DiscordMessageTopic() string {
	return getEnv("DISCORD_MESSAGE_TOPIC")
}

func MongoHost() string {
	return getEnv("MONGO_HOST")
}

func CoflnetBotBaseUrl() (string, error) {
	return optionalEnvVar("COFLNET_BOT_BASE_URL")
}

func ChatApiKey() string {
	return getEnv("CHAT_API_KEY")
}

func StrPtr(s string) *string {
	return &s
}
