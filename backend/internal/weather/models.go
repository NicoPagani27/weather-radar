package weather

// CityWeather representa el clima procesado de una ciudad.
// Hago esto para ordenar los datos.
type CityWeather struct {
	CityID      string
	CityName    string
	Temperature float64
	Humidity    int
	WindSpeed   float64
	Condition   string

	// Error contiene un mensaje si la consulta falló.
	// Si está vacío, la ciudad se procesó correctamente.
	Error string
}