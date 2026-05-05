<template>
  <header class="hero-section-gradient from-purple-600 to-indigo-600 py-16">
    <div class="hero-container">
      <h1 class="hero-title">Our Games</h1>
      <p class="hero-subtitle text-purple-50">Discover games made for TIC-80 and other platforms, with versions playable directly in your browser!</p>
    </div>
  </header>

  <main class="main-container py-12 px-4 mt-0">
    <div class="flex flex-wrap gap-4 mb-8 justify-center filter-buttons pb-4">
      <button
        v-for="f in filters"
        :key="f"
        :class="['filter-btn', { active: activeFilter === f }]"
        @click="activeFilter = f"
      >
        {{ f.charAt(0).toUpperCase() + f.slice(1) }}
      </button>
    </div>

    <div class="grid grid-cols-1 gap-8">
      <template v-for="{ software, releases } in softwares" :key="software.name">
        <div v-show="activeFilter === 'all' || software.status === activeFilter" class="software-card group">
          <div class="software-card-content">
            <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
              <div class="flex-grow">
                <div class="flex items-center gap-3 mb-2">
                  <RouterLink :to="`/catalog/${software.name}`" class="hover:text-purple-600 transition-colors">
                    <h2 class="software-title mb-0">{{ software.title }}</h2>
                  </RouterLink>
                  <span :class="`status-badge status-${software.status}`">{{ software.status }}</span>
                </div>
                <p class="software-desc max-w-2xl">{{ software.desc }}</p>
                <div class="software-meta">
                  <span>Author: <span class="text-gray-900 font-medium">{{ software.author }}</span></span>
                  <span class="mx-2 text-gray-300">•</span>
                  <span>Platform: <span class="text-gray-900 font-medium">{{ software.platform }}</span></span>
                </div>
              </div>

              <div class="flex flex-wrap gap-2 items-center">
                <a
                  v-if="getLatestStable(releases)?.htmlFolderPath"
                  :href="getLatestStable(releases).htmlFolderPath"
                  target="_blank"
                  class="btn-play-sm"
                >▶ Play</a>
                <RouterLink :to="`/catalog/${software.name}`" class="btn-info-sm">
                  ℹ More Info
                </RouterLink>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import softwareApi from '../../api/software.api'
import type { Release, SoftwareEntry } from '../../lib/interfaces/software.interface'

const filters = ['all', 'released', 'demo', 'development', 'archived']
const activeFilter = ref('all')
const softwares = ref<SoftwareEntry[]>([])

function getLatestStable(releases: Release[]): Release {
  return releases
    .filter(r => !r.version.startsWith('dev-'))
    .sort((a, b) => new Date(b.UpdatedAt).getTime() - new Date(a.UpdatedAt).getTime())[0]
}

onMounted(async () => {
  try {
    softwares.value = await softwareApi.index()
  } catch (e) {
    console.error('Failed to load catalog:', e)
  }
})
</script>

<style scoped>
.filter-btn {
  @apply px-4 py-2 rounded-full border-2 border-purple-600 text-purple-600 font-bold transition-all hover:bg-purple-600 hover:text-white;
}
.filter-btn.active {
  @apply bg-purple-600 text-white;
}
.status-badge {
  @apply px-2 py-0.5 rounded text-xs font-bold uppercase;
}
.status-released { @apply bg-green-100 text-green-800 border border-green-200; }
.status-demo { @apply bg-blue-100 text-blue-800 border border-blue-200; }
.status-development { @apply bg-yellow-100 text-yellow-800 border border-yellow-200; }
.status-archived { @apply bg-gray-100 text-gray-800 border border-gray-200; }
.software-card {
  @apply bg-white shadow-lg relative rounded-xl overflow-hidden hover:shadow-2xl transition-shadow duration-300;
}
.software-card-content {
  @apply p-6;
}
.software-title {
  @apply text-2xl font-bold mb-0;
}
.software-desc {
  @apply text-gray-600 mb-3;
}
.software-meta {
  @apply text-sm text-gray-500;
}
.btn-base-sm {
  @apply inline-flex items-center justify-center font-bold py-2.5 px-5 rounded-xl text-sm transition-all active:scale-95;
}
.btn-play-sm { @apply btn-base-sm bg-purple-600 hover:bg-purple-700 text-white shadow-lg shadow-purple-600/20; }
.btn-info-sm { @apply btn-base-sm bg-gray-100 hover:bg-gray-200 text-gray-700; }
</style>
