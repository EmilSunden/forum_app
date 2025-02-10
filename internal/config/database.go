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
			PostgresDB:   os.Getenv("POSTGRES_DATABASE"),
			PostgresUser: os.Getenv("POSTGRES_USER"),
			PostgresPass: os.Getenv("POSTGRES_PASSWORD"),
			PostgresPort: os.Getenv("POSTGRES_PORT"),
		},
	}
}

func (config DatabaseConfig) GetPostgres() string {
	return config.DatabaseConfig.PostgresDB
}

func (config DatabaseConfig) GetPostgresUser() string {
	return config.DatabaseConfig.PostgresUser
}

func (config DatabaseConfig) GetPostgresPass() string {
	return config.DatabaseConfig.PostgresPass
}

func (config DatabaseConfig) GetPostgresPort() string {
	return config.DatabaseConfig.PostgresPort
}
