package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func warnedUsers(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}

func expiredWarnings(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}
