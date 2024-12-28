// Package main is the entry point for the Gemini API application
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gemini-question/internal/config"
	"gemini-question/internal/handler"
	"gemini-question/internal/service"
)

// main initializes and starts the HTTP server
func main() {
	// Load application configuration from environment variables
	cfg := config.LoadConfig()

	// Initialize the Gemini service with API key
	geminiService := service.NewGeminiService(cfg.GeminiAPIKey)

	// Initialize the HTTP handler with the service
	geminiHandler := handler.NewGeminiHandler(geminiService)

	// Setup Gin router with default middleware
	r := gin.Default()

	// Register routes
	r.POST("/generate", geminiHandler.GenerateContent)

	// Start the HTTP server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
