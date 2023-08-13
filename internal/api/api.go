package api

import (
	"log/slog"

	_ "github.com/Coflnet/coflnet-bot/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ApiController struct {
	WebhookController *WebhookController
	UserController    *UserController
}

func NewApiController(userController *UserController, webhookController *WebhookController) *ApiController {
	return &ApiController{
		WebhookController: webhookController,
		UserController:    userController,
	}
}

func (con *ApiController) Start() error {

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/api/swagger/doc.json") // The url pointing to API definition
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	con.UserController.RegisterRoutes(r.Group("/api/player"))
	con.WebhookController.RegisterRoutes(r.Group("/api/webhook"))

	slog.Info("starting api server on port 8080")

	return r.Run(":8080")
}
