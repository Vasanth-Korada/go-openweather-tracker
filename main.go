package main

import (
	"log"
	"net/http"

	"github.com/Vasanth-Korada/weather-tracker/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/weather/", api.WeatherHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
