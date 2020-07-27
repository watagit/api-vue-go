import Vue from 'vue'
import VueRouter from 'vue-router'

import BookIndex  from '../pages/books/index.vue'
import BookIdPage from '../pages/books/id.vue'
import BookEdit   from '../pages/books/edit.vue'
import BookCreate from '../pages/books/new.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/api/books',
    component: BookIndex
  },
  {
    path: '/api/books/new',
    component: BookCreate
  },
  {
    path: '/api/books/:id',
    component: BookIdPage
  },
  {
    path: '/api/books/:id/edit',
    component: BookEdit
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
