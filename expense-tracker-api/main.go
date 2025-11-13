package main

import (
	"log"

	"github.com/Rioba-Ian/expense-tracker-api/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080")

	log.Printf("Server listening on localhost:8080")
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start http server", err)
	}
}
