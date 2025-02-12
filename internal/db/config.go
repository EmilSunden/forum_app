package db

import (
	"os"
)

type DatabaseConfig struct {
	PostgresHost string
	PostgresDB   string
	PostgresUser string
	PostgresPass string
	PostgresPort string
}

// Load the database configuration from the environment variables
func NewDatabaseConfigFromEnv() DatabaseConfig {
	return DatabaseConfig{
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresDB:   os.Getenv("POSTGRES_DATABASE"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASSWORD"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}
}

func (config DatabaseConfig) GetPostgresHost() string {
	return config.PostgresHost
}

func (config DatabaseConfig) GetPostgres() string {
	return config.PostgresDB
}

func (config DatabaseConfig) GetPostgresUser() string {
	return config.PostgresUser
}

func (config DatabaseConfig) GetPostgresPass() string {
	return config.PostgresPass
}

func (config DatabaseConfig) GetPostgresPort() string {
	return config.PostgresPort
}
