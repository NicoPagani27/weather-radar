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
	results := make([]weather.CityWeather, 0)

	for i := 0; i < validCities; i++ {
		result := <-resultsChan
		results = append(results, result)
	}

	// Condiciones extras

	var (
		totalTemp  float64
		totalHum   int
		totalWind  float64
		hottest    weather.CityWeather
		coldest    weather.CityWeather
		windiest   weather.CityWeather
		byCondition = make(map[string][]string)
	)

	for i, r := range results {
		totalTemp += r.Temperature
		totalHum += r.Humidity
		totalWind += r.WindSpeed

		if i == 0 || r.Temperature > hottest.Temperature {
			hottest = r
		}
		if i == 0 || r.Temperature < coldest.Temperature {
			coldest = r
		}
		if i == 0 || r.WindSpeed > windiest.WindSpeed {
			windiest = r
		}

		byCondition[r.Condition] = append(byCondition[r.Condition], r.CityName)
	}

	// Promedios
	count := float64(len(results))
	averages := gin.H{
		"temperature": totalTemp / count,
		"humidity":    float64(totalHum) / count,
		"wind_speed":  totalWind / count,
	}

	// Respuesta final
	c.JSON(http.StatusOK, gin.H{
		"cities": results,
		"summary": gin.H{
			"averages": averages,
			"extremes": gin.H{
				"hottest":  hottest.CityName,
				"coldest":  coldest.CityName,
				"windiest": windiest.CityName,
			},
			"by_condition": byCondition,
		},
	})
}


//Lo que hace esto es:
//leer las request (IDs) de las ciudades.
//crear el channel (canal) para recibir resultados.
//lanzar una goroutine por cada ciudad válida (fan-out).
//esperar a que todas las goroutines terminen y envíen resultados (fan-in).
//procesar los resultados
//responden JSON
