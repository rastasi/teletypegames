import { createRouter, createWebHistory } from 'vue-router'
import { homeRouter } from './home.router'
import { blogRouter } from './blog.router'
import { catalogRouter } from './catalog.router'
import { codeRouter } from './code.router'
import { contactRouter } from './contact.router'
import { howtosRouter } from './howtos.router'
import { teamRouter } from './team.router'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    ...homeRouter,
    ...blogRouter,
    ...catalogRouter,
    ...codeRouter,
    ...contactRouter,
    ...howtosRouter,
    ...teamRouter,
  ],
  scrollBehavior() {
    return { top: 0 }
  },
})
