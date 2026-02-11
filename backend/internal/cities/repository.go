package cities

// cities almacena el mapa de ciudades para búsquedas O(1)
var citiesMap map[string]City
var citiesList []City

func init() {
	// Inicializar ciudades una sola vez
	citiesList = []City{
		{
			ID:        "cordoba",
			Name:      "Córdoba",
			Latitude:  -31.42,
			Longitude: -64.18,
		},
		{
			ID:        "buenos_aires",
			Name:      "Buenos Aires",
			Latitude:  -34.61,
			Longitude: -58.38,
		},
		{
			ID:        "sao_paulo",
			Name:      "São Paulo",
			Latitude:  -23.55,
			Longitude: -46.63,
		},
		{
			ID:        "new_york",
			Name:      "New York",
			Latitude:  40.71,
			Longitude: -74.01,
		},
		{
			ID:        "rosario",
			Name:      "Rosario",
			Latitude:  -32.95,
			Longitude: -60.66,
		},
		{
			ID:        "ciudad_de_mexico",
			Name:      "Ciudad de México",
			Latitude:  19.43,
			Longitude: -99.13,
		},
		{
			ID:        "madrid",
			Name:      "Madrid",
			Latitude:  40.42,
			Longitude: -3.70,
		},
		{
			ID:        "londres",
			Name:      "Londres",
			Latitude:  51.51,
			Longitude: -0.13,
		},
		{
			ID:        "paris",
			Name:      "París",
			Latitude:  48.86,
			Longitude: 2.35,
		},
		{
			ID:        "tokio",
			Name:      "Tokio",
			Latitude:  35.68,
			Longitude: 139.65,
		},
		{
			ID:        "sidney",
			Name:      "Sídney",
			Latitude:  -33.87,
			Longitude: 151.21,
		},
		{
			ID:        "toronto",
			Name:      "Toronto",
			Latitude:  43.65,
			Longitude: -79.38,
		},
		{
			ID:        "rio_de_janeiro",
			Name:      "Río de Janeiro",
			Latitude:  -22.91,
			Longitude: -43.17,
		},
		{
			ID:        "berlin",
			Name:      "Berlín",
			Latitude:  52.52,
			Longitude: 13.41,
		},
		{
			ID:        "lima",
			Name:      "Lima",
			Latitude:  -12.05,
			Longitude: -77.04,
		},
	}

	// Crear mapa para búsquedas rápidas
	citiesMap = make(map[string]City)
	for _, city := range citiesList {
		citiesMap[city.ID] = city
	}
}

// GetAll devuelve la lista de ciudades disponibles.
// Las ciudades están hardcodeadas
// para evitar el uso de base de datos (precargadas).
func GetAll() []City {
	return citiesList
}

// GetByID busca una ciudad por su ID.
// Devuelve la ciudad y un booleano indicando si fue encontrada.
// Usa un map para búsqueda O(1) en lugar de O(n).
func GetByID(id string) (City, bool) {
	city, found := citiesMap[id]
	return city, found
}