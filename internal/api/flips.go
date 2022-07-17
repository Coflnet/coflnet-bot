package api

import (
	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary flipTracked
// @Description webhook for tracked flips
// @Accept json
// @Success 200
// @Param email body model.FlipTrack true "message/rfc822" SchemaExample(Subject: Testmail\r\n\r\nBody Message\r\n)
// @Router /flipTracked [post]
func flipTracked(c *gin.Context) {
	var trackedFlip model.FlipTrack
	err := c.ShouldBindJSON(&trackedFlip)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = discord.FlipTracked(&trackedFlip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
