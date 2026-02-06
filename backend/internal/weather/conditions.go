package weather

// WeatherCodeToCondition traduce el weather_code de Open-Meteo
// a una condición climática más entendible.
// Aparece en WMO Weather interpretation codes (WW).

func WeatherCodeToCondition(code int) string {
	switch {
	case code == 0:
		return "soleado"
	case code >= 1 && code <= 3:
		return "parcialmente nublado"
	case code >= 45 && code <= 48:
		return "neblina"
	case code >= 51 && code <= 67:
		return "lluvia"
	case code >= 71 && code <= 77:
		return "nieve"
	case code >= 80 && code <= 99:
		return "tormenta"
	default:
		return "desconocido"
	}
}
