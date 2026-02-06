package handlers

import (
	"net/http"

	"weather-radar/backend/internal/cities"
	"weather-radar/backend/internal/weather"

	"github.com/gin-gonic/gin"
)


// GetWeatherByCity devuelve el clima actual de una ciudad.
// Recibe el cityId por URL, busca la ciudad y consulta en la api.

func GetWeatherByCity(c *gin.Context) {
	cityID := c.Param("cityId")

	city, found := cities.GetByID(cityID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ciudad no encontrada",
		})
		return
	}

	data, err := weather.GetCurrentWeather(city.Latitude, city.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error consultando clima",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city": city,
		"weather": gin.H{
			"temperature": data.Current.Temperature,
			"humidity":    data.Current.Humidity,
			"wind_speed":  data.Current.WindSpeed,
			"weatherCode": data.Current.WeatherCode,
		},
	})
}
