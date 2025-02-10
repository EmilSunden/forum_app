package server

import (
	"app/internal/config"
	"app/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func Start() error {
	config.LoadEnv()
	mux := routes.SetupRoutes()
	port := config.LoadPortFromEnv().GetPort()
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server is listening on port %s", addr)
	return http.ListenAndServe(addr, mux)
}
