import AuthView from '@/views/auth/AuthView.vue'
import CPanelLocationsView from '@/views/cpanel/CPanelLocationsView.vue'
import CPanelUsersView from '@/views/cpanel/CPanelUsersView.vue'
import CPanelContinentsView from '@/views/cpanel/CPanelContinentsView.vue'
import CPanelCountriesView from '@/views/cpanel/CPanelCountriesView.vue'
import CPanelLanguagesView from '@/views/cpanel/CPanelLanguagesView.vue'
import CPanelGroupsView from '@/views/cpanel/CPanelGroupsView.vue'

const Routes = [
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
        children: [
            {
                path: 'locations/:section/:id?',
                name: 'locations',
                component: CPanelLocationsView,
            },
            {
                path: 'users/:section/:id?',
                name: 'users',
                component: CPanelUsersView
            },
            {
                path: 'continents/:section/:id?',
                name: 'continents',
                component: CPanelContinentsView
            },
            {
                path: 'countries/:section/:id?',
                name: 'countries',
                component: CPanelCountriesView
            },
            {
                path: 'languages/:section/:id?',
                name: 'languages',
                component: CPanelLanguagesView
            },
            {
                path: 'groups/:section/:id?',
                name: 'groups',
                component: CPanelGroupsView
            },
        ]
    },
]

export default Routes