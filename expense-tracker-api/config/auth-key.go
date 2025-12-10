package config

import (
	"os"
)

func GenerateRandomKey() string {
	// bytes := make([]byte, 32)
	// _, err := rand.Read(bytes)
	// if err != nil {
	// 	log.Fatal("Failed to generate key", err)
	// }

	// if helpers.IsProd() {
	// 	return base64.URLEncoding.EncodeToString(bytes)
	// }

	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtSecret
}
