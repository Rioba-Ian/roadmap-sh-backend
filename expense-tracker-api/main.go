package main

import (
	"log"
	"os"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/api"
	"github.com/Rioba-Ian/expense-tracker-api/cmd/database"
	"github.com/Rioba-Ian/expense-tracker-api/config"
	"github.com/Rioba-Ian/expense-tracker-api/helpers"
)

func main() {
	config.LoadEnv()
	jwtKey := config.GenerateRandomKey()
	helpers.SetJWTKey(jwtKey)

	db := database.InitDB()
	defer db.Close()
	port := os.Getenv("PORT")
	log.Printf("port provided is %s\n", port)
	if port == "" {
		port = "8080"
	}
	server := api.NewApiServer(port, db)

	log.Printf("Server listening on port %s", port)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start http server", err)
	}
}
