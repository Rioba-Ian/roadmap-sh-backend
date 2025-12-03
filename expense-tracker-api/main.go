package main

import (
	"log"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/api"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/database"
	"github.com/Rioba-Ian/expense-tracker-api/config"
	"github.com/Rioba-Ian/expense-tracker-api/helpers"
)

func main() {
	config.LoadEnv()
	jwtKey := config.GenerateRandomKey()
	helpers.SetJWTKey(jwtKey)

	database.InitDB()
	server := api.NewApiServer(":8080")

	log.Printf("Server listening on localhost:8080")
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start http server", err)
	}
}
