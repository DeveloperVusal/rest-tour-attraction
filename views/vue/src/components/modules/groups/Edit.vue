<script>
import { inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    mounted() {
        this.loadData()
        inject('initValidateForms')()
    },
    data() {
        return {
            itemId: Number(this.$router.currentRoute.value.params.id),
            isLoading: false,
            renderIsLoading: false,
            sendBodyData: {
                id: 0, 
                name: '',
                is_visible: false,
                is_archive: false,
            },
            renderBodyData: {
                Name: '',
                IsArchive: false,
                IsVisible: false,
            },
            warnings: []
        }
    },
    methods: {
        async loadData() {
            this.renderIsLoading = true
            
            const response = await axios.get('http://localhost:9000'+ReqUrls.group.get+'/'+this.itemId)
            
            setTimeout(() => {
                this.renderIsLoading = false
            }, 500)

            if (response.status === 200) {
                this.renderBodyData = response.data

                this.sendBodyData.name = response.data.Name
                this.sendBodyData.is_visible = response.data.IsVisible
                this.sendBodyData.is_archive = response.data.IsArchive
            }
        },
        async submitForm() {
            this.isLoading = true
            this.warnings = []
            this.sendBodyData.id = this.itemId

            const response = await axios.patch('http://localhost:9000'+ReqUrls.group.update, this.sendBodyData)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                if (response.data.status == 'success') {
                    this.$router.push({name: 'groups', params: {section: 'list'}})
                } else {
                    this.warnings.push(response.data)
                }
            } else {
                alert('Что-то пошло не так, повторите попытку позже!')
            }
        },
        setInputField(event, field) {
            let value = event.target.value
            
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
                <h4 class="mb-0 text-white d-inline">Редактирование групы</h4>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <template v-if="!renderIsLoading">
            <form class="row g-3 needs-validation" novalidate @submit.prevent="submitForm()">
                <div class="row mb-3">
                    <div class="col">
                        <label for="valid-Name" class="form-label">Название*</label>
                        <input type="text" @keyup="setInputField($event, 'name')" :value="renderBodyData.Name" name="name" class="form-control" id="valid-Name" required>
                        <div class="invalid-feedback">
                            Заполните Название группы
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