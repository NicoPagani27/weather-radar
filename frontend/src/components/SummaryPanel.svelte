<script>
  // Información agregada del grupo de ciudades
  export let summary;

  // Helper function para formatear números de forma segura
  function formatNumber(value, decimals = 1) {
    if (value == null || isNaN(value) || !isFinite(value)) {
      return '--';
    }
    return value.toFixed(decimals);
  }

  // Verificar si tenemos datos válidos
  $: hasValidData = summary?.averages &&
                    !isNaN(summary.averages.temperature) &&
                    !isNaN(summary.averages.humidity) &&
                    !isNaN(summary.averages.wind_speed);
</script>

<div class="bg-white rounded-3xl shadow-lg hover:shadow-xl transition-all duration-300" style="padding: 2rem;">
  <!-- Header -->
  <div style="margin-bottom: 2rem;">
    <h3 class="text-xl font-semibold text-[#2c3e50]" style="margin-bottom: 0.5rem;">Resumen General</h3>
  </div>

  <!-- Promedios -->
  <div style="margin-bottom: 2.5rem;">
    <h4 class="text-base font-semibold text-[#2c3e50] uppercase tracking-wide" style="font-size: 0.75rem; color: #219ebc; margin-bottom: 1.5rem;">Valores Promedio</h4>
    
    <div class="grid grid-cols-1 md:grid-cols-3" style="gap: 2rem;">
      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-medium text-gray-500 uppercase tracking-wide">Temperatura</div>
        <div class="text-4xl font-light text-[#2c3e50]">
          {formatNumber(summary?.averages?.temperature, 1)}<span class="text-lg text-[#219ebc]">°C</span>
        </div>
      </div>

      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-medium text-gray-500 uppercase tracking-wide">Humedad</div>
        <div class="text-4xl font-light text-[#2c3e50]">
          {formatNumber(summary?.averages?.humidity, 0)}<span class="text-lg text-[#219ebc]">%</span>
        </div>
      </div>

      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-medium text-gray-500 uppercase tracking-wide">Viento</div>
        <div class="text-4xl font-light text-[#2c3e50]">
          {formatNumber(summary?.averages?.wind_speed, 1)}<span class="text-lg text-[#219ebc]">km/h</span>
        </div>
      </div>
    </div>
  </div>

  <!-- Extremos -->
  <div>
    <h4 class="text-base font-semibold text-[#2c3e50] uppercase tracking-wide" style="font-size: 0.75rem; color: #219ebc; margin-bottom: 1.5rem;">Valores Extremos</h4>
    
    <div class="grid grid-cols-1 md:grid-cols-3" style="gap: 2rem;">
      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-semibold text-[#fb8500] uppercase tracking-wide">Más Caliente</div>
        <div class="text-2xl font-semibold text-[#2c3e50]">{summary?.extremes?.hottest || '--'}</div>
      </div>

      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-semibold text-[#219ebc] uppercase tracking-wide">Más Frío</div>
        <div class="text-2xl font-semibold text-[#2c3e50]">{summary?.extremes?.coldest || '--'}</div>
      </div>

      <div style="display: flex; flex-direction: column; gap: 0.75rem;">
        <div class="text-xs font-semibold text-[#023047] uppercase tracking-wide">Más Ventoso</div>
        <div class="text-2xl font-semibold text-[#2c3e50]">{summary?.extremes?.windiest || '--'}</div>
      </div>
    </div>
  </div>
</div>
