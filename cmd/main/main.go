package main

import (
	"app/internal/app"
	"log"
)

func main() {
	// Initialize the server
	if err := app.App(); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
