<script>
import Sidebar from '@/components/inc/Sidebar.vue'
import ListCountries from '@/components/modules/countries/List.vue'
import AddCountries from '@/components/modules/countries/Add.vue'
import EditCountries from '@/components/modules/countries/Edit.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListCountries,
        AddCountries,
        EditCountries,
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
                <template v-if="$route.params.section == 'list'">
                    <ListCountries />
                </template>
                <template v-else-if="$route.params.section == 'add'">
                    <AddCountries />
                </template>
                <template v-else-if="$route.params.section == 'edit'">
                    <EditCountries />
                </template>
            </div>
        </div>
    </div>
  </main>
</template>