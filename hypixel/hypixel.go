package hypixel

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Coflnet/coflnet-bot/database"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

func User(user *database.User) error {

	cachedUser, err := checkCache(user)

	if err != nil {
		log.Error().Err(err).Msgf("error when checking cache for user %s", user.MinecraftUuid)
	}

	if cachedUser != "" {
		log.Info().Msgf("found cached user %v for minecraft uuid %s", cachedUser, user.MinecraftUuid)
		user.DiscordName = formatDiscordUserName(cachedUser)
		return nil
	}

	key := os.Getenv("API_KEY")

	uuid, err := user.MinecraftUuid.Value()
	if err != nil || uuid == "" {
		l := log.Warn()
		if err != nil {
			l = l.Err(err)
		}
		l.Msgf("user %s has no minecraft uuid", user.Id)
		return nil
	}
	url := fmt.Sprintf("https://api.hypixel.net/player?key=%s&uuid=%s", key, uuid)

	log.Info().Msgf("calling hypixel api, url: %s", url)

	response, err := http.Get(url)

	if err != nil {

		log.Error().Err(err).Msgf("error when fetching url %s", url)
		return err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error when parsing url %s", url)
		return err
	}

	var data HypixelApiResponse

	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Error().Err(err).Msgf("error when parsing hypixel response json, %s", url)
		return err
	}

	discord := data.Player.SocialMedia.Links.Discord

	if discord == "" {
		log.Warn().Msgf("discord name for user %s is not set", user.Id)
		return nil
	} else {
		log.Info().Msgf("got an discord user with the name %s", discord)
	}

	discord = formatDiscordUserName(discord)

	user.DiscordName = discord

	err = saveUserToCache(user)
	if err != nil {
		log.Error().Err(err).Msgf("error when saving user %s to cache", user.DiscordName)
	}

	return nil
}

// checks if the chache already contains the discord name for a given user
func checkCache(user *database.User) (string, error) {
	ctx := context.Background()
	rdb := database.RedisClient()

	result, err := rdb.Get(ctx, user.Id).Result()

	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		log.Error().Err(err).Msgf("error when getting key %s from redis", user.MinecraftUuid)
		return "", err
	}

	return result, nil
}

func saveUserToCache(user *database.User) error {
	ctx := context.Background()
	rdb := database.RedisClient()
	err := rdb.Set(ctx, user.Id, user.DiscordName, time.Minute*5).Err()
	if err != nil {
		log.Error().Err(err).Msgf("error when saving user %s to redis", user.Id)
	} else {
		log.Info().Err(err).Msgf("saved user %s to redis", user.Id)

	}
	return err
}

// remove the #xxxx from the discord name
func formatDiscordUserName(original string) string {
	if !strings.Contains(original, "#") {
		return original
	}

	arr := strings.Split(original, "#")
	return arr[0]
}
