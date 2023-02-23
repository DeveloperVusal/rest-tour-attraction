<script>
import Sidebar from '@/components/inc/Sidebar.vue'
import ListLanguages from '@/components/modules/languages/List.vue'
import AddLanguages from '@/components/modules/languages/Add.vue'
import EditLanguages from '@/components/modules/languages/Edit.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListLanguages,
        AddLanguages,
        EditLanguages
    },
    mounted() {
        this.storeModules.setModuleActive('languages')
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
                    <ListLanguages />
                </template>
                <template v-else-if="$route.params.section == 'add'">
                    <AddLanguages />
                </template>
                <template v-else-if="$route.params.section == 'edit'">
                    <EditLanguages />
                </template>
            </div>
        </div>
    </div>
  </main>
</template>