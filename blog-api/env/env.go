package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	PostgresUser            EnvKey = "DB_USER"
	PostgresPassword        EnvKey = "DB_PASSWORD"
	PostgresDefaultDatabase EnvKey = "DB_NAME"
	PostgresDatabaseAddress EnvKey = "DB_HOST"
	PostgresDatabasePort    EnvKey = "DB_PORT"
)

func Load() error {

	if os.Getenv("DB_USER") != "" && os.Getenv("DB_HOST") != "" {
		log.Println("Env variables already loaded from the deployment platform")
		return nil
	}

	err := godotenv.Load(".env.prod")
	if err != nil {
		log.Println("Loaded env variables from the .env.prod")
		return nil
	}

	log.Println("No .env.prod found, trying .env")
	return godotenv.Load(".env")
}
