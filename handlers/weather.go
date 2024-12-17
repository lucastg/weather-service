package handlers

import (
	"net/http"
	"weather-service/services"
	"weather-service/utils"

	"github.com/gin-gonic/gin"
)

func GetWeatherByZipcode(c *gin.Context) {
	zipcode := c.Param("zipcode")

	if len(zipcode) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	location, err := services.GetLocationByZipcode(zipcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	tempC, err := services.GetTemperature(location.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch temperature"})
		return
	}

	tempF := utils.CelsiusToFahrenheit(tempC)
	tempK := utils.CelsiusToKelvin(tempC)

	c.JSON(http.StatusOK, gin.H{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	})
}
