package config

import (
	"app/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// To initialize generic configuration models for the application

func LoadEnv() {
	// Initialize the configuration
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found; using system environment variables")
	}
}

// Load the database configuration from the environment variables
func NewDatabaseConfigFromEnv() models.Config {
	return models.Config{
		Database: models.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
}
