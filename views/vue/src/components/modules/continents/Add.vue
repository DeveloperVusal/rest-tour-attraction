<script>
import { inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    mounted() {
        this.languages = inject('loadLanguages')()
        this.languages.then(r => this.languages = r)

        inject('initValidateForms')()
    },
    data() {
        return {
            languages: null,
            isLoading: false,
            bodyData: {
                language_id: 2,
                name: '',
                is_visible: true,
            },
            warnings: []
        }
    },
    methods: {
        async submitForm() {
            if (!this.bodyData.name.length) return false

            this.isLoading = true
            this.warnings = []

            const response = await axios.post('http://localhost:9000'+ReqUrls.continent.add, this.bodyData)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 201) {
                if (response.data.hasOwnProperty('continent_id')) {
                    this.$router.push({name: 'continents', params: {section: 'list'}})
                } else {
                    this.warnings.push(response.data)
                }
            } else if (response.status === 200) {
                this.warnings.push(response.data)
            } else {
                alert('Что-то пошло не так, повторите попытку позже!')
            }
        },
        setInputField(event, field) {
            let value = event.target.value

            if (field === 'language_id') value = Number(value)

            this.bodyData[field] = value
        }
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Добавление материка</h4>
                <!-- <button class="btn btn-primary">Добавить</button> -->
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <form class="row g-3 needs-validation" novalidate @submit.prevent="submitForm()">
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Lang" class="form-label">Язык</label>
                    <select name="language_id" @change="setInputField($event, 'language_id')" class="form-select" id="valid-Lang" required>
                        <template v-if="languages">
                            <option v-for="lang in languages" :value="lang.Id">
                                {{ lang.Name }}
                            </option>
                        </template>
                        <template v-else>
                            <option value="0">Пусто</option>
                        </template>
                    </select>
                    <div class="invalid-feedback">
                        Выберите Язык записи
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Name" class="form-label">Название</label>
                    <input type="text" @keyup="setInputField($event, 'name')" name="name" class="form-control" id="valid-Name" required>
                    <div class="invalid-feedback">
                        Заполните Название материка
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" name="is_visible" checked @change="setInputField($event, 'is_visible')" id="invalidCheck">
                        <label class="form-check-label" for="invalidCheck">
                            Показывать
                        </label>
                    </div>
                </div>
            </div>
            <div class="row mb-3" v-if="warnings.length">
                <div class="col">
                    <div class="alert alert-danger p-1" role="alert">
                        <ul class="list-group list-group-flush">
                            <li 
                                class="list-group-item list-group-item-danger"
                                v-for="item in warnings"
                            >
                                {{ item }}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <button class="btn btn-primary" type="submit">
                        <span v-if="isLoading">
                            <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                            Loading...
                        </span>
                        <span v-else>Добавить</span>
                    </button>
                </div>
            </div>
        </form>
    </div>
</template>