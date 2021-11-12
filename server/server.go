package server

import (
	"strconv"

	"github.com/Flou21/discord-role-updater/database"
	"github.com/Flou21/discord-role-updater/hypixel"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// @title Fiber Example API
// @version 1.0
// @description This is the documentation for the discord role service
// @BasePath /
func Start(errorCh chan<- error) {
	app := fiber.New()

	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/user/:id", getUser)

	app.Post("/message/:msg", sendMessage)

	log.Info().Msg("starting server")

	errorCh <- app.Listen(":3000")
}

// @Summary get user
// @Description get details of a user
// @ID get-user
// @Param id path int true "User ID"
// @Success 200 {object} database.User
// @Failure 400 {object} string
// @Router /user/:id [get]
func getUser(c *fiber.Ctx) error {

	id := c.Params("id")

	if id == "" {
		log.Warn().Msg("getUser no id given: " + id)
		return c.SendStatus(400)
	}
	log.Info().Msg("getUser called, id: " + id)

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		log.Err(err).Msg("got an id that is not a number " + id)
		return c.SendStatus(400)
	}

	user, err := database.UserById(parsedId)

	if err != nil {
		log.Error().Err(err).Msgf("error when loading the user %s from the database", user.Id)
	}

	if err = hypixel.User(&user); err != nil {
		log.Error().Err(err).Msgf("error when adding information to user %s from hypixel", user.Id)
	}

	return c.JSON(user)
}

func sendMessage(c *fiber.Ctx) error {
	return nil
}
