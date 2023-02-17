<script>
import Sidebar from '@/components/inc/Sidebar.vue'
import ListContinents from '@/components/modules/continents/List.vue'
import AddContinents from '@/components/modules/continents/Add.vue'
import EditContinents from '@/components/modules/continents/Edit.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListContinents,
        AddContinents,
        EditContinents,
    },
    mounted() {
        this.storeModules.setModuleActive('continents')
    },
    setup() {
        const storeModules = useModulesStore()

        return {
            storeModules
        }
    },
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
                <template v-if="$route.params.section == 'list'">
                    <ListContinents />
                </template>
                <template v-else-if="$route.params.section == 'add'">
                    <AddContinents />
                </template>
                <template v-else-if="$route.params.section == 'edit'">
                    <EditContinents />
                </template>
            </div>
        </div>
    </div>
  </main>
</template>