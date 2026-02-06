package handlers

import (
	"net/http"

	"weather-radar/backend/internal/cities"
	"weather-radar/backend/internal/weather"

	"github.com/gin-gonic/gin"
)


// CompareWeather recibe una lista de IDs de ciudades
// y devuelve un resumen comparativo usando concurrencia.
func CompareWeather(c *gin.Context) {

	var request struct {
		CityIDs []string `json:"city_ids"`
	}


	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "body inválido",
		})
		return
	}

	// Canal donde todas las goroutines enviarán resultados
	resultsChan := make(chan weather.CityWeather)

	// Contador de ciudades válidas puede no haber validas
	validCities := 0

	// Fan-out: una goroutine por ciudad
	for _, id := range request.CityIDs {
		city, found := cities.GetByID(id)
		if !found {
			continue
		}

		validCities++
		go weather.FetchCityWeather(city, resultsChan)
	}

	// Slice donde se juntan los resultados (fan-in)
	successful := make([]weather.CityWeather, 0)
	failed := make([]weather.CityWeather, 0)

	for i := 0; i < validCities; i++ {
		result := <-resultsChan

	// Si la ciudad falló, la separamos
		if result.Error != "" {
			failed = append(failed, result)
		} else {
			// Si salió bien, va al grupo de exitosas
			successful = append(successful, result)
		}	
}
	// Usamos nuestra función Aggregate para calcular resúmenes
	summary := weather.Aggregate(successful)

	// Devolvemos JSON
	c.JSON(http.StatusOK, gin.H{
		"cities":  successful,
		"errors":  failed,
		"summary": summary,
	})
}


//Lo que hace esto es:
//leer las request (IDs) de las ciudades.
//crear el channel (canal) para recibir resultados.
//lanzar una goroutine por cada ciudad válida (fan-out).
//esperar a que todas las goroutines terminen y envíen resultados (fan-in), separanto exitosos de fallidos.
//se calcula agregados usando aggregate.
//respondemos JSON con ciudades exitosas, errores y resumen.
