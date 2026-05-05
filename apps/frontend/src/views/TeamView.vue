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
      <div v-for="member in teamMembers" :key="member.name" class="team-card">
        <div class="team-card-image-container">
          <img :src="member.image" :alt="member.name" class="team-card-image" />
        </div>
        <div class="team-card-content">
          <h2 class="team-card-title">{{ member.name }}</h2>
          <p class="team-card-desc">{{ member.description }}</p>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import mrZeroImg from '../assets/team/mr.zero.png'
import mrOneImg from '../assets/team/mr.one.png'
import mrTwoImg from '../assets/team/mr.two.png'
import mrThreeImg from '../assets/team/mr.three.png'

const teamMembers = [
  { name: 'Mr. Zero: Tasi', description: 'His dream is to become an open-source knight.', image: mrZeroImg },
  { name: 'Mr. One: Ballz', description: 'The cheese is half-eaten. Something here went very wrong.', image: mrOneImg },
  { name: 'Mr. Two: Z', description: "The egg was first or the chicken? It doesn't matter, we eat both.", image: mrTwoImg },
  { name: 'Mr. Three: gBird', description: 'Life is beautiful because it contains the possibility of becoming human.', image: mrThreeImg },
]
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
