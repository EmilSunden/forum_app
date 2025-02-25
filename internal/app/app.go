package app

import (
	"app/internal/config"
	"app/internal/psql"
	"app/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func App() error {
	config.LoadEnv()
	portConf := config.LoadServerConfigFromEnv()
	port := portConf.GetPort()
	addr := fmt.Sprintf(":%s", port)

	// initialize DB
	db, err := psql.InitializeGormDB()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	ginEngine := routes.Routes(db)

	log.Printf("Server is listening on port %s", addr)
	return http.ListenAndServe(addr, ginEngine)
}
