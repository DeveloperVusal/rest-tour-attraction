<script>
import Sidebar from '@/components/inc/Sidebar.vue'
import ListGroups from '@/components/modules/groups/List.vue'
import AddGroups from '@/components/modules/groups/Add.vue'
import EditGroups from '@/components/modules/groups/Edit.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListGroups,
        AddGroups,
        EditGroups,
    },
    mounted() {
        this.storeModules.setModuleActive('groups')
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
                    <ListGroups />
                </template>
                <template v-else-if="$route.params.section == 'add'">
                    <AddGroups />
                </template>
                <template v-else-if="$route.params.section == 'edit'">
                    <EditGroups />
                </template>
            </div>
        </div>
    </div>
  </main>
</template>