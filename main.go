package main

import (
	"log"
	"os"
	"weather-service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	weatherAPIKey := os.Getenv("WEATHERAPI_KEY")
	if weatherAPIKey == "" {
		log.Fatal("WeatherAPI key not set")
	}

	r := gin.Default()

	r.GET("/weather/:zipcode", handlers.GetWeatherByZipcode)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}
}