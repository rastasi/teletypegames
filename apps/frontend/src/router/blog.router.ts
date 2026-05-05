import type { RouteRecordRaw } from 'vue-router'

export const blogRouter: RouteRecordRaw[] = [
  { path: '/blog', name: 'blogIndex', component: () => import('../page/blog/BlogIndexPage.vue') },
  { path: '/blog/:slug(.*)', name: 'blogShow', component: () => import('../page/blog/BlogShowPage.vue') },
]
