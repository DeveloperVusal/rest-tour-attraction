import { createRouter, createWebHistory } from 'vue-router'

import AuthView from '@/views/auth/AuthView.vue'
import CPanelLocationsView from '@/views/cpanel/CPanelLocationsView.vue'
import CPanelUsersView from '@/views/cpanel/CPanelUsersView.vue'
import CPanelContinentsView from '@/views/cpanel/CPanelContinentsView.vue'
import CPanelCountriesView from '@/views/cpanel/CPanelCountriesView.vue'
import CPanelLanguagesView from '@/views/cpanel/CPanelLanguagesView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/auth'
    },
    {
      path: '/auth',
      name: 'auth',
      component: AuthView
    },
    {
      path: '/cpanel',
      name: 'cpanel',
      redirect: '/cpanel/locations/list'
    },
    {
      path: '/cpanel/locations/list',
      name: 'cpanel_locations',
      component: CPanelLocationsView
    },
    {
      path: '/cpanel/users/list',
      name: 'cpanel_users',
      component: CPanelUsersView
    },
    {
      path: '/cpanel/continents/list',
      name: 'cpanel_continents',
      component: CPanelContinentsView
    },
    {
      path: '/cpanel/countries/list',
      name: 'cpanel_countries',
      component: CPanelCountriesView
    },
    {
      path: '/cpanel/languages/list',
      name: 'cpanel_languages',
      component: CPanelLanguagesView
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // }
  ]
})

export default router
