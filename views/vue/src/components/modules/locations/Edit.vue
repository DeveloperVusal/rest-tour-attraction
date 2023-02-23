<script>
import 'bootstrap-icons/font/bootstrap-icons.css'

import { inject } from 'vue'
import axios from 'axios'
import cookie from 'cookiejs'

import { ReqUrls } from '@/requstes'

export default {
    setup() {
        const loadCountries = inject('loadCountries')

        return {
            loadCountries
        }
    },
    mounted() {
        inject('loadLanguages')().then(r => this.languages = r)

        this.loadData(() => {
            this.loadCountries(this.sendBodyData.language_id).then(r => this.countries = r)
        })
        
        inject('initValidateForms')()
    },
    data() {
        return {
            itemId: Number(this.$router.currentRoute.value.params.id),
            languages: null,
            countries: null,
            isLoading: false,
            renderIsLoading: false,
            sendBodyData: {
                id: 0, 
                language_id: 0,
                country_id: 0,
                name: '',
                city: '',
                description: '',
                is_archive: false,
                is_visible: false,
            },
            renderBodyData: {
                Name: '',
                LanguageId: 0,
                CountryId: 0,
                Name: '',
                City: '',
                Description:'',
                IsArchive: false,
                IsVisible: false,
            },
            warnings: []
        }
    },
    methods: {
        async loadData(callback) {
            this.renderIsLoading = true
            
            const response = await axios.get('http://localhost:9000'+ReqUrls.location.get+'/'+this.itemId)
            
            setTimeout(() => {
                this.renderIsLoading = false
            }, 500)

            if (response.status === 200) {
                this.renderBodyData = response.data

                this.sendBodyData.name = response.data.Name
                this.sendBodyData.city = response.data.City
                this.sendBodyData.description = response.data.Description
                this.sendBodyData.language_id = response.data.LanguageId
                this.sendBodyData.country_id = response.data.CountryId
                this.sendBodyData.is_visible = response.data.IsVisible
                this.sendBodyData.is_archive = response.data.IsArchive

                callback()
            }
        },
        async submitForm() {
            this.isLoading = true
            this.warnings = []
            this.sendBodyData.id = this.itemId

            const accessToken = cookie.get('access_token')
            const response = await axios.patch('http://localhost:9000'+ReqUrls.location.update, this.sendBodyData, {
                headers: {
                    Authorization: `Bearer ${accessToken}`
                }
            })
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                if (response.data.status == 'success') {
                    this.$router.push({name: 'locations', params: {section: 'list'}, query: {lang: this.sendBodyData.language_id}})
                } else {
                    this.warnings.push(response.data)
                }
            } else {
                alert('Что-то пошло не так, повторите попытку позже!')
            }
        },
        setInputField(event, field) {
            let value = event.target.value

            if (field === 'language_id' || field === 'country_id') value = Number(value)
            if (
                field === 'is_visible' ||
                field === 'is_archive'
            ) value = event.target.checked

            this.sendBodyData[field] = value
        }
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Сохранение места</h4>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <template v-if="!renderIsLoading">
            <form class="row g-3 needs-validation" novalidate @submit.prevent="submitForm($event.target)">
                <div class="row mb-3">
                    <div class="col">
                        <label for="valid-Lang" class="form-label">Язык*</label>
                        <select name="language_id" @change="setInputField($event, 'language_id')" class="form-select" id="valid-Lang" required>
                            <template v-if="languages">
                                <option 
                                    v-for="lang in languages" 
                                    :value="lang.Id"
                                    :selected="(lang.Id == renderBodyData.LanguageId) ? true : null"
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
                        <label for="valid-Country" class="form-label">Страна*</label>
                        <select name="country_id" @change="setInputField($event, 'country_id')" class="form-select" id="valid-Country" required>
                            <template v-if="countries">
                                <option 
                                    v-for="country in countries" 
                                    :value="country.Id"
                                    :selected="(country.Id == renderBodyData.CountryId) ? true : null"
                                >
                                    {{ country.Name }}
                                </option>
                            </template>
                            <template v-else>
                                <option value="0">Пусто</option>
                            </template>
                        </select>
                        <div class="invalid-feedback">
                            Выберите Страну
                        </div>
                    </div>
                </div>
                <div class="row mb-3">
                    <div class="col">
                        <label for="valid-Name" class="form-label">Название*</label>
                        <input type="text" :value="renderBodyData.Name" @keyup="setInputField($event, 'name')" name="name" class="form-control" id="valid-Name" required>
                        <div class="invalid-feedback">
                            Заполните Название места
                        </div>
                    </div>
                </div>
                <div class="row mb-3">
                    <div class="col">
                        <label for="valid-City" class="form-label">Город*</label>
                        <input type="text" :value="renderBodyData.City" @keyup="setInputField($event, 'city')" name="city" class="form-control" id="valid-City" required>
                        <div class="invalid-feedback">
                            Заполните Город
                        </div>
                    </div>
                </div>
                <div class="row mb-3">
                    <div class="col">
                        <div class="mb-3">
                            <label for="valid-Description" class="form-label">Описание*</label>
                            <textarea class="form-control" id="valid-Description" rows="3" name="description" @keyup="setInputField($event, 'description')" required>{{ renderBodyData.Description }}</textarea>
                            <div class="invalid-feedback">
                                Заполните Описание
                            </div>
                        </div>
                    </div>
                </div>
                <div class="row mb-3">
                    <div class="col">
                        <div class="form-check">
                            <input class="form-check-input" type="checkbox" name="is_visible" :checked="(renderBodyData.IsVisible) ? true : null" @change="setInputField($event, 'is_visible')" id="validCheck-IsVisible">
                            <label class="form-check-label" for="validCheck-IsVisible">
                                Показывать
                            </label>
                        </div>
                        <div class="form-check">
                            <input class="form-check-input" type="checkbox" name="is_archive" :checked="(renderBodyData.IsArchive) ? true : null" @change="setInputField($event, 'is_archive')" id="validCheck-IsArchive">
                            <label class="form-check-label" for="validCheck-IsArchive">
                                Архив
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
                            <span v-else>Сохранить</span>
                        </button>
                    </div>
                </div>
            </form>
        </template>
        <template v-else>
            <span>
                <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                Loading...
            </span>
        </template>
    </div>
</template>