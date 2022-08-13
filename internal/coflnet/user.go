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

	user, err := mongo.SearchByUserId(id)
	if err != nil {

		// if the error is not an usernotfound return it
		if _, ok := err.(*model.UserNotFoundError); !ok {
			return nil, err
		}
	}

	if user != nil {
		if user.LastRefresh.After(time.Now().Add(-time.Hour * 24)) {
			log.Info().Msgf("found user %d in db, last refresh: %v", id, user.LastRefresh)
			return user, nil
		}
	}

	return LoadUserById(id)
}

// LoadUserById loads the user from different api's
func LoadUserById(id int) (*model.User, error) {

	user := &model.User{
		UserId: id,
	}

	mcUser, err := UserMcConnect(user.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from mcConnect", user.UserId)
		return nil, err
	}

	premiumTime, err := PaymentUserById(user.UserId)
	if err != nil {
		log.Error().Err(err).Msgf("can not load with id %d from payment", user.UserId)
		return nil, err
	}

	var discordNames []string
	for _, uuid := range mcUser.MinecraftUuids {
		player, err := hypixel.PlayerData(uuid)
		if err != nil {
			return nil, err
		}
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
