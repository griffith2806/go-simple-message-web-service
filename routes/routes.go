package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/messages", getMessages)
	server.GET("/messages/:id", getMessage)
	server.POST("/messages", createMessage)
}
