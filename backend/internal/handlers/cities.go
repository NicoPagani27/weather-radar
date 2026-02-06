package handlers

import (
	"net/http"
	"weather-radar/backend/internal/cities"
	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {
	c.JSON(http.StatusOK, cities.GetAll())
}