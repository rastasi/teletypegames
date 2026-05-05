import type { RouteRecordRaw } from 'vue-router'

export const teamRouter: RouteRecordRaw[] = [
  { path: '/team', name: 'teamIndex', component: () => import('../page/team/TeamIndexPage.vue') },
]
