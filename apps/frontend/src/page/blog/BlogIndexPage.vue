<template>
  <div class="blog-container">
    <header class="blog-header">
      <div class="blog-header-decor">
        <div class="hero-decor-blob -top-24 -left-24 h-96 w-96 bg-purple-600"></div>
        <div class="hero-decor-blob -bottom-24 -right-24 h-96 w-96 bg-blue-600"></div>
      </div>

      <div class="hero-container">
        <div class="hero-badge">Community Stories</div>
        <h1 class="hero-title">
          Developer <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-400">Blog</span>
        </h1>
        <p class="hero-subtitle mb-10">
          News, updates, and behind-the-scenes stories from the Teletype Games community.
        </p>
      </div>
    </header>

    <main class="blog-main">
      <div v-if="error" class="error-banner" role="alert">
        <span class="text-2xl">⚠️</span>
        <div>
          <h3 class="error-banner-title">Failed to connect to Wiki</h3>
          <p class="error-banner-desc">{{ error }}</p>
        </div>
      </div>

      <div v-else-if="!loading && blogPages.length === 0" class="empty-state">
        <div class="empty-state-icon">📭</div>
        <h2 class="empty-state-title">No blog posts found</h2>
        <p class="empty-state-desc">We haven't published any blog posts yet. Check back soon!</p>
      </div>

      <div v-else class="blog-posts-list">
        <article v-for="page in blogPages" :key="page.id" class="blog-post-card">
          <div class="blog-post-content">
            <div class="blog-post-meta">
              <div class="flex items-center gap-4">
                <div class="blog-post-timestamp">
                  <span v-if="isNew(page.createdAt)" class="new-post-badge">
                    <span class="new-post-dot"></span>
                    New Post
                  </span>
                  <span>{{ formatDateTime(page.createdAt) }}</span>
                </div>
              </div>
            </div>

            <h2 class="blog-post-title">
              <RouterLink :to="getPermalink(page.path)" class="hover:text-indigo-600 transition-colors">
                {{ page.title }}
              </RouterLink>
            </h2>

            <div class="blog-post-preview">
              <p v-if="page.description" class="blog-post-description">{{ page.description }}</p>
              <p v-if="page.content" class="blog-post-body">{{ getCleanPreview(page.content) }}</p>
            </div>

            <div class="mt-6">
              <RouterLink :to="getPermalink(page.path)" class="read-more-btn">
                Read more
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </RouterLink>
            </div>
          </div>
        </article>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { formatDateTime } from '../../lib/dateFormat'
import wikiApi from '../../api/wiki.api'
import type { WikiPageWithContent } from '../../lib/interfaces/wiki.interface'

const blogPages = ref<WikiPageWithContent[]>([])
const error = ref<string | null>(null)
const loading = ref(true)

function isNew(createdAt: string): boolean {
  return (new Date().getTime() - new Date(createdAt).getTime()) < 7 * 24 * 60 * 60 * 1000
}

function getPermalink(path: string): string {
  const slug = path.startsWith('blog/') ? path.replace('blog/', '') : path
  return `/blog/${slug}`
}

function getCleanPreview(content: string): string {
  if (!content) return ''
  return content.replace(/[#*`_\[\]()]/g, '').trim().slice(0, 300) + '...'
}

onMounted(async () => {
  try {
    blogPages.value = await wikiApi.listBlogPages()
  } catch (e: any) {
    error.value = `Failed to fetch wiki data: ${e.message}`
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.blog-container {
  @apply bg-gray-50 min-h-screen pb-20;
}
.blog-header {
  @apply relative overflow-hidden py-24 text-white text-center bg-indigo-900;
}
.blog-header-decor {
  @apply absolute inset-0 opacity-30;
}
.blog-main {
  @apply max-w-5xl mx-auto px-4 md:px-8 -mt-12 relative z-10;
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
.blog-posts-list {
  @apply flex flex-col gap-12;
}
.blog-post-card {
  @apply bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden;
}
.blog-post-content {
  @apply p-6 md:p-10;
}
.blog-post-meta {
  @apply flex items-center justify-between mb-4;
}
.blog-post-timestamp {
  @apply flex items-center gap-2 text-xs font-semibold uppercase tracking-wider text-gray-400;
}
.new-post-badge {
  @apply flex items-center gap-1 text-green-600 bg-green-50 px-2 py-0.5 rounded-full border border-green-100;
}
.new-post-dot {
  @apply w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse;
}
.blog-post-title {
  @apply text-2xl md:text-3xl font-black text-gray-900 mb-4 leading-tight;
}
.blog-post-preview {
  @apply relative overflow-hidden;
  max-height: 8rem;
}
.blog-post-preview::after {
  content: "";
  @apply absolute bottom-0 left-0 w-full h-16 bg-gradient-to-b from-transparent to-white pointer-events-none;
}
.blog-post-description {
  @apply text-lg text-gray-800 font-bold mb-2 leading-relaxed;
}
.blog-post-body {
  @apply text-base text-gray-500 font-medium leading-relaxed;
}
.read-more-btn {
  @apply inline-flex items-center text-indigo-600 font-bold hover:text-indigo-800 transition-colors;
}
</style>
