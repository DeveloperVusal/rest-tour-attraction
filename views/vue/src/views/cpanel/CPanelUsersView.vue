<script>
import Sidebar from '@/components/inc/Sidebar.vue'
import ListUsers from '@/components/modules/users/List.vue'
import AddUsers from '@/components/modules/users/Add.vue'
import EditUsers from '@/components/modules/users/Edit.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListUsers,
        AddUsers,
        EditUsers,
    },
    mounted() {
        this.storeModules.setModuleActive('users')
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
                    <ListUsers />
                </template>
                <template v-else-if="$route.params.section == 'add'">
                    <AddUsers />
                </template>
                <template v-else-if="$route.params.section == 'edit'">
                    <EditUsers />
                </template>
            </div>
        </div>
    </div>
  </main>
</template>