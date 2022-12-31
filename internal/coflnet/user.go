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
func LoadUserById(id int) (*model.User, error) {
	user := &model.User{
		UserId: id,
	}

	mcUser, err := UserMcConnect(user.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from mcConnect", user.UserId)
		return nil, err
	}

	return LoadUser(mcUser)
}

func LoadUserByUUID(uuid string) (*model.User, error) {
  user, err := UserMcConnectByUUID(uuid)
  if err != nil {
    return nil, err
  }

  return LoadUserById(user.UserId)
}

func LoadUser(mcUser *model.User) (*model.User, error) {
	premiumTime, err := PaymentUserById(mcUser.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from payment", mcUser.UserId)
		return nil, err
	}

	var discordNames []string
	for _, uuid := range mcUser.MinecraftUuids {
		player, err := hypixel.PlayerData(uuid)
		if err != nil {
			return nil, err
		}
    discordName := player.Player.SocialMedia.Links.Discord

    if discordName == "" {
      continue
    }

    log.Info().Msgf("found the discord name %s", discordName)
		discordNames = append(discordNames, discordName)
    time.Sleep(1 * time.Second)
	}

  user := &model.User{
		UserId:         mcUser.UserId,
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
