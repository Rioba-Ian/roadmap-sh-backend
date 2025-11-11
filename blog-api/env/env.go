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
	PostgresUser            EnvKey = "DB_USER"
	PostgresPassword        EnvKey = "DB_PASSWORD"
	PostgresDefaultDatabase EnvKey = "DB_NAME"
	PostgresDatabaseAddress EnvKey = "DB_HOST"
	PostgresDatabasePort    EnvKey = "DB_PORT"
)

func Load() error {
	return godotenv.Load(".env")
}
