<script>
import ListLocations from '@/components/modules/locations/List.vue'
import Sidebar from '@/components/inc/Sidebar.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListLocations,
    },
    mounted() {
        this.storeModules.setModuleActive('locations')
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
                <Sidebar module="locations" />
            </div>
            <div class="col">
                <ListLocations />
            </div>
        </div>
    </div>
  </main>
</template>