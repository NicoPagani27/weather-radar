package weather

import (
	"context"
	"weather-radar/backend/internal/cities"
)

// FetchCityWeather consulta el clima de una ciudad y
// envía el resultado al channel (canal).
// Esta función se ejecuta como goroutine.
func FetchCityWeather(city cities.City, ch chan CityWeather) {
	data, err := GetCurrentWeather(city.Latitude, city.Longitude)
	if err != nil {
		ch <- CityWeather{
			CityID:   city.ID,
			CityName: city.Name,
			Error:    err.Error(),
		}
		return
	}

	ch <- CityWeather{
		CityID:      city.ID,
		CityName:    city.Name,
		Temperature: data.Current.Temperature,
		Humidity:    data.Current.Humidity,
		WindSpeed:   data.Current.WindSpeed,
		WeatherCode: data.Current.WeatherCode,
		Condition:   WeatherCodeToCondition(data.Current.WeatherCode),
	}
}

// FetchCityWeatherWithContext consulta el clima con timeout usando context.
func FetchCityWeatherWithContext(ctx context.Context, city cities.City, ch chan CityWeather) {
	// Canal para recibir el resultado de la goroutine
	type result struct {
		data OpenMeteoResponse
		err  error
	}
	doneChan := make(chan result, 1)

	// Ejecutar la consulta en una goroutine
	go func() {
		data, err := GetCurrentWeather(city.Latitude, city.Longitude)
		doneChan <- result{data: data, err: err}
	}()

	// Esperar resultado o timeout
	select {
	case res := <-doneChan:
		if res.err != nil {
			ch <- CityWeather{
				CityID:   city.ID,
				CityName: city.Name,
				Error:    res.err.Error(),
			}
			return
		}

		ch <- CityWeather{
			CityID:      city.ID,
			CityName:    city.Name,
			Temperature: res.data.Current.Temperature,
			Humidity:    res.data.Current.Humidity,
			WindSpeed:   res.data.Current.WindSpeed,
			WeatherCode: res.data.Current.WeatherCode,
			Condition:   WeatherCodeToCondition(res.data.Current.WeatherCode),
		}
	case <-ctx.Done():
		// Timeout o cancelación
		ch <- CityWeather{
			CityID:   city.ID,
			CityName: city.Name,
			Error:    "timeout consultando clima",
		}
	}
}

