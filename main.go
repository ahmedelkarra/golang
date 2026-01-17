package main

import (
	"log"
	"ping-api/database"
	"ping-api/middleware"
	"ping-api/models"
	"ping-api/user"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/api/register", user.Register)
	r.POST("/api/login", user.Login)
	r.GET("/api/me", middleware.IsUser, user.Me)

	r.Run()
}
