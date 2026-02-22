package main

import (
	"example-crud/config"
	"example-crud/models"
	"example-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Item{}) // creates table automatically

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
