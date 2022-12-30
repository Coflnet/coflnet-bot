package coflnet

import (
	"github.com/Coflnet/coflnet-bot/internal/model"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/hypixel"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

// UserById loads a user, potentially cached
func UserById(id int) (*model.User, error) {
	return LoadUserById(id)
}

// LoadUserById loads the user from different api's
func LoadUserById(id int) (*model.User, error) {

	user := &model.User{
		UserId: id,
	}

  log.Info().Msgf("loading user with id %d", id)
	mcUser, err := UserMcConnect(user.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from mcConnect", user.UserId)
		return nil, err
	}
  log.Info().Msgf("he has %d minecraft uuids", len(mcUser.MinecraftUuids))

	premiumTime, err := PaymentUserById(user.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from payment", user.UserId)
		return nil, err
	}
  log.Info().Msgf("he has %v premium days", premiumTime)

	var discordNames []string
	for _, uuid := range mcUser.MinecraftUuids {
		player, err := hypixel.PlayerData(uuid)
		if err != nil {
			return nil, err
		}

    log.Info().Msgf("found the discord name %s", player.Player.SocialMedia.Links.Discord)
		discordNames = append(discordNames, player.Player.SocialMedia.Links.Discord)
	}

	user = &model.User{
		UserId:         id,
		MinecraftUuids: mcUser.MinecraftUuids,
		DiscordNames:   discordNames,
		PremiumUntil:   premiumTime,
	}
	user.LastRefresh = time.Now()

	err = mongo.SaveUser(user)
	if err != nil {
		log.Error().Err(err).Msgf("error when saving user with id %d", user.UserId)
		return nil, err
	}

	metrics.UserLoaded()

	return user, nil
}
