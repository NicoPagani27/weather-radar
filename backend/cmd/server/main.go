package main

import (
	"net/http"
	// Handlers HTTP de la aplicación
	"weather-radar/backend/internal/handlers"
	"weather-radar/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creo el router con la configuración default de Gin
	router := gin.Default()
	// Middleware para CORS (Cross-Origin Resource Sharing)
	router.Use(middleware.CORSMiddleware())
	// Endpoint simple para probar que el servidor funciona
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})
	// Grupo de rutas bajo /api para mantener ordenada la API
	api := router.Group("/api")
	{
		// Devuelve la lista de ciudades disponibles
		api.GET("/cities", handlers.GetCities)

		// Devuelve el clima actual de una ciudad específica
		api.GET("/weather/:cityId", handlers.GetWeatherByCity)

		// Compara el clima de lista de ciudades
		api.POST("/weather/compare", handlers.CompareWeather)
	}
	// Levanto el servidor en el puerto 8080
	router.Run(":8080")
}









