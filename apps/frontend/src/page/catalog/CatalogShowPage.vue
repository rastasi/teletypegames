<template>
  <div class="bg-gray-50 min-h-screen pb-20">
    <header class="hero-section-gradient from-purple-700 to-indigo-800 py-12 md:py-20">
      <div class="hero-container">
        <nav class="mb-8">
          <RouterLink to="/catalog" class="text-purple-200 hover:text-white transition-colors flex items-center gap-2 font-medium">
            ← Back to Catalog
          </RouterLink>
        </nav>

        <div v-if="software" class="flex flex-col md:flex-row md:items-end justify-between gap-6">
          <div class="flex-grow">
            <div class="flex items-center gap-3 mb-4">
              <span :class="`status-badge status-${software.status} px-3 py-1`">{{ software.status }}</span>
              <span class="text-purple-300 font-mono text-sm tracking-widest uppercase">{{ software.platform }}</span>
            </div>
            <h1 class="text-4xl md:text-6xl font-black text-white mb-4 leading-tight">{{ software.title }}</h1>
            <p class="text-xl text-purple-100 max-w-3xl leading-relaxed">{{ software.desc }}</p>
          </div>
        </div>
      </div>
    </header>

    <main v-if="software" class="main-container -mt-10">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">

        <!-- Sidebar Info -->
        <div class="lg:col-span-1 space-y-6">
          <div class="bg-white rounded-3xl shadow-xl p-8 border border-gray-100">
            <h2 class="text-xl font-bold text-gray-900 mb-6 flex items-center gap-2">
              <span class="text-purple-600">ℹ</span> Project Info
            </h2>
            <dl class="space-y-4">
              <div>
                <dt class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1">Author</dt>
                <dd class="text-gray-900 font-medium">{{ software.author }}</dd>
              </div>
              <div>
                <dt class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1">Platform</dt>
                <dd class="text-gray-900 font-medium">{{ software.platform }}</dd>
              </div>
              <div v-if="software.license">
                <dt class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1">License</dt>
                <dd class="text-gray-900 font-medium">{{ software.license }}</dd>
              </div>
              <div v-for="link in software.externalLinks" :key="link.url">
                <dt class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-1">{{ link.label }}</dt>
                <dd>
                  <a :href="link.url" target="_blank" class="text-purple-600 hover:text-purple-800 font-bold break-all">
                    {{ link.url }} ↗
                  </a>
                </dd>
              </div>
            </dl>
          </div>
        </div>

        <!-- Main Content Area -->
        <div class="lg:col-span-2 space-y-8">
          <div v-if="software.story" class="bg-white rounded-3xl shadow-xl p-8 border border-gray-100">
            <h2 class="text-xl font-bold text-gray-900 mb-4">About</h2>
            <p class="text-gray-700 leading-relaxed whitespace-pre-line">{{ software.story }}</p>
          </div>

          <div v-if="latestStable" class="bg-gradient-to-br from-indigo-600 to-purple-700 rounded-3xl shadow-xl p-8 text-white">
            <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
              <div>
                <div class="flex items-center gap-2 mb-2">
                  <span class="px-2 py-0.5 bg-white/20 text-white text-[10px] font-black uppercase rounded tracking-widest">Latest Stable</span>
                </div>
                <h2 class="text-4xl font-black mb-1">{{ latestStable.version }}</h2>
                <p class="text-indigo-100 text-sm">Released: {{ formatDateTime(latestStable.UpdatedAt) }}</p>
              </div>

              <div class="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-2 xl:grid-cols-4 gap-2 w-full md:w-auto">
                <a v-if="latestStable.htmlFolderPath" :href="latestStable.htmlFolderPath" target="_blank" class="flex items-center justify-center py-3 px-4 bg-white text-indigo-900 font-bold rounded-xl transition-all hover:bg-indigo-50 shadow-lg shadow-black/10 text-sm whitespace-nowrap">▶ Play Now</a>
                <a v-if="latestStable.cartridgePath" :href="latestStable.cartridgePath" target="_blank" class="flex items-center justify-center py-3 px-4 bg-indigo-500/30 border border-indigo-400/30 text-white font-bold rounded-xl transition-all hover:bg-indigo-500/40 text-sm whitespace-nowrap">💾 Download</a>
                <a v-if="latestStable.sourcePath" :href="latestStable.sourcePath" target="_blank" class="flex items-center justify-center py-3 px-4 bg-indigo-500/30 border border-indigo-400/30 text-white font-bold rounded-xl transition-all hover:bg-indigo-500/40 text-sm whitespace-nowrap">📄 Source</a>
                <a v-if="latestStable.docsFolderPath" :href="latestStable.docsFolderPath" target="_blank" class="flex items-center justify-center py-3 px-4 bg-indigo-500/30 border border-indigo-400/30 text-white font-bold rounded-xl transition-all hover:bg-indigo-500/40 text-sm whitespace-nowrap">📖 Docs</a>
              </div>
            </div>
          </div>

          <div class="bg-white rounded-3xl shadow-xl overflow-hidden border border-gray-100">
            <div class="p-8 border-b border-gray-100">
              <h2 class="text-2xl font-bold text-gray-900">All Releases</h2>
            </div>
            <div class="p-0">
              <!-- Desktop Table -->
              <div class="hidden md:block">
                <table class="w-full text-left">
                  <thead class="bg-gray-50 text-gray-500 text-xs font-bold uppercase tracking-widest">
                    <tr>
                      <th class="px-8 py-4">Version</th>
                      <th class="px-8 py-4">Release Date</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100">
                    <tr v-for="release in stableReleases" :key="release.version" class="hover:bg-gray-50/50 transition-colors">
                      <td class="px-8 py-4">
                        <div class="font-bold text-gray-900 text-lg mb-2">{{ release.version }}</div>
                        <div class="flex flex-wrap gap-4">
                          <a v-if="release.htmlFolderPath" :href="release.htmlFolderPath" target="_blank" class="text-purple-600 hover:text-purple-900 font-bold text-sm flex items-center gap-1">▶ Play</a>
                          <a v-if="release.cartridgePath" :href="release.cartridgePath" target="_blank" class="text-blue-600 hover:text-blue-900 font-bold text-sm flex items-center gap-1">💾 Download</a>
                          <a v-if="release.sourcePath" :href="release.sourcePath" target="_blank" class="text-green-600 hover:text-green-900 font-bold text-sm flex items-center gap-1">📄 Source</a>
                          <a v-if="release.docsFolderPath" :href="release.docsFolderPath" target="_blank" class="text-yellow-600 hover:text-yellow-900 font-bold text-sm flex items-center gap-1">📖 Docs</a>
                        </div>
                      </td>
                      <td class="px-8 py-4 text-gray-500 align-top pt-5">{{ formatDateTime(release.UpdatedAt) }}</td>
                    </tr>

                    <tr v-if="devReleases.length > 0" class="bg-yellow-50/30">
                      <td colspan="2" class="px-8 py-3 text-[10px] font-black text-yellow-700 uppercase tracking-tighter">Development Versions</td>
                    </tr>

                    <tr v-for="release in devReleases" :key="release.version" class="bg-yellow-50/10 hover:bg-yellow-50/30 transition-colors">
                      <td class="px-8 py-4">
                        <div class="font-bold text-yellow-800 text-lg mb-2">{{ release.version }}</div>
                        <div class="flex flex-wrap gap-4">
                          <a v-if="release.htmlFolderPath" :href="release.htmlFolderPath" target="_blank" class="text-purple-600 hover:text-purple-900 font-bold text-sm flex items-center gap-1">▶ Play</a>
                          <a v-if="release.cartridgePath" :href="release.cartridgePath" target="_blank" class="text-blue-600 hover:text-blue-900 font-bold text-sm flex items-center gap-1">💾 Download</a>
                        </div>
                      </td>
                      <td class="px-8 py-4 text-gray-500 align-top pt-5">{{ formatDateTime(release.UpdatedAt) }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>

              <!-- Mobile View -->
              <div class="md:hidden divide-y divide-gray-100">
                <div
                  v-for="release in allReleases"
                  :key="release.version"
                  :class="['p-6', release.version.startsWith('dev-') ? 'bg-yellow-50/20' : '']"
                >
                  <div class="flex justify-between items-start mb-4">
                    <div>
                      <div class="font-bold text-gray-900 text-lg">{{ release.version }}</div>
                      <div class="text-gray-500 text-xs mt-1">{{ formatDateTime(release.UpdatedAt) }}</div>
                    </div>
                    <span v-if="release.version.startsWith('dev-')" class="px-2 py-0.5 bg-yellow-100 text-yellow-800 text-[10px] font-black uppercase rounded">Dev</span>
                  </div>
                  <div class="grid grid-cols-2 gap-2">
                    <a v-if="release.htmlFolderPath" :href="release.htmlFolderPath" target="_blank" class="flex items-center justify-center py-2 bg-purple-100 text-purple-700 font-bold rounded-lg text-sm">Play</a>
                    <a v-if="release.cartridgePath" :href="release.cartridgePath" target="_blank" class="flex items-center justify-center py-2 bg-blue-100 text-blue-700 font-bold rounded-lg text-sm">Download</a>
                    <a v-if="release.sourcePath" :href="release.sourcePath" target="_blank" class="flex items-center justify-center py-2 bg-green-100 text-green-700 font-bold rounded-lg text-sm">Source</a>
                    <a v-if="release.docsFolderPath" :href="release.docsFolderPath" target="_blank" class="flex items-center justify-center py-2 bg-yellow-100 text-yellow-700 font-bold rounded-lg text-sm">Docs</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <main v-else class="main-container -mt-10">
      <div class="bg-white rounded-3xl shadow-xl p-12 text-center text-gray-400">
        {{ error || 'Loading...' }}
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { formatDateTime } from '../../lib/dateFormat'
import softwareApi from '../../api/software.api'
import type { Release, Software } from '../../lib/interfaces/software.interface'

const route = useRoute()
const software = ref<Software | null>(null)
const releases = ref<Release[]>([])
const error = ref<string | null>(null)

const stableReleases = computed(() =>
  releases.value
    .filter(r => !r.version.startsWith('dev-'))
    .sort((a, b) => new Date(b.UpdatedAt).getTime() - new Date(a.UpdatedAt).getTime())
)

const devReleases = computed(() =>
  releases.value
    .filter(r => r.version.startsWith('dev-'))
    .sort((a, b) => new Date(b.UpdatedAt).getTime() - new Date(a.UpdatedAt).getTime())
)

const allReleases = computed(() => [...stableReleases.value, ...devReleases.value])
const latestStable = computed(() => stableReleases.value[0] ?? null)

onMounted(async () => {
  try {
    const all = await softwareApi.index()
    const item = all.find(s => s.software.name === route.params.name)
    if (item) {
      software.value = item.software
      releases.value = item.releases
    } else {
      error.value = 'Game not found'
    }
  } catch (e: any) {
    error.value = `Failed to load: ${e.message}`
  }
})
</script>

<style scoped>
.status-badge {
  @apply rounded text-[10px] font-black uppercase tracking-wider;
}
.status-released { @apply bg-green-500 text-white shadow-lg shadow-green-500/20; }
.status-demo { @apply bg-blue-500 text-white shadow-lg shadow-blue-500/20; }
.status-development { @apply bg-yellow-500 text-white shadow-lg shadow-yellow-500/20; }
.status-archived { @apply bg-gray-500 text-white shadow-lg shadow-gray-500/20; }
</style>
