package router

import (
	"github.com/gin-gonic/gin"
	"vediomeeting/internal/server/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// meeting
	r.POST("/meeting/create", service.MeetingCreate)
	return r
}
