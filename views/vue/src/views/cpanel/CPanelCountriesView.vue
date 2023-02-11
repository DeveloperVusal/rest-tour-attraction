<script>
import ListCountries from '@/components/modules/countries/List.vue'
import Sidebar from '@/components/inc/Sidebar.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListCountries,
    },
    mounted() {
        this.storeModules.setModuleActive('countries')
    },
    setup() {
        const storeModules = useModulesStore()

        return {
            storeModules
        }
    }
}
</script>

<template>
  <main>
    <div class="container-xl mt-3">
        <div class="row">
            <div class="col col-3">
                <Sidebar />
            </div>
            <div class="col">
                <ListCountries />
            </div>
        </div>
    </div>
  </main>
</template>