import Vue from 'vue'
import VueRouter from 'vue-router'

import BookIndex from '../pages/books/index.vue'
import BookIdPage from '../pages/books/id.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/api/books',
    component: BookIndex
  },
  {
    path: '/api/books/:id',
    component: BookIdPage
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
