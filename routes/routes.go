package routes

import (
	docs "example.com/gg-messages/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) //Swagger relative route is /swagger/index.html

	server.GET("/messages", getMessages)
	server.GET("/messages/:id", getMessage)
	server.POST("/messages", createMessage)
}
