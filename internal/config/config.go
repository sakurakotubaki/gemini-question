// Package config handles the application configuration
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration values for the application
type Config struct {
	GeminiAPIKey string // API key for accessing Gemini API
}

// LoadConfig loads the configuration from environment variables
// It automatically loads variables from .env file if present
// Returns a pointer to Config struct containing all configuration values
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
	}
}
