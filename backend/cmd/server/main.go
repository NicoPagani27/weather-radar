package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creo el router con la configuración default de Gin
	router := gin.Default()

	// Endpoint simple para probar que el servidor funciona
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	// Levanto el servidor en el puerto 8080
	router.Run(":8080")
}
