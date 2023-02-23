import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useModulesStore = defineStore('modules', () => {
    const list = ref({})
    // const prefix = '/cpanel'

    list.value = {
        locations: {name: 'Места', path: '/locations', active: true},
        countries: {name: 'Страны', path: '/countries', active: false},
        continents: {name: 'Материки', path: '/continents', active: false},
        languages: {name: 'Языки', path: '/languages', active: false},
        groups: {name: 'Группы', path: '/groups', active: false},
        users: {name: 'Пользователи', path: '/users', active: false},
    }
    const getList = computed(() => list.value)

    const setModuleActive = (module) => {
        for (let key in list.value) list.value[key].active = false

        for (let key in list.value) {
            if (key == module) {
                list.value[key].active = true
            }
        }
    }

    return { list, getList, setModuleActive }
})
