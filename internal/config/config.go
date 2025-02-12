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

// ServerConfig holds the server port
type ServerConfig struct {
	Port string
}

// LoadJWTConfigFromEnv loads the JWT secret key from the environment variables
func LoadJWTConfigFromEnv() JWTConfig {
	return JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}

// GetJWTSecret returns the JWT secret key
func (config JWTConfig) GetJWTSecret() string {
	return config.Secret
}

// LoadServerConfigFromEnv loads the server configuration from the environment variables
func LoadServerConfigFromEnv() ServerConfig {
	return ServerConfig{
		Port: os.Getenv("PORT"),
	}

}

// GetPort returns the server port
func (config ServerConfig) GetPort() string {
	return config.Port
}
