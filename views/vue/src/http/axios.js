import axios from 'axios'
import Cookies from 'js-cookie'

import { ReqUrls } from '@/requstes'
// import { buildParamsUri } from './functions'

const $api = axios.create({
    withCredentials: true,
    baseURL: 'http://localhost:9000'
})

$api.interceptors.request.use((config) => {
    const accessToken = Cookies.get('access_token')

    config.headers.Authorization = `Bearer ${(accessToken.length) ? accessToken : ''}`
    config.headers.crossDomain = true

    return config
})

$api.interceptors.response.use((config) => {
    return config;
}, async (error) => {
    if (error.response.status === 401) {
        if (Cookies.get('access_token').length) {
            let response = await $api.get(ReqUrls.account.refresh)

            response = await response.json()

            if (response.status === 'success') {
                cookie('access_token', response.data.access_token, { 'expires': 30, 'path': '/', 'domain':'', 'httponly': true })
                localStorage.setItem('refresh_token', response.data.refresh_token)

                console.log(response.data)
                // const params = buildParamsUri()

                // if (typeof params === 'object' && params.redirect) {
                //     window.location.href = params.redirect
                // } else {
                //     window.location.href = process.env.VUE_APP_DEFAULT_REDIRECT
                // }
            }
        }
    }
})

export default $api