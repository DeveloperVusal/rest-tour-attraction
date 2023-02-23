<script>
import { ref, inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    setup() {
        const inputRefs = ref([])
        const onRemoveClick = inject('requestRemove')

        return {
            inputRefs,
            onRemoveClick
        }
    },
    mounted() {
        const currentLanguageId = Number(this.$router.currentRoute.value.query.lang)
        this.languages = inject('loadLanguages')()
        this.languages.then(r => {
            this.languages = r

            if (currentLanguageId) this.currentLanguageId = currentLanguageId
            else  this.currentLanguageId = r[0].Id

            this.loadData()
        })
    },
    data() {
        return {
            isLoading: false,
            bodyData: [],
            currentLanguageId: 0,
            languages: null,
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

            const response = await axios.get('http://localhost:9000'+ReqUrls.country.get, {
                params: {
                    language_id: this.currentLanguageId,
                    full: true
                }
            })
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                this.bodyData = response.data.list
            }
        },
        changeLanguage(event) {
            const language_id = Number(event.target.value)

            if (language_id) {
                this.currentLanguageId = language_id
                this.loadData()
            }
        }
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Страны</h4>

                <div class="d-flex">
                    <select 
                        name="lang" 
                        class="form-select"
                        @change="changeLanguage($event)"
                    >
                        <option 
                            v-for="lang in languages" 
                            :value="lang.Id"
                            :selected="(lang.Id == currentLanguageId) ? true : null"
                        >
                            {{ lang.Name }}
                        </option>
                    </select>
                    <router-link 
                        :to="{
                            name: 'countries', 
                            params: { section: 'add' },
                            query: { lang: currentLanguageId }
                        }"
                        style="margin-left: 1rem;"
                    >
                        <button class="btn btn-primary">Добавить</button>
                    </router-link>
                </div>
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
                    <th scope="col">Название</th>
                    <th scope="col">Материк</th>
                    <th scope="col" style="text-align: center;">Видимость</th>
                    <th scope="col" style="text-align: center;">Архив</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in bodyData">
                    <td scope="row">{{ item.Id }}</td>
                    <td scope="col">{{ item.Name }}</td>
                    <td scope="col">{{ item.Continent.Name }}</td>
                    <td scope="col" align="center">
                        <div class="form-check form-check-inline" style="margin-right: -0.6rem;">
                            <input 
                                class="form-check-input" 
                                type="checkbox" 
                                :checked="(item.IsVisible) ? true : null"
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
                    <td scope="col" align="center">
                        <div class="d-flex justify-content-evenly">
                            <router-link 
                                :to="{
                                    name: 'countries', 
                                    params: { section: 'edit', id: item.Id }
                                }"
                                title="Редактировать"
                            >
                                <i class="bi bi-pencil-fill link-secondary"></i>
                            </router-link>
                            <div 
                                to="#"
                                title="Удалить"
                                @click="onRemoveClick('country', item.Id)"
                            >
                                <i class="bi bi-trash3-fill link-secondary"></i>
                            </div>
                        </div>
                    </td>
                </tr>
            </tbody>
        </table>
        <!-- <nav aria-label="Page navigation example">
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
        </nav> -->
    </div>
</template>