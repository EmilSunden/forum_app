package server

import (
	"app/internal/config"
	"app/internal/db"
	"app/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func Start() error {
	config.LoadEnv()
	portConf := config.LoadPortFromEnv()
	port := portConf.GetPort()
	addr := fmt.Sprintf(":%s", port)
	// initialize DB
	db, err := db.InitializeGormDB()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	mux := routes.Routes(db)

	log.Printf("Server is listening on port %s", addr)
	return http.ListenAndServe(addr, mux)
}
