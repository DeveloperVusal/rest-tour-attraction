<script>
import { provide } from 'vue'
import { RouterView } from 'vue-router'
import axios from 'axios'
import cookie from 'cookiejs'

import { useModulesStore } from '@/stores/modules'
import $api from '@/http/axios'
import { ReqUrls } from '@/requstes'

export default {
    setup() {
        const loadLanguages = async function() {
            const response = await axios.get('http://localhost:9000'+ReqUrls.language.get)

            if (response.status === 200) {
                return response.data
            } else {
                return false
            }
        }
        const initValidateForms = function() {
            // Fetch all the forms we want to apply custom Bootstrap validation styles to
            const forms = document.querySelectorAll('.needs-validation')

            // Loop over them and prevent submission
            Array.from(forms).forEach(form => {
                form.addEventListener('submit', event => {
                    if (!form.checkValidity()) {
                        event.preventDefault()
                        event.stopPropagation()
                    }

                    form.classList.add('was-validated')
                }, false)
            })
        }
        const formatDate = function(tm) {
            const fmDate = new Date(tm)

            let day = Number(fmDate.getUTCDate()),
                month = Number(fmDate.getUTCMonth()+1),
                year = Number(fmDate.getUTCFullYear()),
                hour = Number(fmDate.getHours()),
                min = Number(fmDate.getMinutes())

            if (day < 10) day = '0'+day
            if (month < 10) month = '0'+month
            if (hour < 10) hour = '0'+hour
            if (min < 10) month = '0'+min

            return day+'.'+month+'.'+year+' в '+hour+':'+min
        }
        const requestRemove = async function(module, id)  {
            const con = confirm('Вы действительно хотите удалить?')

            if (con) {
                const response = await axios.delete('http://localhost:9000'+ReqUrls[module].delete+'/'+id)

                if (response.data.status === 'success') {
                    window.location.reload()
                } else {
                    return false
                }
            }
        }

        provide('requestRemove', requestRemove)
        provide('loadLanguages', loadLanguages)
        provide('formatDate', formatDate)
        provide('initValidateForms', initValidateForms)

        const storeModules = useModulesStore()

        return {
            storeModules
        }
    },
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
                    this.redirectAfterAuth()
                } else {
                    console.log('Unauthorized')
                    this.$router.push('/auth')
                }
            } else {
                console.log('Unauthorized')
                this.$router.push('/auth')
            }
        },
        redirectAfterAuth() {
            if (
                this.$router.currentRoute.value.name === 'cpanel' || 
                this.$router.currentRoute.value.name === 'auth'
            ) {
                let fowardLink = this.$router.currentRoute.value.path+this.storeModules.getList['locations'].path+'/list'

                this.$router.push(fowardLink)
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
