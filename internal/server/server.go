package server

import (
	"app/internal/config"
	"app/internal/routes"
	"log"
	"net/http"
)

func Start() error {
	config.LoadEnv()
	mux := routes.SetupRoutes()
	addr := config.GetPort()
	log.Printf("Server is listening on port %s", addr)
	return http.ListenAndServe(addr, mux)
}
