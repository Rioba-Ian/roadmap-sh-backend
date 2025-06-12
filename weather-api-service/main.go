package main

import (
	"log"
)

func main() {

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
