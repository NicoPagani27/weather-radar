package cities

func GetAll() []City {
	return []City{
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
	}
}
