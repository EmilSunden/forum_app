package main

import (
	"app/internal/server"
	"log"
)

func main() {
	// Main function
	// config.LoadEnv()

	// Initialize the server
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
