<template>
  <div class="home-container">
    <!-- Header Section -->
    <header class="hero-section-slate">
      <div class="home-header-decor">
        <div class="hero-decor-blob -top-24 -left-24 h-96 w-96 bg-indigo-600"></div>
        <div class="hero-decor-blob -bottom-24 -right-24 h-96 w-96 bg-purple-600"></div>
      </div>

      <div class="hero-container">
        <div class="hero-badge">Welcome to</div>
        <h1 class="hero-title">
          Teletype <span class="text-transparent bg-clip-text bg-gradient-to-r from-violet-400 to-fuchsia-400">Games</span>
        </h1>
        <p class="hero-subtitle">
          Teletype Games is an independent game development community built on creative freedom, equality, and an open-source mindset. Our goal is to create experimental and full-fledged games with short development cycles.
        </p>
      </div>
    </header>

    <main class="main-container max-w-6xl">
      <!-- Countdown Section -->
      <section v-if="nextEvent" class="mb-12">
        <div class="countdown-card">
          <h2 class="countdown-event-name">{{ nextEvent.name }}</h2>
          <div class="countdown-timer">{{ countdownText }}</div>
          <p class="countdown-date">Starts on: <span class="font-semibold">{{ nextEvent.dateText }}</span></p>

          <div v-if="followingEvents.length > 0" class="upcoming-events-section">
            <h3 class="upcoming-events-title">Upcoming Events</h3>
            <div class="space-y-4">
              <div v-for="event in followingEvents" :key="event.name" class="upcoming-event-item">
                <span class="font-bold text-gray-800">{{ event.name }}</span>
                <span class="text-sm text-gray-500 font-mono">{{ event.dateText }}</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Featured Software Section -->
      <section v-if="highlightedSoftware" class="mb-16">
        <div class="featured-card group">
          <div class="featured-content">
            <div class="flex items-center gap-2 mb-4">
              <span class="featured-badge">Featured Game</span>
              <span :class="`status-badge status-${highlightedSoftware.software.status}`">{{ highlightedSoftware.software.status }}</span>
              <span v-if="highlightedStableRelease" class="featured-version-badge">{{ highlightedStableRelease.version }}</span>
            </div>
            <h2 class="featured-title">{{ highlightedSoftware.software.title }}</h2>
            <p class="featured-desc">{{ highlightedSoftware.software.desc }}</p>

            <div class="flex flex-wrap gap-4 mt-8">
              <a
                v-if="highlightedStableRelease?.htmlFolderPath"
                :href="BASE + highlightedStableRelease.htmlFolderPath"
                target="_blank"
                class="featured-btn-play"
              >▶ Play in Browser</a>
              <RouterLink :to="`/catalog/${highlightedSoftware.software.name}`" class="featured-btn-secondary">
                View Project Details
              </RouterLink>
            </div>
          </div>
        </div>
      </section>

      <!-- Latest YouTube Video -->
      <section class="mb-16">
        <div class="yt-featured-wrapper group">
          <div class="yt-glow"></div>

          <div class="yt-header-area">
            <div class="flex items-center gap-3">
              <div class="yt-status-blob">
                <div class="yt-status-dot"></div>
              </div>
              <span class="yt-label-text">Latest from YouTube</span>
            </div>
            <a href="https://www.youtube.com/@teletypegames" target="_blank" class="yt-channel-link">
              Visit Channel ↗
            </a>
          </div>

          <div class="yt-content-grid">
            <div class="yt-player-container">
              <div class="player-wrap relative w-full aspect-video">
                <div v-if="ytState === 'loading'" class="state-overlay absolute inset-0 flex items-center justify-center bg-slate-950">
                  <div class="w-10 h-10 border-4 border-white/5 border-t-indigo-500 rounded-full animate-spin"></div>
                </div>
                <div v-else-if="ytState === 'error'" class="state-overlay absolute inset-0 flex items-center justify-center bg-slate-950 flex-col gap-2">
                  <div class="text-xl">⚠️</div>
                  <div class="text-[11px] text-slate-500 text-center px-4 font-bold uppercase tracking-tighter">{{ ytError }}</div>
                </div>
                <div v-else-if="ytState === 'no-config'" class="state-overlay absolute inset-0 flex items-center justify-center bg-slate-950">
                  <div class="text-slate-500 text-xs">Missing API Configuration</div>
                </div>
                <iframe
                  v-else-if="ytVideoId"
                  :src="`https://www.youtube.com/embed/${ytVideoId}?autoplay=0&rel=0&modestbranding=1`"
                  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                  allowfullscreen
                  class="w-full h-full border-none block"
                ></iframe>
              </div>
            </div>

            <div v-if="ytVideoId" class="yt-info-sidebar">
              <div class="mb-auto">
                <span class="yt-new-tag">New Release</span>
                <h3 class="yt-video-title">{{ ytTitle }}</h3>
                <div class="yt-stats-row">
                  <span v-if="ytPublishDate">📅 {{ ytPublishDate }}</span>
                  <span v-if="ytViewCount">👁 {{ ytViewCount }}</span>
                </div>
              </div>

              <div class="mt-6">
                <a :href="`https://www.youtube.com/watch?v=${ytVideoId}`" target="_blank" class="yt-play-button">
                  <span class="text-xl">▶</span> Watch on YouTube
                </a>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Navigation Grid -->
      <section class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <div class="nav-card group card-interactive">
          <div class="nav-card-icon bg-purple-50 group-hover:bg-purple-100">
            <span class="text-3xl">🎮</span>
          </div>
          <h2 class="nav-card-title">Catalog</h2>
          <p class="nav-card-desc">Here you can find our games. Discover our creations for TIC-80 and other platforms, many of which you can try directly in your browser.</p>
          <RouterLink to="/catalog" class="btn-primary bg-purple-600 hover:bg-purple-700 text-white hover:shadow-purple-200">
            Browse Games
          </RouterLink>
        </div>

        <div class="nav-card group card-interactive">
          <div class="nav-card-icon bg-green-50 group-hover:bg-green-100">
            <span class="text-3xl">📚</span>
          </div>
          <h2 class="nav-card-title">Wiki</h2>
          <p class="nav-card-desc">We gather our public interest development knowledge here. Read through technical descriptions and development logs.</p>
          <a href="https://wiki.teletype.hu" target="_blank" rel="noopener noreferrer" class="btn-primary bg-green-600 hover:bg-green-700 text-white hover:shadow-green-200">
            Knowledge Base
          </a>
        </div>

        <div class="nav-card group card-interactive">
          <div class="nav-card-icon bg-slate-50 group-hover:bg-slate-100">
            <span class="text-3xl">💻</span>
          </div>
          <h2 class="nav-card-title">Git Repos</h2>
          <p class="nav-card-desc">All our projects are open source. Browse our code, study our solutions, or even participate in the development.</p>
          <a href="https://git.teletype.hu" target="_blank" rel="noopener noreferrer" class="btn-primary bg-slate-800 hover:bg-slate-900 text-white hover:shadow-slate-200">
            Gitea
          </a>
        </div>

        <div class="nav-card group card-interactive">
          <div class="nav-card-icon bg-red-50 group-hover:bg-red-100">
            <span class="text-3xl">📺</span>
          </div>
          <h2 class="nav-card-title">YouTube</h2>
          <p class="nav-card-desc">Watch our development vlogs, gameplay videos, and tutorials on our official YouTube channel.</p>
          <a href="https://www.youtube.com/@teletypegames" target="_blank" rel="noopener noreferrer" class="btn-primary bg-red-600 hover:bg-red-700 text-white hover:shadow-red-200">
            Watch on YouTube
          </a>
        </div>
      </section>

      <!-- AI Content Notice -->
      <div class="mb-12 mt-12 bg-white rounded-2xl border border-indigo-100 p-6 shadow-sm flex items-center gap-4">
        <div class="w-12 h-12 rounded-full bg-indigo-50 flex items-center justify-center text-2xl flex-shrink-0">🤖</div>
        <div class="flex-grow text-slate-600 text-sm md:text-base italic">
          <strong>Note:</strong> Our HowTo-s and blog posts are written by AI with human supervision.
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'
import { formatDateTime } from '../utils/dateFormat'
import { GAMES_BASE } from '../config'

