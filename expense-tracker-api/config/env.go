package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err != nil {

		err := godotenv.Load()

		if err != nil {
			log.Println("Failed to load .env")
		} else {
			log.Println("Successfully loaded .env file")
		}
	} else {
		log.Println(".env file not found, using env variables from container")
	}
}
