<script>
import { inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    mounted() {
        inject('loadGroups')().then(r => this.groups = r)
        inject('initValidateForms')()
    },
    data() {
        return {
            groups: null,
            isLoading: false,
            bodyData: {
                group_id: 0,
                username: '',
                password: '',
                email: '',
                name: '',
                surname: '',
                age: 0,
                gender: '',
                is_confirm: true,
            },
            warnings: []
        }
    },
    methods: {
        async submitForm() {
            this.warnings = []

            if (
                !this.bodyData.username.length ||
                !this.bodyData.password.length ||
                !this.bodyData.group_id
            ) {
                if (!this.bodyData.name.length) this.warnings.push('Заполните Логин пользователя')
                if (!this.bodyData.password.length) this.warnings.push('Заполните Пароль пользователя')
                if (!this.bodyData.group_id) this.warnings.push('Выберите группу пользователя')
                
                return false
            }

            this.isLoading = true
            this.warnings = []

            const response = await axios.post('http://localhost:9000'+ReqUrls.user.add, this.bodyData)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 201) {
                if (response.data.hasOwnProperty('user_id')) {
                    this.$router.push({name: 'users', params: {section: 'list'}})
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

            if (field === 'group_id' || field === 'age') value = Number(value)
            if (field === 'is_confirm') value = event.target.checked

            this.bodyData[field] = value
        }
    }
}
</script>

<template>
    <div class="container bg-secondary rounded border pl-3 pr-3 pt-2 pb-2 mb-3">
        <div class="row">
            <div class="col d-flex align-items-center justify-content-between">
                <h4 class="mb-0 text-white d-inline">Добавление пользователя</h4>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <form class="row g-3 needs-validation" novalidate @submit.prevent="submitForm()">
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Group" class="form-label">Группа*</label>
                    <select name="group_id" @change="setInputField($event, 'group_id')" class="form-select" id="valid-Group" required>
                        <template v-if="groups">
                            <option value="0">Выбрать</option>
                            <option 
                                v-for="group in groups" 
                                :value="group.Id"
                            >
                                {{ group.Name }}
                            </option>
                        </template>
                        <template v-else>
                            <option value="0">Пусто</option>
                        </template>
                    </select>
                    <div class="invalid-feedback">
                        Выберите Группу
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Username" class="form-label">Логин*</label>
                    <input type="text" @keyup="setInputField($event, 'username')" name="username" class="form-control" id="valid-Username" required>
                    <div class="invalid-feedback">
                        Заполните Логин пользователя
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Password" class="form-label">Пароль*</label>
                    <input type="text" @keyup="setInputField($event, 'password')" name="password" class="form-control" id="valid-Password" required>
                    <div class="invalid-feedback">
                        Заполните Пароль пользователя
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Email" class="form-label">Email*</label>
                    <input type="text" @keyup="setInputField($event, 'email')" name="email" class="form-control" id="valid-Email" required>
                    <div class="invalid-feedback">
                        Заполните Email пользователя
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Name" class="form-label">Имя</label>
                    <input type="text" @keyup="setInputField($event, 'name')" name="name" class="form-control" id="valid-Name">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Surname" class="form-label">Фамилия</label>
                    <input type="text" @keyup="setInputField($event, 'surname')" name="surname" class="form-control" id="valid-Surname">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Age" class="form-label">Возраст</label>
                    <input type="number" @keyup="setInputField($event, 'age')" name="age" min="14" max="200" class="form-control" id="valid-Age">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Gender" class="form-label">Пол</label>
                    <select name="gender" @change="setInputField($event, 'gender')" class="form-select" id="valid-Gender">
                        <option value="male">мужской</option>
                        <option value="female">женский</option>
                    </select>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" name="is_confirm" checked @change="setInputField($event, 'is_confirm')" id="invalidCheck-IsConfirm">
                        <label class="form-check-label" for="invalidCheck-IsConfirm">
                            Подтвержден
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