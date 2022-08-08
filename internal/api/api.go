package api

import (
	_ "github.com/Coflnet/coflnet-bot/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api
func Start() error {

	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/api/xp/:player", xpOfPlayer)
	r.GET("/api/player/:id", getUserById)
	r.GET("/api/player/discord/:tag", GetUserByDiscordTag)

	r.GET("/api/warned", warnedUsers)
	r.GET("/api/expiredWarnings", expiredWarnings)

	return r.Run()
}
