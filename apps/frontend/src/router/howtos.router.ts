import type { RouteRecordRaw } from 'vue-router'

export const howtosRouter: RouteRecordRaw[] = [
  { path: '/howtos', name: 'howtosIndex', component: () => import('../page/howtos/HowtosIndexPage.vue') },
]
