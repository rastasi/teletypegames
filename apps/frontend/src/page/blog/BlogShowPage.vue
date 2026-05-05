<template>
  <div class="blog-container">
    <header class="blog-header">
      <div class="blog-header-decor">
        <div class="hero-decor-blob -top-24 -left-24 h-96 w-96 bg-purple-600"></div>
        <div class="hero-decor-blob -bottom-24 -right-24 h-96 w-96 bg-blue-600"></div>
      </div>

      <div class="post-hero-container">
        <RouterLink to="/blog" class="back-link">← Back to Blog</RouterLink>
        <template v-if="pageContent">
          <h1 class="hero-title mt-4">{{ pageContent.title }}</h1>
          <p class="hero-subtitle mb-4">{{ pageContent.description }}</p>
          <div class="blog-post-timestamp text-indigo-200">{{ formatDateTime(pageContent.createdAt) }}</div>
        </template>
      </div>
    </header>

    <main class="blog-main">
      <div v-if="error" class="error-banner" role="alert">
        <span class="text-2xl">⚠️</span>
        <div>
          <h3 class="error-banner-title">Error</h3>
          <p class="error-banner-desc">{{ error }}</p>
        </div>
      </div>

      <article v-else-if="pageContent" class="blog-post-card">
        <div class="blog-post-content">
          <div v-if="pageContent.render" class="wiki-content max-w-none text-gray-700 leading-relaxed" v-html="pageContent.render" />
          <div v-else class="blog-post-fallback">No content available for this post.</div>
        </div>
      </article>

      <div v-else-if="loading" class="blog-post-card">
        <div class="blog-post-content text-center text-gray-400 py-12">Loading...</div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { formatDateTime } from '../../lib/dateFormat'
import wikiApi from '../../api/wiki.api'
import type { WikiPageContent } from '../../lib/interfaces/wiki.interface'

const route = useRoute()
const pageContent = ref<WikiPageContent | null>(null)
const error = ref<string | null>(null)
const loading = ref(true)

onMounted(async () => {
  const slug = route.params.slug as string
  try {
    const result = await wikiApi.getBlogPage(slug)
    if (result) {
      pageContent.value = result
    } else {
      error.value = 'Post not found'
    }
  } catch (e: any) {
    error.value = e.message
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
.post-hero-container {
  @apply relative z-10 max-w-4xl mx-auto px-4;
}
.back-link {
  @apply inline-block text-indigo-300 hover:text-white transition-colors font-medium;
}
.hero-title {
  @apply text-4xl md:text-5xl font-black mb-4 leading-tight;
}
.hero-subtitle {
  @apply text-xl text-indigo-100 font-medium;
}
.blog-post-timestamp {
  @apply text-sm font-semibold uppercase tracking-wider;
}
.blog-main {
  @apply max-w-4xl mx-auto px-4 md:px-8 -mt-12 relative z-10;
}
.blog-post-card {
  @apply bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden;
}
.blog-post-content {
  @apply p-6 md:p-12;
}
.error-banner {
  @apply max-w-7xl mx-auto bg-red-50 border-l-4 border-red-500 p-6 rounded-r-xl shadow-lg mb-12 flex items-start gap-4;
}
.error-banner-title { @apply text-red-800 font-bold text-lg; }
.error-banner-desc { @apply text-red-700 mt-1; }
</style>