const BASE = GAMES_BASE
const YOUTUBE_API_KEY = import.meta.env.YOUTUBE_API_KEY
const YOUTUBE_CHANNEL_ID = import.meta.env.YOUTUBE_CHANNEL_ID

interface Event { name: string; date: string }

const nextEvent = ref<(Event & { dateISO: string; dateText: string }) | null>(null)
const followingEvents = ref<(Event & { dateText: string })[]>([])
const highlightedSoftware = ref<any>(null)
const highlightedStableRelease = ref<any>(null)
const countdownText = ref('-- : -- : -- : --')

const ytState = ref<'loading' | 'loaded' | 'error' | 'no-config'>('loading')
const ytVideoId = ref('')
const ytTitle = ref('')
const ytPublishDate = ref('')
const ytViewCount = ref('')
const ytError = ref('')

let countdownInterval: ReturnType<typeof setInterval> | null = null

function startCountdown(targetISO: string) {
  const targetDate = new Date(targetISO).getTime()

  function update() {
    const distance = targetDate - Date.now()
    if (distance < 0) {
      countdownText.value = 'LIVE NOW!'
      if (countdownInterval) clearInterval(countdownInterval)
      return
    }
    const d = String(Math.floor(distance / 86400000)).padStart(2, '0')
    const h = String(Math.floor((distance % 86400000) / 3600000)).padStart(2, '0')
    const m = String(Math.floor((distance % 3600000) / 60000)).padStart(2, '0')
    const s = String(Math.floor((distance % 60000) / 1000)).padStart(2, '0')
    countdownText.value = `${d}d ${h}h ${m}m ${s}s`
  }

  update()
  countdownInterval = setInterval(update, 1000)
}

