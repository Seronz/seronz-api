package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seronz/seronz-api/internal/email"
	"github.com/seronz/seronz-api/internal/health"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin, Content-Type"},
	}))

	r.GET("/health", health.Check)
	r.POST("/api/send", email.Send)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("seronz-api jalan di :" + port)
	r.Run(":" + port)
}
