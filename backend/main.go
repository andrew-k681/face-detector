package main

import (
	"log"
	"os"

	"face-detection-app/internal/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

func main() {
	// check OpenCV
	if gocv.OpenCVVersion() == "" {
		log.Fatal("OpenCV library not found. Please install OpenCV.")
		return
	}

	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))

	// API routes
	api := router.Group("/api")
	{
		api.POST("/detect-face", handlers.DetectFaceHandler)
		api.GET("/health", handlers.HealthHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
