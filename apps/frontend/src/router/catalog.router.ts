import type { RouteRecordRaw } from 'vue-router'

export const catalogRouter: RouteRecordRaw[] = [
  { path: '/catalog', name: 'catalogIndex', component: () => import('../page/catalog/CatalogIndexPage.vue') },
  { path: '/catalog/:name', name: 'catalogShow', component: () => import('../page/catalog/CatalogShowPage.vue') },
]
