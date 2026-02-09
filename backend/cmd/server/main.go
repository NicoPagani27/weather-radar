package main

import (
	"net/http"
	// Handlers HTTP de la aplicación
	"weather-radar/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creo el router con la configuración default de Gin
	router := gin.Default()
	// Middleware para CORS (Cross-Origin Resource Sharing) 
	// Tuve un error de failed to fech y agrego este CORS para permitir 
	// que el frontend (que corre en otro puerto) pueda hacer peticiones al backend.
	router.Use(func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
})
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









