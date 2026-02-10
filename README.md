# Weather Radar

Aplicación web para consultar y comparar el clima de diferentes ciudades en tiempo real. El backend está hecho en Go y consulta datos de Open-Meteo.

## ¿Qué hace?

Seleccionás una o varias ciudades y obtenés información del clima actual como temperatura, humedad y velocidad del viento. También podés comparar el clima entre ciudades para ver cuál tiene mejor (o peor) tiempo.

## Estructura del proyecto

```
.
├── backend/           # API en Go
│   ├── cmd/
│   │   └── server/    # Punto de entrada del servidor
│   └── internal/
│       ├── cities/    # Manejo de ciudades disponibles
│       ├── handlers/  # Endpoints HTTP
│       └── weather/   # Lógica para consultar clima
├── frontend/          # Interfaz web (Svelte)
└── postman/          # Colección para probar la API
```

## Tecnologías

- **Backend**: Go 1.25 con Gin
- **API externa**: [Open-Meteo](https://open-meteo.com/) (sin necesidad de API key)
- **Frontend**: Svelte

## Cómo levantar el proyecto

### Backend

Primero instalá las dependencias:

```bash
cd backend
go mod download
```

Después inicia el servidor:

```bash
go run cmd/server/main.go
```

El servidor va a estar funcionando en `http://localhost:8080`.

### Endpoints disponibles

**Obtener ciudades**
```
GET /api/cities
```

Devuelve la lista de ciudades disponibles con su ID, nombre y coordenadas.

**Clima de una ciudad**
```
GET /api/weather/:cityId
```

Ejemplo: `/api/weather/bsas` te da el clima actual de Buenos Aires.

**Comparar ciudades**
```
POST /api/weather/compare
Content-Type: application/json

{
  "city_ids": ["bsas", "cordoba", "mendoza"]
}
```

Compara el clima de varias ciudades y te dice cuál tiene mejor condiciones. Usa goroutines para consultar todas en paralelo, así que es rápido incluso con muchas ciudades.

## Testing con Postman

En la carpeta `postman/` hay una colección con ejemplos de las requests. Importala en Postman y podés probar todos los endpoints fácilmente.

## Notas técnicas

- El backend maneja CORS para poder ser consumido desde el frontend que corre en otro puerto
- Usa goroutines y canales para consultar múltiples APIs en paralelo cuando se comparan ciudades
- Timeout de 3 segundos en las llamadas externas para que no se cuelgue si Open-Meteo anda lento
- Los códigos de clima vienen de Open-Meteo y hay que interpretarlos (0 = despejado, 1-3 = nublado, etc.)


-----
Desarrollado como proyecto personal para aprender Go y practicar concurrencia.