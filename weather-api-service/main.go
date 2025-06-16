package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables at startup with the directory being ./weather-api-service
	if err := godotenv.Load("./weather-api-service/.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
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
