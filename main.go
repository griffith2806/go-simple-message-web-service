package main

import (
	"example.com/gg-messages/db"
	"example.com/gg-messages/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
