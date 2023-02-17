<script>
import ListUsers from '@/components/modules/users/List.vue'
import Sidebar from '@/components/inc/Sidebar.vue'

import { useModulesStore } from '@/stores/modules'

export default {
    components: {
        Sidebar,
        ListUsers,
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
                <!-- <template v-else-if="$route.params.section == 'add'">
                    <AddUsers />
                </template>
                <template v-else-if="$route.params.section == 'save'">
                    <SaveUsers />
                </template> -->
            </div>
        </div>
    </div>
  </main>
</template>