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

type JWTSecretKey struct {
	JWTSecretKey string
}

type Port struct {
	Port string
}

func LoadJWTSecretFromEnv() JWTSecretKey {
	return JWTSecretKey{
		JWTSecretKey: os.Getenv("JWT_SECRET"),
	}
}

func (jwt JWTSecretKey) GetJWTSecret() string {
	return jwt.JWTSecretKey
}

func LoadPortFromEnv() Port {
	return Port{
		Port: os.Getenv("PORT"),
	}

}

func (port Port) GetPort() string {
	return port.Port
}
