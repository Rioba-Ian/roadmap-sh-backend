package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"github.com/Rioba-Ian/expense-tracker-api/helpers"
)

func GenerateRandomKey() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Failed to generate key", err)
	}

	if helpers.IsProd() {
		return base64.URLEncoding.EncodeToString(bytes)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtSecret
}
