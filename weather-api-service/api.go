package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type APIServer struct {
	addr string
}

type ApiResponse struct {
	data WeatherRes
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/", handleHomeHandler)
	router.HandleFunc("/weather", handleWeatherResp)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started on %s", s.addr)

	return server.ListenAndServe()
}

func handleWeatherResp(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "City parameter is missing. Usage: /weather?city=yourcity", http.StatusBadRequest)
		return
	}

	fmt.Printf("The city chosen= %s", city)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	resp, err := http.Get(fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&include=days%%2Ccurrent%%2Chours&key=%s&contentType=json", city, os.Getenv("WEATHER_API_KEY")))
	if err != nil {
		log.Fatalf("Failed to get response %v", err)
	}
	r.Header.Set("Content-Type", "application/json")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body %v", err)
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse.data); err != nil {
		log.Fatalf("Error unmarshalling Json %v", err)
		return
	}

	indentedJson, err := json.MarshalIndent(apiResponse.data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON with indent:", err)
		return
	}

	fmt.Println(string(indentedJson))

	fmt.Printf("The city is: %s\n", apiResponse.data.Address)

	w.WriteHeader(http.StatusOK)
	w.Write(indentedJson)
}

func handleHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Welcome to weather api, use /weather?city='Your city' to interact with service."))
}
