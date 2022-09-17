package discord

import (
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

func coflChatId() string {
	return getEnv("DISCORD_COFLCHAT_ID")
}

func muteRoles() []string {
	return strings.Split(getEnv("DISCORD_MUTE_ROLE"), ",")
}

func helperRole() string {
	return getEnv("HELPER_ROLE")
}

func devRole() string {
	return getEnv("DEV_ROLE")
}

func modRole() string {
	return getEnv("MOD_ROLE")
}

func mutesWebhook() string {
	return getEnv("DISCORD_MUTES_WEBHOOK")
}

func warningsWebhook() string {
	return getEnv("DISCORD_WARNINGS_WEBHOOK")
}

func ciSuccessWebhook() string {
	return getEnv("DISCORD_CI_SUCCESS_WEBHOOK")
}

func ciFailureWebhook() string {
	return getEnv("DISCORD_CI_FAILURE_WEBHOOK")
}

func flipperRoleWebhook() string {
	return getEnv("DISCORD_FLIPPER_ROLE_WEBHOOK")
}

func feedbackWebhook() string {
	return getEnv("DISCORD_FEEDBACK_WEBHOOK")
}

func getEnv(e string) string {
	v := os.Getenv(e)
	if v == "" {
		log.Panic().Msgf("%s not set", e)
	}
	return v
}
