package api

import (
	"github.com/Coflnet/coflnet-bot/internal/model"
	"net/http"
	"strconv"

	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/mongo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (con *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/:id", con.getUserById)
	router.GET("/discord/:tag", con.getUserByDiscordTag)
}

func (con *UserController) getUserById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		log.Warn().Msgf("get player id was called with empty id")
		c.String(http.StatusBadRequest, "no id given")
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Warn().Err(err).Msgf("get player id was called, but the id is not a int")
		c.String(http.StatusBadRequest, "user id has to be an int")
		return
	}

	log.Info().Msgf("search player with id %d", userId)

	user, err := coflnet.LoadUserById(userId)
	if err != nil {
		if nf, ok := err.(*model.UserNotFoundError); ok {
			log.Warn().Msgf("user with id %d not found", userId)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": nf.Error(),
			})
			return
		}
		log.Error().Err(err).Msgf("internal error occurred when searching user with id %d", userId)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (con *UserController) getUserByDiscordTag(c *gin.Context) {
	tag := c.Param("tag")

	if tag == "" {
		log.Warn().Msgf("get player tag was called with empty tag")
		c.String(http.StatusBadRequest, "no tag given")
		return
	}

	log.Info().Msgf("search player with tag %s", tag)

	user, err := mongo.SearchByDiscordTag(tag)
	if err != nil {

		if nf, ok := err.(*model.UserNotFoundError); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": nf.Error(),
			})
			return
		}

		log.Error().Err(err).Msgf("internal error occurred when searching user with tag %s", tag)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}
