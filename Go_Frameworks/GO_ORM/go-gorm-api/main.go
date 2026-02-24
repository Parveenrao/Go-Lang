package main

import (
	"go-gorm-api/config"
	"go-gorm-api/models"
	"go-gorm-api/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	config.Connect()

	
	// Auto migrate
	config.DB.AutoMigrate(&models.User{})

	routes.SetupRoute(r)

	r.Run(":8080")

} 
