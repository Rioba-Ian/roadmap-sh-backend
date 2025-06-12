package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

type ApiResponse struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/weather", func(http.ResponseWriter, *http.Request) {
		resp, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/nairobi?unitGroup=us&include=days&key=5M9ZW8K38956LGNUVN695WQTR&contentType=json")
		if err != nil {
			log.Fatalf("Failed to get response %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error reading response body %v", err)
		}

		var apiResponse ApiResponse
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Fatalf("Error unmarshalling Json %v", err)
		}
		prettyJson, err := json.Indent(apiResponse, "", "\t")

		fmt.Println(prettyJson)

		fmt.Printf("The city is: %s\n", apiResponse.Address)

	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started on %s", s.addr)

	return server.ListenAndServe()
}
