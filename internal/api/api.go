package api

import "github.com/gin-gonic/gin"

func Start() error {

	r := gin.Default()

	r.GET("/api/xp/:player", xpOfPlayer)
	r.GET("/api/player/:id", getUserById)
	r.GET("/api/player/discord/:tag", GetUserByDiscordTag)

	return r.Run()
}
