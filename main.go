package main

import (
	"api_go/config"
	"api_go/models"
	"api_go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Conectar y migrar DB
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run(":8080")
}
