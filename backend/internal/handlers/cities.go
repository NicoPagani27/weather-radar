package handlers

import (
	"net/http"
	"weather-radar/backend/internal/cities"
	"github.com/gin-gonic/gin"
)

// GetCities devuelve todas las ciudades disponibles.
// Este endpoint es para el frontend
// para mostrar el selector de ciudades.

func GetCities(c *gin.Context) {
	c.JSON(http.StatusOK, cities.GetAll())
}