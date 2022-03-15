package api

import (
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/xp"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func xpOfPlayer(c *fiber.Ctx) error {

	name := c.Params("player")

	if name == "" {
		log.Error().Msg("got request for xp of player, but no name is given")
		return c.Status(fiber.StatusBadRequest).SendString("no name is given")
	}

	player, err := xp.XpOfPlayer(name)
	if err != nil {
		log.Error().Err(err).Msgf("there was an error when searching xp for player: %s", name)
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("could not retrieve xp for player %s", name))
	}

	serialized, err := json.Marshal(player)
	if err != nil {
		log.Error().Err(err).Msgf("error serializing player %s", name)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).Send(serialized)
}
