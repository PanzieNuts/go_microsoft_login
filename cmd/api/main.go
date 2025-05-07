package main

import (
	"bondly/config"
	"bondly/services"
	"bondly/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables and config
	config.LoadConfig()
	utils.InitLogger()

	// ✅ Initialize OAuth2 config
	services.InitializeOAuth()

	// Initialize DB
	_, err := utils.InitDatabase()
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	r := gin.Default()

	// Microsoft OAuth routes
	r.GET("/auth/microsoft", services.HandleMicrosoftLogin)
	r.GET("/auth/microsoft/callback", services.HandleMicrosoftCallback)

	// Start server on port 8000
	r.Run(":8000")
}
