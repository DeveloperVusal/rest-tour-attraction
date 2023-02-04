<script>
import cookie from 'cookiejs'
import { RouterView } from 'vue-router'

import $api from '@/http/axios'
import { ReqUrls } from '@/requstes'

export default {
    created() {
        this.isAuth()
    },
    methods: {
        async isAuth() {
            if (cookie.get('access_token').length) {
                const response = await $api.get(ReqUrls.account.auth_verify)
                
                // Если все успешно
                if (response.status === 200) {
                    console.log('Authorized successfuly')
                    this.$router.push('/cpanel')
                } else {
                    console.log('Unauthorized')
                    this.$router.push('/auth')
                }
            } else {
                console.log('Unauthorized')
                this.$router.push('/auth')
            }
        }
    }
}
</script>

<template>
    <RouterView />
</template>

<style scoped>
    @import './assets/styles/fonts.scss';
</style>
