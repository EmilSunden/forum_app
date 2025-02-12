package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// To load the environment variables from the .env file
func LoadEnv() {
	// Initialize the configuration
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found; using system environment variables")
	}
}

// JWTConfig holds the JWT secret key
type JWTConfig struct {
	Secret string
}

type ServerConfig struct {
	Port string
}

// LoadJWTConfigFromEnv loads the JWT secret key from the environment variables
func LoadJWTConfigFromEnv() JWTConfig {
	return JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}

func (config JWTConfig) GetJWTSecret() string {
	return config.Secret
}

func LoadServerConfigFromEnv() ServerConfig {
	return ServerConfig{
		Port: os.Getenv("PORT"),
	}

}

func (config ServerConfig) GetPort() string {
	return config.Port
}
