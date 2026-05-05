import type { RouteRecordRaw } from 'vue-router'

export const codeRouter: RouteRecordRaw[] = [
  { path: '/code', name: 'codeIndex', component: () => import('../page/code/CodeIndexPage.vue') },
]