function formatYtDate(iso: string): string {
  const d = new Date(iso)
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')} ${String(d.getHours()).padStart(2, '0')}:${String(d.getMinutes()).padStart(2, '0')}`
}

function formatViews(n: number): string {
  if (n >= 1_000_000) return (n / 1_000_000).toFixed(1) + 'M views'
  if (n >= 1_000) return Math.round(n / 1_000) + 'K views'
  return n + ' views'
}

async function loadLatestVideo() {
  if (!YOUTUBE_API_KEY || !YOUTUBE_CHANNEL_ID) {
    ytState.value = 'no-config'
    return
  }

  try {
    const searchRes = await fetch(`https://www.googleapis.com/youtube/v3/search?key=${YOUTUBE_API_KEY}&channelId=${YOUTUBE_CHANNEL_ID}&part=snippet&order=date&maxResults=1&type=video`)
    if (!searchRes.ok) throw new Error(`API error: ${searchRes.status}`)
    const searchData = await searchRes.json()

    if (!searchData.items?.length) {
      ytError.value = 'No videos found on this channel.'
      ytState.value = 'error'
      return
    }

    const item = searchData.items[0]
    const videoId = item.id.videoId
    const title = item.snippet.title
    const pubDate = item.snippet.publishedAt

    const statsRes = await fetch(`https://www.googleapis.com/youtube/v3/videos?key=${YOUTUBE_API_KEY}&id=${videoId}&part=statistics`)
    const statsData = await statsRes.json()
    const views = statsData.items?.[0]?.statistics?.viewCount ?? null

    ytVideoId.value = videoId
    ytTitle.value = title
    ytPublishDate.value = formatYtDate(pubDate)
    ytViewCount.value = views ? formatViews(Number(views)) : ''
    ytState.value = 'loaded'
  } catch (err: any) {
    ytError.value = 'Error loading video: ' + err.message
    ytState.value = 'error'
  }
}

onMounted(async () => {
  // Load events
  try {
    const res = await fetch(BASE + '/api/events')
    if (res.ok) {
      const events: Event[] = await res.json()
      if (events.length > 0) {
        const first = events[0]
        const firstDate = new Date(first.date)
        nextEvent.value = {
          name: first.name,
          date: first.date,
          dateISO: firstDate.toISOString(),
          dateText: formatDateTime(firstDate),
        }
        followingEvents.value = events.slice(1).map(e => ({
          name: e.name,
          date: e.date,
          dateText: formatDateTime(new Date(e.date)),
        }))
        startCountdown(firstDate.toISOString())
      }
    }
  } catch (e) {
    console.error('Failed to load events:', e)
  }

  // Load highlighted software
  try {
    const res = await fetch(BASE + '/api/software/highlighted')
    if (res.ok) {
      const data = await res.json()
      highlightedSoftware.value = data
      if (data?.releases) {
        const stable = data.releases
          .filter((r: any) => !r.version.startsWith('dev-'))
          .sort((a: any, b: any) => new Date(b.UpdatedAt).getTime() - new Date(a.UpdatedAt).getTime())
        highlightedStableRelease.value = stable[0] ?? null
      }
    }
  } catch (e) {
    console.error('Failed to fetch highlighted software:', e)
  }

  // Load YouTube
  await loadLatestVideo()
})

onUnmounted(() => {
  if (countdownInterval) clearInterval(countdownInterval)
})
</script>

<style scoped>
.home-container {
  @apply bg-gray-50 min-h-screen pb-20;
}
.home-header-decor {
  @apply absolute inset-0 opacity-30;
}

