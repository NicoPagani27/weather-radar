package weather

import "weather-radar/backend/internal/cities"

// FetchCityWeather consulta el clima de una ciudad y
// envía el resultado al channel (canal).
// Esta función se ejecuta como goroutine.

func FetchCityWeather(city cities.City, ch chan<- CityWeather) {
	data, err := GetCurrentWeather(city.Latitude, city.Longitude)
	if err != nil {
		// En este caso no enviamos nada
		// (podría mejorarse con manejo de errores)
		return
	}

	result := CityWeather{
		CityID:      city.ID,
		CityName:    city.Name,
		Temperature: data.Current.Temperature,
		Humidity:    data.Current.Humidity,
		WindSpeed:   data.Current.WindSpeed,
		Condition:   WeatherCodeToCondition(data.Current.WeatherCode),
	}

	// fan-in(junto resultados) ---> se envía el resultado al channel compartido 
	ch <- result
}
