package psql

import (
	"fmt"
	"log"
	"strconv"

	"app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeGormDB() (*gorm.DB, error) {
	dbConfig := NewDatabaseConfigFromEnv()
	host := dbConfig.GetPostgresHost()
	portStr := dbConfig.GetPostgresPort()
	user := dbConfig.GetPostgresUser()
	password := dbConfig.GetPostgresPass()
	dbName := dbConfig.GetPostgres()

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port: %s", portStr)
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = migrate(
		db,
		&models.User{},
		&models.FriendRequest{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return db, nil
}
