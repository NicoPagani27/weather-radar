<script>
  import { onMount } from "svelte";
  import { getCities, compareWeather } from "$lib/api";

  import CityCard from "../components/CityCard.svelte";
  import SummaryPanel from "../components/SummaryPanel.svelte";

  let cities = [];
  let selectedCities = [];
  let result = null;
  let loading = false;
  let error = "";
  let showCitySelector = false;
  
  // Animación de typewriter
  let displayedTitle = "";
  let showCursor = true;
  const fullTitle = "Weather Radar";
  
  // Función para mostrar selector y hacer scroll
  function showCitySelectorAndScroll() {
    showCitySelector = true;
    setTimeout(() => {
      document.getElementById('city-selector')?.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }, 100);
  }
  
  // Cargar ciudades al iniciar la página
  onMount(async () => {
    // Animación de typewriter para el título
    for (let i = 0; i <= fullTitle.length; i++) {
      displayedTitle = fullTitle.slice(0, i);
      await new Promise(resolve => setTimeout(resolve, 100)); // Ajusta la velocidad del typewriter aquí
    }
    
    // Ocultar el cursor cuando termine la animación
    showCursor = false;
    
    try {
      cities = await getCities();
    } catch (e) {
      error = e.message;
    }
  });

  // Ejecuta la comparación
  async function handleCompare() {
    if (selectedCities.length === 0) return;

    loading = true;
    error = "";
    result = null;

    try {
      result = await compareWeather(selectedCities);
    } catch (e) {
      error = e.message;
    } finally {
      loading = false;
    }
  }
</script>

<!-- Hero Section -->
<div class="h-[70vh] flex items-center justify-center">
  <div class="text-center px-4">
    <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold text-white mb-4 drop-shadow-2xl">
      {displayedTitle}{#if showCursor}<span class="typewriter-cursor">|</span>{/if}
    </h1>
    <p class="text-lg md:text-xl text-white/90 font-light animate-fade-in-delay" style="margin-bottom: 3rem;">
      Compará el clima de múltiples ciudades en tiempo real
    </p>
    
    <button class="cssbuttons-io animate-fade-in-delay-2" on:click={showCitySelectorAndScroll}>
      <span>Seleccionar Ciudades</span>
    </button>
  </div>
</div>

<!-- Content Section -->
<div class="w-full pb-4 flex flex-col items-center px-6 md:px-10 lg:px-16">
  {#if error}
    <div class="bg-red-500/90 backdrop-blur-sm text-white px-6 py-4 rounded-2xl mb-6 shadow-lg">
      <p class="font-medium">⚠️ {error}</p>
    </div>
  {/if}

  {#if showCitySelector}
    <!-- Selector de ciudades -->
    <div id="city-selector" class="animate-slide-up flex flex-col items-center w-full">
    <div class="grid grid-cols-4 gap-3 mb-16 justify-items-center w-full max-w-[1100px]">
      {#each cities as city, index}
        <label 
          class="cursor-pointer animate-city-appear w-full max-w-[250px]" 
          class:col-start-2={index === 4 && cities.length === 6}
          class:col-span-1={true}
          style="animation-delay: {index * 0.05}s;"
        >
          <input
            type="checkbox"
            value={city.id}
            bind:group={selectedCities}
            class="hidden"
          />
          <div class="city-card">
            <span class="city-name">{city.name}</span>
          </div>
        </label>
      {/each}
    </div>

    <div style="margin-top: 4rem;"></div>

    <button
      class="cssbuttons-io disabled:opacity-50 disabled:cursor-not-allowed"
      on:click={handleCompare}
      disabled={loading || selectedCities.length === 0}
    >
      <span>
        {#if loading}
          Comparando...
        {:else}
          Comparar Clima
        {/if}
      </span>
    </button>
  </div>
  {/if}

  {#if result}
    <!-- Tarjetas de ciudades -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-10" style="margin-top: 6rem; margin-bottom: 4rem; width: calc(100% - 4rem); max-width: 1400px;">
      {#each result.cities as city}
        <CityCard {city} />
      {/each}
    </div>

    <!-- Panel de resumen -->
    <div style="width: calc(100% - 4rem); max-width: 1400px; margin-bottom: 4rem;">
      <SummaryPanel summary={result.summary} />
    </div>
  {/if}
</div>

<style>
  :global(body) {
    overflow-x: hidden;
  }

  @keyframes fade-in {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes slide-up {
    from {
      opacity: 0;
      transform: translateY(60px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes city-appear {
    from {
      opacity: 0;
      transform: translateY(20px) scale(0.9);
    }
    to {
      opacity: 1;
      transform: translateY(0) scale(1);
    }
  }

  @keyframes blink {
    0%, 100% {
      opacity: 1;
    }
    50% {
      opacity: 0;
    }
  }

  .typewriter-cursor {
    animation: blink 1s infinite;
  }

  .animate-fade-in {
    animation: fade-in 1s ease-out;
  }

  .animate-fade-in-delay {
    animation: fade-in 1s ease-out 0.3s both;
  }

  .animate-fade-in-delay-2 {
    animation: fade-in 1s ease-out 0.8s both;
  }

  .animate-slide-up {
    animation: slide-up 1s ease-out both;
  }

  .animate-city-appear {
    animation: city-appear 0.6s ease-out both;
  }

  /* Botón personalizado */
  .cssbuttons-io {
    position: relative;
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    font-weight: 300;
    font-size: 18px;
    letter-spacing: 0.05em;
    border-radius: 0.8em;
    cursor: pointer;
    border: none;
    background: #219ebc;
    color: white;
    overflow: hidden;
    box-shadow: 0 4px 15px rgba(33, 158, 188, 0.3);
  }

  .cssbuttons-io span {
    position: relative;
    z-index: 10;
    transition: color 0.4s;
    display: inline-flex;
    align-items: center;
    padding: 0.8em 1.2em 0.8em 1.05em;
  }

  .cssbuttons-io::before,
  .cssbuttons-io::after {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
  }

  .cssbuttons-io::before {
    content: "";
    background: #0d4863;
    width: 120%;
    left: -10%;
    transform: skew(30deg);
    transition: transform 0.4s cubic-bezier(0.3, 1, 0.8, 1);
  }

  .cssbuttons-io:hover::before {
    transform: translate3d(100%, 0, 0);
  }

  .cssbuttons-io:active {
    transform: scale(0.95);
  }

  /* Botones de ciudad */
  .city-card {
    position: relative;
    padding: 1.3em 0;
    font-size: 12px;
    text-transform: uppercase;
    letter-spacing: 2.5px;
    font-weight: 500;
    color: #000;
    background-color: #fff;
    border: none;
    border-radius: 45px;
    box-shadow: 0px 8px 15px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease 0s;
    cursor: pointer;
    outline: none;
    text-align: center;
    width: 100%;
  }

  .city-card:hover {
    background-color: #8ecae6;
    box-shadow: 0px 15px 20px rgba(142, 202, 230, 0.4);
    color: #fff;
    transform: translateY(-7px);
  }

  .city-card:active {
    transform: translateY(-1px);
  }

  .city-name {
    font-size: inherit;
    font-weight: inherit;
    color: inherit;
    transition: color 0.3s ease;
  }

  input[type="checkbox"]:checked + .city-card {
    background-color: #8ecae6;
    box-shadow: 0px 15px 20px rgba(142, 202, 230, 0.4);
    color: #fff;
    transform: translateY(-7px);
  }

  input[type="checkbox"]:checked + .city-card .city-name {
    color: #fff;
  }

  input[type="checkbox"]:checked + .city-card:hover {
    transform: translateY(-9px);
  }
</style>  