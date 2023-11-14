package main

import (
	"log"
	"net/http"

	"github.com/Vasanth-Korada/weather-tracker/handlers"
	"github.com/Vasanth-Korada/weather-tracker/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	loggingMiddleware := middleware.LoggingMiddleware

	http.Handle("/weather/", loggingMiddleware(http.HandlerFunc(handlers.WeatherHandler)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
