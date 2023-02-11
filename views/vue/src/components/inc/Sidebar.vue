<script>
import { useModulesStore } from '@/stores/modules'

export default {
    setup() {
        const storeModules = useModulesStore()

        return {
            storeModules
        }
    },
    mounted() {
        this.modules = this.storeModules.getList
    },
    watch: {
        'storeModules.getList': {
            handler(list) {
                this.modules = list
            },
            deep: true
        }
    },
    data() {
        return {
            modules: {}
        }
    }
}
</script>

<template>
    <div class="list-group">
        <template v-for="item in this.modules">
            <router-link :to="'/cpanel'+item.path" :class="'list-group-item list-group-item-action'+((item.active) ? ' active' : '')">
                {{ item.name }}
            </router-link>
        </template>
    </div>
</template>