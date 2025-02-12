package main

import (
	"app/internal/server"
	"log"
)

func main() {
	// Initialize the server
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
