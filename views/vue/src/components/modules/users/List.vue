<script>
import { ref } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    setup() {
        const inputRefs = ref([])

        return {
            inputRefs
        }
    },
    mounted() {
        this.loadData()
    },
    data() {
        return {
            isLoading: false,
            bodyData: []
        }
    },
    watch: {
        'inputRefs': {
            handler(refs) {
                refs.map(item => {
                    if (!item.checked) item.indeterminate = true
                })
            },
            deep: true
        }
    },
    methods: {
        async loadData() {
            this.isLoading = true

            const response = await axios.get('http://localhost:9000'+ReqUrls.user.get)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                this.bodyData = response.data
            }
        },
        formatDate(tm) {
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
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Пользователи</h4>
                <button class="btn btn-primary">Добавить</button>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border p-3 text-center">
        <template v-if="isLoading">
            <div class="spinner-grow text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </template>
        <table class="table text-start" v-else>
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Логин</th>
                    <th scope="col">Почта</th>
                    <th scope="col">Группа</th>
                    <th scope="col" style="text-align: center;">Подтвержден</th>
                    <th scope="col" style="text-align: center;">Архив</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in bodyData">
                    <td scope="row">{{ item.Id }}</td>
                    <td scope="col">{{ item.Username }}</td>
                    <td scope="col">{{ item.Email }}</td>
                    <td scope="col">{{ item.Group.Name }}</td>
                    <td scope="col" align="center">
                        <div class="form-check form-check-inline" style="margin-right: -0.6rem;">
                            <input 
                                class="form-check-input" 
                                type="checkbox" 
                                :checked="(item.IsConfirm) ? true : null"
                                disabled 
                                ref="inputRefs"
                            />
                        </div>
                    </td>
                    <td scope="col" align="center">
                        <div class="form-check form-check-inline" style="margin-right: -0.6rem;">
                            <input 
                                class="form-check-input" 
                                type="checkbox" 
                                :checked="(item.IsArchive) ? true : null"
                                disabled 
                                ref="inputRefs"
                            />
                        </div>
                    </td>
                    <td align="center">
                        <i class="bi bi-pencil-fill"></i> - 
                        <i class="bi bi-trash3-fill"></i>
                    </td>
                </tr>
            </tbody>
        </table>
        <nav aria-label="Page navigation example">
            <ul class="pagination justify-content-end">
                <li class="page-item">
                    <a class="page-link" href="#" aria-label="Previous">
                        <span aria-hidden="true">&laquo;</span>
                    </a>
                </li>
                <li class="page-item"><a class="page-link" href="#">1</a></li>
                <li class="page-item"><a class="page-link" href="#">2</a></li>
                <li class="page-item"><a class="page-link" href="#">3</a></li>
                <li class="page-item">
                    <a class="page-link" href="#" aria-label="Next">
                        <span aria-hidden="true">&raquo;</span>
                    </a>
                </li>
            </ul>
        </nav>
    </div>
</template>