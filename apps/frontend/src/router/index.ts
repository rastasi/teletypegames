import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: HomeView },
    { path: '/blog', component: () => import('../views/BlogView.vue') },
    { path: '/blog/:slug(.*)', component: () => import('../views/BlogPostView.vue') },
    { path: '/catalog', component: () => import('../views/CatalogView.vue') },
    { path: '/catalog/:name', component: () => import('../views/CatalogItemView.vue') },
    { path: '/code', component: () => import('../views/CodeView.vue') },
    { path: '/contact', component: () => import('../views/ContactView.vue') },
    { path: '/howtos', component: () => import('../views/HowtosView.vue') },
    { path: '/team', component: () => import('../views/TeamView.vue') },
  ],
  scrollBehavior() {
    return { top: 0 }
  }
})
