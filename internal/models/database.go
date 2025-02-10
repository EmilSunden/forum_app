package models

// DatabaseConfig holds the configuration for the database
type DatabaseConfig struct {
	PostgresDB   string
	PostgresUser string
	PostgresPass string
	PostgresPort string
}
