package config

import (
	"log"

	"github.com/joho/godotenv"
)

// To load the environment variables from the .env file
func LoadEnv() {
	// Initialize the configuration
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found; using system environment variables")
	}
}

func GetPort() string {
	return ":8080"
}
