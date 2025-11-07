package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	PostgresUser            EnvKey = "POSTGRES_USER"
	PostgresPassword        EnvKey = "DB_PASSWORD"
	PostgresDefaultDatabase EnvKey = "DB_NAME"
	PostgresDatabaseAddress EnvKey = "POSTGRES_DB_ADDRESS"
	PostgresDatabasePort    EnvKey = "DB_PORT"
)

func Load() error {
	return godotenv.Load(".env")
}
