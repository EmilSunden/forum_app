package server

import (
	"app/internal/routes"
	"log"
	"net/http"
)

func Start() error {
	mux := routes.SetupRoutes()
	addr := ":8080"
	log.Printf("Server is listening on port %s", addr)
	return http.ListenAndServe(addr, mux)
}
