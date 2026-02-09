// URL base del backend
const BASE_URL = "http://localhost:8080/api";

// Trae la lista de ciudades disponibles
export async function getCities() {
  const res = await fetch(`${BASE_URL}/cities`);
  if (!res.ok) {
    throw new Error("No se pudieron cargar las ciudades");
  }
  return res.json();
}

// Compara el clima entre varias ciudades
export async function compareWeather(cityIds) {
  const res = await fetch(`${BASE_URL}/weather/compare`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ city_ids: cityIds })
  });

  if (!res.ok) {
    throw new Error("Error comparando ciudades");
  }

  return res.json();
}
