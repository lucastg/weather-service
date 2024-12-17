package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInvalidZipcode(t *testing.T) {
	r := gin.Default()
	r.GET("/weather/:zipcode", handlers.GetWeatherByZipcode)

	req, _ := http.NewRequest("GET", "/weather/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.JSONEq(t, `{"message": "invalid zipcode"}`, w.Body.String())
}
