<template>
  <header class="hero-section-gradient from-teal-600 to-green-600 py-20">
    <div class="hero-container max-w-4xl">
      <h1 class="hero-title">Codebase</h1>
      <p class="hero-subtitle text-teal-50 mb-10">
        Explore our collection of open-source projects, study our source code, and contribute to our independent game development tools.
      </p>
      <a href="https://git.teletype.hu" target="_blank" class="btn-hero-teal">
        Open Gitea
      </a>
    </div>
  </header>

  <main class="main-container">
    <div v-if="error" class="error-alert" role="alert">
      <strong class="font-bold">Error!</strong>
      <span class="block sm:inline"> {{ error }}</span>
    </div>

    <div v-if="!error && publicRepos.length > 0" class="card-grid mb-16">
      <div v-for="repo in publicRepos" :key="repo.name" class="card-base">
        <h3 class="repo-card-title">
          <a :href="repo.html_url" target="_blank" class="repo-link">
            {{ repo.owner.login }}/{{ repo.name }}
          </a>
        </h3>
        <p v-if="repo.description" class="repo-card-desc">{{ repo.description }}</p>
        <p v-if="repo.language" class="repo-card-meta">Language: {{ repo.language }}</p>
        <p class="repo-card-meta">Updated: {{ formatDateTime(repo.updated_at) }}</p>
      </div>
    </div>

    <h2 class="commits-section-title">Recent Commits</h2>

    <p v-if="!error && recentCommits.length === 0" class="empty-commits-msg">
      No recent commits found or accessible from public repositories.
    </p>

    <div class="commits-list">
      <div v-for="commit in recentCommits" :key="commit.sha" class="commit-item">
        <div class="commit-icon">🚀</div>
        <div class="commit-content">
          <p class="commit-message">
            <a :href="commit.repo.html_url" target="_blank" class="commit-repo-link">
              <span class="font-bold">{{ commit.repo.owner }}</span> / <span class="font-bold">{{ commit.repo.name }}</span>
            </a>
            : <span class="text-gray-700">{{ commit.message.split('\n')[0] }}</span>
          </p>
          <p class="commit-meta">
            Authored by <span class="font-bold">{{ commit.author.name }}</span> on {{ formatDateTime(commit.date) }}
            <a :href="commit.url" target="_blank" class="commit-sha-link">
              <span class="font-mono text-xs">{{ commit.sha.substring(0, 7) }}</span>
            </a>
          </p>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { formatDateTime } from '../../lib/dateFormat'
import gitApi from '../../api/git.api'
import type { GiteaRepo, Commit } from '../../lib/interfaces/git.interface'

const publicRepos = ref<GiteaRepo[]>([])
const recentCommits = ref<Commit[]>([])
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    publicRepos.value = await gitApi.repos()

    const commitPromises = publicRepos.value.map(repo =>
      gitApi.commits(repo.owner.login, repo.name, repo.html_url)
    )

    const all = (await Promise.all(commitPromises)).flat()
    all.sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
    recentCommits.value = all.slice(0, 20)
  } catch (e: any) {
    error.value = e.message
  }
})
</script>

<style scoped>
.error-alert {
  @apply bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-8;
}
.repo-card-title {
  @apply text-xl font-semibold mb-2;
}
.repo-link {
  @apply text-blue-600 hover:text-blue-800;
}
.repo-card-desc {
  @apply text-gray-700 text-sm;
}
.repo-card-meta {
  @apply text-gray-500 text-xs mt-1;
}
.commits-section-title {
  @apply text-3xl font-bold mb-8 text-center text-gray-800;
}
.empty-commits-msg {
  @apply text-center text-gray-600 text-lg;
}
.commits-list {
  @apply space-y-4 mb-12;
}
.commit-item {
  @apply bg-white shadow-lg rounded-lg p-4 flex items-center space-x-4 border border-gray-100;
}
.commit-icon {
  @apply flex-shrink-0 text-3xl;
}
.commit-content {
  @apply flex-grow;
}
.commit-message {
  @apply text-gray-800 font-semibold;
}
.commit-repo-link {
  @apply text-blue-600 hover:text-blue-800;
}
.commit-meta {
  @apply text-gray-600 text-sm mt-1;
}
.commit-sha-link {
  @apply ml-2 text-blue-500 hover:text-blue-700;
}
</style>
