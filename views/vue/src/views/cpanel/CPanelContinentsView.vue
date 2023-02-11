<script>
import ListContinents from '@/components/modules/continents/List.vue'
import Sidebar from '@/components/inc/Sidebar.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListContinents,
    },
    mounted() {
        this.storeModules.setModuleActive('continents')
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
                <ListContinents />
            </div>
        </div>
    </div>
  </main>
</template>