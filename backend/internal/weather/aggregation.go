package weather

// AggregationResult guarda los datos que calculamos
// a partir de varias ciudades.
// Por ejemplo promedios, extremos y agrupación por clima.
type AggregationResult struct {
	Averages    map[string]float64   `json:"averages"`    // Promedios: temp, humedad, viento
	Extremes    map[string]string    `json:"extremes"`    // Ciudades más caliente, fría y ventosa
	ByCondition map[string][]string  `json:"by_condition"`// Agrupación por condición climática
}

// Aggregate recibe un slice de CityWeather y devuelve
// los valores agregados que necesitamos para el resumen
func Aggregate(results []CityWeather) AggregationResult {
	// Validación: si no hay resultados, devolver estructura vacía
	if len(results) == 0 {
		return AggregationResult{
			Averages:    make(map[string]float64),
			Extremes:    make(map[string]string),
			ByCondition: make(map[string][]string),
		}
	}

	var totalTemp float64
	var totalHum int
	var totalWind float64
	var hottest CityWeather
	var coldest CityWeather
	var windiest CityWeather
	byCondition := make(map[string][]string)

	// Recorremos cada ciudad
	for i, r := range results {
		totalTemp += r.Temperature
		totalHum += r.Humidity
		totalWind += r.WindSpeed

		// Extras
		if i == 0 || r.Temperature > hottest.Temperature {
			hottest = r
		}
		if i == 0 || r.Temperature < coldest.Temperature {
			coldest = r
		}
		if i == 0 || r.WindSpeed > windiest.WindSpeed {
			windiest = r
		}

		// Agrupación por condición
		byCondition[r.Condition] = append(byCondition[r.Condition], r.CityName)
	}

	// Se calcula los promedios
	count := float64(len(results))
	averages := map[string]float64{
		"temperature": totalTemp / count,
		"humidity":    float64(totalHum) / count,
		"wind_speed":  totalWind / count,
	}

	// Se arma el resultado final
	return AggregationResult{
		Averages: averages,
		Extremes: map[string]string{
			"hottest":  hottest.CityName,
			"coldest":  coldest.CityName,
			"windiest": windiest.CityName,
		},
		ByCondition: byCondition,
	}
}
