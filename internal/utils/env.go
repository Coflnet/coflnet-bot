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

func KafkaHost() string {
  return Env("KAFKA_HOST")
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

func RedisHost() string {
  return Env("REDIS_HOST")
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

func getEnv(e string) string {
	v := os.Getenv(e)
	if v == "" {
        err := fmt.Errorf("missing %s environment variable", e)
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

func ChatBaseUrl() string {
    return getEnv("CHAT_BASE_URL")
}


func ChatChannelId() string {
    return getEnv("CHAT_CHANNEL_ID")
}

func DiscordInGameChannelId()  string {
    return getEnv("DISCORD_IN_GAME_CHANNEL_ID")
}

func McConnectBaseUrl() string {
    return getEnv("MC_CONNECT_BASE_URL")
}

func PaymentBaseUrl() string {
    return getEnv("PAYMENT_BASE_URL")
}

func ApiBaseUrl() string {
    return getEnv("API_BASE_URL")
}
