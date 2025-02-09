package config

import (
	"app/internal/models"
	"os"
)

type DatabaseConfig struct {
	DatabaseConfig models.DatabaseConfig
}

// Load the database configuration from the environment variables
func NewDatabaseConfigFromEnv() DatabaseConfig {
	return DatabaseConfig{
		DatabaseConfig: models.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
}

func (config DatabaseConfig) GetHost() string {
	return config.DatabaseConfig.Host
}

func (config DatabaseConfig) GetPort() string {
	return config.DatabaseConfig.Port
}

func (config DatabaseConfig) GetUsername() string {
	return config.DatabaseConfig.Username
}

func (config DatabaseConfig) GetPassword() string {
	return config.DatabaseConfig.Password
}

func (config DatabaseConfig) GetDatabase() string {
	return config.DatabaseConfig.Database
}
