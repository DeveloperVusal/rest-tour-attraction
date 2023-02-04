import axios from 'axios'
import cookie from 'cookiejs'

import { ReqUrls } from '@/requstes'
// import { buildParamsUri } from './functions'

const $api = axios.create({
    // withCredentials: true,
    baseURL: 'http://localhost:9000'
})

$api.interceptors.request.use((config) => {
    const accessToken = cookie.get('access_token')

    config.headers.Authorization = `Bearer ${(accessToken.length) ? accessToken : ''}`
    config.headers.crossDomain = true

    return config
})

$api.interceptors.response.use((config) => {
    return config;
}, async (error) => {
    if (error.response.status === 401) {
        if (cookie.get('access_token').length) {
            let response = await $api.get(ReqUrls.account.refresh, {
                params: {
                    token: localStorage.getItem('refresh_token')
                }
            })

            if (response.status === 200) {
                if (response.data.hasOwnProperty('access_token')) {
                    cookie.set('access_token', response.data.access_token, {
                        'expires': 30, 
                        'path': '/', 
                        'domain':'', 
                        'httponly': true 
                    })
                    localStorage.setItem('refresh_token', response.data.refresh_token)
                } else {
                    if (window.location.pathname != '/auth') window.location.href = '/auth'
                }
            }
        }
    }
})

export default $api