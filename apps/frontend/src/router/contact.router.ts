import type { RouteRecordRaw } from 'vue-router'

export const contactRouter: RouteRecordRaw[] = [
  { path: '/contact', name: 'contact', component: () => import('../page/contact/ContactPage.vue') },
]