/* Featured Card */
.featured-card {
  @apply relative overflow-hidden bg-gradient-to-br from-indigo-900 to-slate-900 rounded-3xl shadow-2xl flex flex-col md:flex-row min-h-[400px];
}
.featured-content {
  @apply p-8 md:p-12 flex flex-col justify-center flex-1 z-10;
}
.featured-badge {
  @apply px-3 py-1 bg-indigo-500/20 text-indigo-300 border border-indigo-400/30 rounded-full text-xs font-bold uppercase tracking-wider;
}
.featured-title {
  @apply text-4xl md:text-5xl font-black text-white mb-6;
}
.featured-desc {
  @apply text-indigo-100/80 text-lg leading-relaxed max-w-xl;
}
.featured-btn-play {
  @apply px-8 py-4 bg-white text-indigo-900 font-bold rounded-xl transition-all hover:scale-105 hover:shadow-xl hover:shadow-indigo-500/20 active:scale-95;
}
.featured-btn-secondary {
  @apply px-8 py-4 bg-indigo-500/10 text-white border border-indigo-400/30 font-bold rounded-xl transition-all hover:bg-indigo-500/20;
}
.featured-version-badge {
  @apply px-2 py-0.5 rounded text-[10px] font-bold bg-white/10 text-white/70 border border-white/20;
}

/* Status Badges */
.status-badge {
  @apply px-2 py-0.5 rounded text-[10px] font-bold uppercase;
}
.status-released { @apply bg-green-500/20 text-green-300 border border-green-500/30; }
.status-demo { @apply bg-blue-500/20 text-blue-300 border border-blue-500/30; }
.status-development { @apply bg-yellow-500/20 text-yellow-300 border border-yellow-500/30; }
.status-archived { @apply bg-gray-500/20 text-gray-300 border border-gray-500/30; }

/* Countdown Card */
.countdown-card {
  @apply bg-white rounded-2xl shadow-xl overflow-hidden border border-gray-100 p-8 md:p-12 text-center transition-all duration-300;
}
.countdown-event-name {
  @apply text-2xl font-bold text-gray-400 uppercase tracking-widest mb-2;
}
.countdown-timer {
  @apply text-5xl md:text-7xl font-mono font-black text-transparent bg-clip-text bg-gradient-to-r from-violet-600 to-fuchsia-600 my-6;
}
.countdown-date {
  @apply text-gray-500 text-lg;
}
.upcoming-events-section {
  @apply mt-12 pt-12 border-t border-gray-100 text-left max-w-2xl mx-auto;
}
.upcoming-events-title {
  @apply text-xl font-bold text-gray-900 mb-6;
}
.upcoming-event-item {
  @apply flex items-center justify-between p-4 bg-gray-50 rounded-xl border border-gray-100;
}

/* Nav Cards */
.nav-card {
  @apply flex flex-col h-full;
}
.nav-card-icon {
  @apply w-14 h-14 rounded-xl flex items-center justify-center mb-6 transition-colors;
}
.nav-card-title {
  @apply text-2xl font-bold text-gray-900 mb-3;
}
.nav-card-desc {
  @apply text-gray-600 mb-6 flex-grow leading-relaxed;
}

/* YouTube Featured Wrapper */
.yt-featured-wrapper {
  @apply relative overflow-hidden bg-gradient-to-br from-slate-900 via-slate-900 to-indigo-950 rounded-3xl shadow-2xl border border-slate-800 p-6 md:p-8;
}
.yt-glow {
  @apply absolute -top-24 -right-24 w-96 h-96 bg-indigo-500/10 rounded-full blur-[100px] pointer-events-none;
}
.yt-header-area {
  @apply flex items-center justify-between mb-6 relative z-10;
}
.yt-status-blob {
  @apply w-3 h-3 bg-red-500/20 rounded-full flex items-center justify-center;
}
.yt-status-dot {
  @apply w-1.5 h-1.5 bg-red-500 rounded-full animate-pulse;
}
.yt-label-text {
  @apply text-xs font-bold uppercase tracking-widest text-slate-400;
}
.yt-channel-link {
  @apply text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition-colors;
}
.yt-content-grid {
  @apply flex flex-col lg:flex-row gap-8 relative z-10;
}
.yt-player-container {
  @apply flex-[1.6] overflow-hidden rounded-2xl border border-white/5 shadow-inner bg-black;
}
.yt-info-sidebar {
  @apply flex-1 flex flex-col justify-center;
}
.yt-new-tag {
  @apply inline-block px-2 py-1 bg-red-500/10 text-red-400 border border-red-500/20 rounded-md text-[10px] font-black uppercase tracking-tighter mb-4;
}
.yt-video-title {
  @apply text-2xl font-black text-white mb-4 leading-tight group-hover:text-indigo-200 transition-colors;
}
.yt-stats-row {
  @apply flex gap-4 text-sm text-slate-400 font-medium;
}
.yt-play-button {
  @apply flex items-center justify-center gap-2 w-full py-4 bg-white text-slate-900 font-black rounded-xl transition-all hover:scale-[1.02] hover:bg-indigo-50 active:scale-95 shadow-xl;
}
</style>
