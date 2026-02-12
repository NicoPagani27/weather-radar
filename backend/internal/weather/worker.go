package weather

import "weather-radar/backend/internal/cities"

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
