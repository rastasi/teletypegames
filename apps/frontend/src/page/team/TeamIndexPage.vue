<template>
  <header class="hero-section-gradient from-purple-600 to-blue-700 py-20 text-white">
    <div class="hero-container text-center">
      <h1 class="hero-title mb-6">Our Team</h1>
      <p class="hero-subtitle mb-10 text-purple-100">
        Meet the brilliant minds behind Teletype Games.
      </p>
    </div>
  </header>

  <main class="main-container py-16 px-4">
    <div class="team-grid">
      <div v-for="member in members" :key="member.nick" class="team-card">
        <div class="team-card-image-container">
          <img :src="avatarUrl(member.avatar_filename)" :alt="member.nick" class="team-card-image" />
        </div>
        <div class="team-card-content">
          <h2 class="team-card-title">{{ member.nick }}: {{ member.real_nick }}</h2>
          <p class="team-card-desc">{{ member.motto }}</p>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import memberApi from '../../api/member.api'
import type { Member } from '../../lib/interfaces/member.interface'

const members = ref<Member[]>([])

function avatarUrl(filename: string): string {
  return new URL(`../../assets/team/${filename}`, import.meta.url).href
}

onMounted(async () => {
  try {
    members.value = await memberApi.index()
  } catch (e) {
    console.error('Failed to load team members:', e)
  }
})
</script>

<style scoped>
.team-grid {
  @apply grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8 max-w-7xl mx-auto;
}
.team-card {
  @apply bg-white rounded-2xl shadow-xl overflow-hidden border border-gray-100 transition-all hover:scale-105;
}
.team-card-image-container {
  @apply aspect-square overflow-hidden bg-gray-200;
}
.team-card-image {
  @apply w-full h-full object-cover;
}
.team-card-content {
  @apply p-6;
}
.team-card-title {
  @apply text-xl font-black text-gray-900 mb-2;
}
.team-card-desc {
  @apply text-gray-600 leading-relaxed;
}
</style>
