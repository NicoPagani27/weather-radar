<script>
  // Datos de una ciudad individual
  export let city;
</script>

<div class="bg-white rounded-3xl shadow-lg hover:shadow-xl transition-all duration-300 min-h-[240px]" style="padding: 1.25rem;">
  <!-- Nombre de la ciudad -->
  <div>
    <h3 class="text-lg font-semibold text-[#2c3e50]">{city.CityName}</h3>
  </div>

  <!-- Temperatura principal -->
  <div style="margin-top: 0.75rem; margin-bottom: 0.75rem;">
    <span class="text-5xl font-light text-[#2c3e50] leading-none">{city.Temperature}°</span>
  </div>

  <!-- Detalles del clima -->
  <div class="space-y-2">
    <div class="flex items-center justify-between" style="padding: 0.3rem 0; border-bottom: 1px solid rgba(44, 62, 80, 0.1);">
      <span class="text-base text-[#2c3e50]">Humedad</span>
      <span class="text-base text-[#2c3e50] font-medium">{city.Humidity}%</span>
    </div>

    <div class="flex items-center justify-between" style="padding: 0.3rem 0; border-bottom: 1px solid rgba(44, 62, 80, 0.1);">
      <span class="text-base text-[#2c3e50]">Viento</span>
      <span class="text-base text-[#2c3e50] font-medium">{city.WindSpeed}km/h</span>
    </div>
  </div>
</div>
