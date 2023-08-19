package api

import (
	"github.com/Coflnet/coflnet-bot/internal/discord"
	pubdiscord "github.com/Coflnet/coflnet-bot/pkg/discord"
	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	DiscordHandler *discord.DiscordHandler
}

func NewWebhookController(handler *discord.DiscordHandler) *WebhookController {
	return &WebhookController{
		DiscordHandler: handler,
	}
}

func (con *WebhookController) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/message", con.messageWebhook)
}

func (con *WebhookController) messageWebhook(c *gin.Context) {
	// read the message body
	var message pubdiscord.Message

	err := c.BindJSON(&message)
	if err != nil {
		c.String(400, "request body is not a valid discord message")
		return
	}

	err = con.DiscordHandler.SendMessage(c.Request.Context(), &message)
	if err != nil {
		c.String(500, "error sending message to discord")
		return
	}

	c.String(200, "message sent")
}
