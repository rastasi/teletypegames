<template>
  <div class="howtos-container">
    <header class="hero-section-slate">
      <div class="howtos-header-decor">
        <div class="hero-decor-blob -top-24 -left-24 h-96 w-96 bg-indigo-600"></div>
        <div class="hero-decor-blob -bottom-24 -right-24 h-96 w-96 bg-purple-600"></div>
      </div>

      <div class="hero-container">
        <div class="hero-badge">Knowledge Base</div>
        <h1 class="hero-title">
          Tech <span class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-400 to-purple-400">HowTo Center</span>
        </h1>
        <p class="hero-subtitle mb-10">
          The latest documentation, technical guides, and development logs from our knowledge base.
        </p>
        <a :href="WIKIJS_BASE_URL" target="_blank" class="btn-hero-indigo">
          Knowledge Base
        </a>
      </div>
    </header>

    <main class="main-container">
      <div v-if="error" class="error-banner" role="alert">
        <span class="text-2xl">⚠️</span>
        <div>
          <h3 class="error-banner-title">Failed to connect to Wiki</h3>
          <p class="error-banner-desc">{{ error }}</p>
        </div>
      </div>

      <div v-else-if="!loading && recentPages.length === 0" class="empty-state">
        <div class="empty-state-icon">📭</div>
        <h2 class="empty-state-title">No pages found</h2>
        <p class="empty-state-desc">It seems like there aren't any public pages available on the wiki at the moment.</p>
      </div>

      <div v-else class="card-grid">
        <a
          v-for="page in recentPages"
          :key="page.id"
          :href="`${WIKIJS_BASE_URL}/${page.locale}/${page.path}`"
          target="_blank"
          class="group card-interactive flex flex-col"
        >
          <div class="howto-card-meta">
            <div class="howto-card-timestamp">
              <span v-if="isNew(page.createdAt)" class="new-badge">
                <span class="new-dot"></span>
                New
              </span>
              <span>{{ formatDateTime(page.updatedAt) }}</span>
            </div>
          </div>

          <h2 class="howto-card-title">{{ page.title }}</h2>

          <p v-if="page.description" class="howto-card-desc">{{ page.description }}</p>
          <div v-else class="howto-card-no-desc">No description provided.</div>

          <div class="howto-card-footer">
            <span class="howto-card-path">{{ page.path }}</span>
            <span class="howto-card-read-link">Read <span class="text-lg">→</span></span>
          </div>
        </a>
      </div>

      <div v-if="!error && recentPages.length > 0" class="howtos-footer">
        <a :href="WIKIJS_BASE_URL" target="_blank" class="browse-wiki-btn">
          Browse Wiki Home
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd" />
          </svg>
        </a>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { formatDateTime } from '../utils/dateFormat'
import { WIKI_BASE } from '../config'

const WIKIJS_BASE_URL = WIKI_BASE
const WIKIJS_TOKEN = import.meta.env.WEBAPP_WIKIJS_TOKEN

interface WikiPage {
  id: number
  path: string
  title: string
  description: string
  updatedAt: string
  createdAt: string
  locale: string
}

const recentPages = ref<WikiPage[]>([])
const error = ref<string | null>(null)
const loading = ref(true)

function isNew(createdAt: string): boolean {
  return (new Date().getTime() - new Date(createdAt).getTime()) < 7 * 24 * 60 * 60 * 1000
}

onMounted(async () => {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  }
  if (WIKIJS_TOKEN) {
    headers['Authorization'] = `Bearer ${WIKIJS_TOKEN}`
  }

  const GRAPHQL_QUERY = `
    {
      pages {
        list(orderBy: UPDATED, orderByDirection: DESC, tags: ["howto"]) {
          id path title description updatedAt createdAt locale
        }
      }
    }
  `

  try {
    const response = await fetch(`${WIKIJS_BASE_URL}/graphql`, {
      method: 'POST',
      headers,
      body: JSON.stringify({ query: GRAPHQL_QUERY }),
    })

    if (!response.ok) throw new Error(`WikiJS responded with status ${response.status}`)

    const json = await response.json()
    if (json.errors) throw new Error(`GraphQL error: ${json.errors.map((e: any) => e.message).join(', ')}`)

    const pages: any[] = json?.data?.pages?.list ?? []
    recentPages.value = pages.map((p: any) => ({
      id: p.id,
      path: p.path,
      title: p.title || p.path,
      description: p.description ?? '',
      updatedAt: p.updatedAt,
      createdAt: p.createdAt,
      locale: p.locale,
    })).slice(0, 30)
  } catch (e: any) {
    error.value = `Failed to fetch wiki data: ${e.message}`
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.howtos-container {
  @apply bg-gray-50 min-h-screen pb-20;
}
.howtos-header-decor {
  @apply absolute inset-0 opacity-30;
}
.banner-base {
  @apply max-w-7xl mx-auto border-l-4 p-6 rounded-r-xl shadow-lg mb-12 flex items-start gap-4;
}
.error-banner { @apply banner-base bg-red-50 border-red-500; }
.error-banner-title { @apply text-red-800 font-bold text-lg; }
.error-banner-desc { @apply text-red-700 mt-1; }
.empty-state {
  @apply max-w-7xl mx-auto bg-white rounded-2xl shadow-xl p-12 text-center border border-gray-100;
}
.empty-state-icon { @apply text-6xl mb-4; }
.empty-state-title { @apply text-2xl font-bold text-gray-900 mb-2; }
.empty-state-desc { @apply text-gray-500 max-w-md mx-auto; }
.howto-card-meta {
  @apply flex items-center justify-end mb-4;
}
.howto-card-timestamp {
  @apply flex items-center gap-2 text-xs font-semibold uppercase tracking-wider text-gray-400;
}
.new-badge {
  @apply flex items-center gap-1 text-green-600 bg-green-50 px-2 py-0.5 rounded-full border border-green-100;
}
.new-dot {
  @apply w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse;
}
.howto-card-title {
  @apply text-xl font-bold text-gray-900 group-hover:text-indigo-600 transition-colors mb-2 line-clamp-2 leading-tight;
}
.howto-card-desc {
  @apply text-gray-600 text-sm mb-4 line-clamp-3 leading-relaxed flex-grow;
}
.howto-card-no-desc {
  @apply flex-grow mb-4 italic text-gray-400 text-sm;
}
.howto-card-footer {
  @apply pt-4 border-t border-gray-100 mt-auto flex items-center justify-between;
}
.howto-card-path {
  @apply text-xs font-mono text-gray-400 truncate max-w-[180px];
}
.howto-card-read-link {
  @apply text-indigo-600 text-sm font-bold flex items-center gap-1 group-hover:translate-x-1 transition-transform;
}
.howtos-footer {
  @apply mt-16 text-center;
}
.browse-wiki-btn {
  @apply inline-flex items-center gap-2 bg-slate-900 text-white px-8 py-4 rounded-xl hover:bg-slate-800 transition-all font-bold shadow-lg hover:shadow-indigo-200/50;
}
</style>
