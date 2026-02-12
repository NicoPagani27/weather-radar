package weather

// AggregationResult guarda los datos que calculamos
// a partir de varias ciudades.
// Por ejemplo promedios, extremos y agrupación por clima.
type AggregationResult struct {
	Averages    map[string]float64  `json:"averages"`     // Promedios: temp, humedad, viento
	Extremes    map[string]string   `json:"extremes"`     // Ciudades más caliente, fría y ventosa
	ByCondition map[string][]string `json:"by_condition"` // Agrupación por condición climática
}

// Aggregate recibe un slice de CityWeather y devuelve
// los valores agregados en un AggregationResult.
func Aggregate(results []CityWeather) AggregationResult {

	// Si no hay resultados se evita la división por cero (NaN)
	// y devuelve valores por defecto válidos para JSON.
	if len(results) == 0 {
		return AggregationResult{
			Averages: map[string]float64{
				"temperature": 0,
				"humidity":    0,
				"wind_speed":  0,
			},
			Extremes:    map[string]string{},
			ByCondition: map[string][]string{},
		}
	}

	var totalTemp float64
	var totalHum int
	var totalWind float64

	// Inicializamos extremos con la primera ciudad
	hottest := results[0]
	coldest := results[0]
	windiest := results[0]

	// Mapa para agrupar ciudades por condición climática
	byCondition := make(map[string][]string)

	// Recorre todas las ciudades
	for i, r := range results {
		totalTemp += r.Temperature
		totalHum += r.Humidity
		totalWind += r.WindSpeed

		// A partir del segundo elemento evalua extremos
		if i > 0 {
			if r.Temperature > hottest.Temperature {
				hottest = r
			}
			if r.Temperature < coldest.Temperature {
				coldest = r
			}
			if r.WindSpeed > windiest.WindSpeed {
				windiest = r
			}
		}

		// Agrupación por condición
		byCondition[r.Condition] = append(byCondition[r.Condition], r.CityName)
	}

	// Se calculan los promedios
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
