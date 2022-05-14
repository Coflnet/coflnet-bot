package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Coflnet/coflnet-bot/internal/xp"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func xpOfPlayer(c *gin.Context) {

	name := c.Param("player")

	if name == "" {
		log.Error().Msg("got request for xp of player, but no name is given")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name not found",
		})

		return
	}

	player, err := xp.XpOfPlayer(name)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error when searching xp for player: %s", name)

		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not retrieve xp for player %s", name)})
		return
	}

	serialized, err := json.Marshal(player)
	if err != nil {
		log.Error().Err(err).Msgf("error serializing player %s", name)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not serialize player %s", name)})
		return
	}

	c.JSON(http.StatusOK, serialized)
}
