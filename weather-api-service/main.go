package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables at startup with the directory being ./weather-api-service
	// This is optional - if .env file doesn't exist (like in Docker), environment variables
	// should already be provided by the container runtime
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: Could not load .env file (this is normal in Docker):", err)
	}

	weatherApi := APIServer{
		addr: ":8080",
	}

	err := weatherApi.Run()

	fatalError("Failed to start service", err)
}

func fatalError(message string, err error) {
	if err != nil {
		log.Fatalf(message, ":", err)
	}
}
