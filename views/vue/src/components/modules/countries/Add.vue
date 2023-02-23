<script>
import { inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    mounted() {
        inject('loadLanguages')().then(r => this.languages = r)
        inject('loadContinents')(this.bodyData.language_id).then(r => this.continents = r)
        inject('initValidateForms')()
    },
    data() {
        return {
            currentLanguageId: this.$router.currentRoute.value.query.lang,
            languages: null,
            continents: null,
            isLoading: false,
            bodyData: {
                language_id: Number(this.$router.currentRoute.value.query.lang) ?? 2,
                continent_id: 0,
                name: '',
                is_visible: true,
            },
            warnings: []
        }
    },
    methods: {
        async submitForm() {
            this.warnings = []

            if (
                !this.bodyData.name.length ||
                !this.bodyData.continent_id
            ) {
                if (!this.bodyData.name.length) this.warnings.push('Заполните Название материка')
                if (!this.bodyData.continent_id) this.warnings.push('Выберите Континент')
                
                return false
            }

            this.isLoading = true
            this.warnings = []

            const response = await axios.post('http://localhost:9000'+ReqUrls.country.add, this.bodyData)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 201) {
                if (response.data.hasOwnProperty('country_id')) {
                    this.$router.push({name: 'countries', params: {section: 'list'}, query: {lang: this.bodyData.language_id}})
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

            if (field === 'language_id' || field === 'continent_id') value = Number(value)
            if (field === 'is_visible') value = event.target.checked

            this.bodyData[field] = value
        }
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Добавление страны</h4>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <form class="row g-3 needs-validation" novalidate @submit.prevent="submitForm()">
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Lang" class="form-label">Язык*</label>
                    <select name="language_id" @change="setInputField($event, 'language_id')" class="form-select" id="valid-Lang" required>
                        <template v-if="languages">
                            <option 
                                v-for="lang in languages" 
                                :value="lang.Id"
                                :selected="(lang.Id == currentLanguageId) ? true : null"
                            >
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
                    <label for="valid-Continent" class="form-label">Континент*</label>
                    <select name="continent_id" @change="setInputField($event, 'continent_id')" class="form-select" id="valid-Continent" required>
                        <template v-if="continents">
                            <option value="0">Выбрать</option>
                            <option 
                                v-for="continent in continents" 
                                :value="continent.Id"
                            >
                                {{ continent.Name }}
                            </option>
                        </template>
                        <template v-else>
                            <option value="0">Пусто</option>
                        </template>
                    </select>
                    <div class="invalid-feedback">
                        Выберите Континент
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Name" class="form-label">Название*</label>
                    <input type="text" @keyup="setInputField($event, 'name')" name="name" class="form-control" id="valid-Name" required>
                    <div class="invalid-feedback">
                        Заполните Название страны
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