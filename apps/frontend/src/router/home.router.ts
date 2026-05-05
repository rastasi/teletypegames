import type { RouteRecordRaw } from 'vue-router'
import HomePage from '../page/home/HomePage.vue'

export const homeRouter: RouteRecordRaw[] = [
  { path: '/', name: 'home', component: HomePage },
]
