package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Cliente HTTP reutilizable (singleton)
var httpClient = &http.Client{
	Timeout: 3 * time.Second,
}

// OpenMeteoResponse representa únicamente
// los campos que nos interesan de la respuesta de la API.
type OpenMeteoResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
		Humidity    int     `json:"relative_humidity_2m"`
		WindSpeed   float64 `json:"wind_speed_10m"`
		WeatherCode int     `json:"weather_code"`
	} `json:"current"`
}

// GetCurrentWeather consulta la API usando latitud y longitud
// y devuelve el clima actual.
func GetCurrentWeather(lat, lon float64) (OpenMeteoResponse, error) {
	// Validar coordenadas
	if lat < -90 || lat > 90 {
		return OpenMeteoResponse{}, fmt.Errorf("latitud inválida: %f (debe estar entre -90 y 90)", lat)
	}
	if lon < -180 || lon > 180 {
		return OpenMeteoResponse{}, fmt.Errorf("longitud inválida: %f (debe estar entre -180 y 180)", lon)
	}
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code&timezone=auto",
		lat,
		lon,
	)

	// Timeout de 3 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return OpenMeteoResponse{}, fmt.Errorf("error creando request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return OpenMeteoResponse{}, fmt.Errorf("error llamando API: %w", err)
	}
	defer resp.Body.Close()

	var data OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return OpenMeteoResponse{}, fmt.Errorf("error decodificando JSON: %w", err)
	}

	return data, nil
}
