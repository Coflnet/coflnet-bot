package xp

import (
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/rs/zerolog/log"
)

type Player struct {
	Name string `json:"name"`
	Xp   int32  `json:"xp"`
}

//goland:noinspection GoNameStartsWithPackageName
func XpOfPlayer(playerName string) (*Player, error) {

	messagesOfPlayer, err := mongo.CountMessagesOfPlayer(playerName)

	if err != nil {
		log.Error().Err(err).Msgf("error getting messages of player %s", playerName)
		return nil, err
	}

	xp := messagesOfPlayer * 1

	p := &Player{
		Name: playerName,
		Xp:   xp,
	}

	return p, nil
}
