package cities
// City representa una ciudad disponible para consultar clima.
// tiene los datos necesarios para usar la API.
type City struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}