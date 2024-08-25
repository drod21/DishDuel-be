package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	if err := initDB(); err != nil {
		log.Fatal(err)
	}

	// Start the server
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}
