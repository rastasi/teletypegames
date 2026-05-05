<template>
  <header class="bg-gray-800 text-white p-4 flex justify-between items-center relative">
    <RouterLink to="/" class="text-xl font-bold">Teletype Games</RouterLink>

    <!-- Desktop Navigation -->
    <nav class="hidden md:flex space-x-4">
      <RouterLink to="/" class="p-2 hover:text-purple-300 transition-colors duration-200">Home</RouterLink>
      <RouterLink to="/catalog" class="p-2 hover:text-purple-300 transition-colors duration-200">Catalog</RouterLink>
      <RouterLink to="/blog" class="p-2 hover:text-purple-300 transition-colors duration-200">Blog</RouterLink>
      <RouterLink to="/howtos" class="p-2 hover:text-purple-300 transition-colors duration-200">How-tos</RouterLink>
      <RouterLink to="/code" class="p-2 hover:text-purple-300 transition-colors duration-200">Code</RouterLink>
      <RouterLink to="/team" class="p-2 hover:text-purple-300 transition-colors duration-200">Team</RouterLink>
      <RouterLink to="/contact" class="p-2 hover:text-purple-300 transition-colors duration-200">Contact us</RouterLink>
    </nav>

    <!-- Hamburger Button (Mobile only) -->
    <button @click="toggleMenu" class="md:hidden p-2 focus:outline-none focus:ring-2 focus:ring-purple-500 rounded-md">
      <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
      </svg>
    </button>

    <!-- Mobile Navigation -->
    <nav v-show="menuOpen" class="md:hidden absolute top-full left-0 w-full bg-gray-800 flex flex-col items-center py-4 space-y-2 z-50">
      <RouterLink to="/" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Home</RouterLink>
      <RouterLink to="/catalog" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Catalog</RouterLink>
      <RouterLink to="/blog" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Blog</RouterLink>
      <RouterLink to="/howtos" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">How-tos</RouterLink>
      <RouterLink to="/code" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Code</RouterLink>
      <RouterLink to="/team" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Team</RouterLink>
      <RouterLink to="/contact" class="block p-2 w-full text-center hover:bg-gray-700 transition-colors duration-200">Contact us</RouterLink>
    </nav>
  </header>
  <main>
    <slot />
  </main>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'

const menuOpen = ref(false)
const route = useRoute()

function toggleMenu() {
  menuOpen.value = !menuOpen.value
}

watch(() => route.path, () => {
  menuOpen.value = false
})
</script>
