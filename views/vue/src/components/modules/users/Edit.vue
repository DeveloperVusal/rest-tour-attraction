<script>
import { inject } from 'vue'
import axios from 'axios'
import 'bootstrap-icons/font/bootstrap-icons.css'

import { ReqUrls } from '@/requstes'

export default {
    mounted() {
        inject('loadGroups')().then(r => this.groups = r)
        this.loadData()
        inject('initValidateForms')()
    },
    data() {
        return {
            itemId: Number(this.$router.currentRoute.value.params.id),
            groups: null,
            isLoading: false,
            renderIsLoading: false,
            sendBodyData: {
                id: 0, 
                group_id: 0,
                username: '',
                password: '',
                email: '',
                name: '',
                surname: '',
                age: 0,
                gender: '',
                is_confirm: false,
                is_blocked: false,
                is_archive: false,
            },
            renderBodyData: {
                GroupId: 0,
                Username: '',
                Password: '',
                Email: '',
                Name: '',
                Surname: '',
                Age: 0,
                Gender: '',
                IsConfirm: false,
                IsArchive: false,
                IsBlocked: false,
            },
            warnings: []
        }
    },
    methods: {
        async loadData() {
            this.renderIsLoading = true
            
            const response = await axios.get('http://localhost:9000'+ReqUrls.user.get+'/'+this.itemId)
            
            setTimeout(() => {
                this.renderIsLoading = false
            }, 500)

            if (response.status === 200) {
                this.renderBodyData = response.data

                this.sendBodyData.group_id = response.data.GroupId
                this.sendBodyData.username = response.data.Username
                this.sendBodyData.email = response.data.Email
                this.sendBodyData.name = response.data.Name
                this.sendBodyData.surname = response.data.Surname
                this.sendBodyData.gender = response.data.Gender
                this.sendBodyData.age = response.data.Age
                this.sendBodyData.is_confirm = response.data.IsConfirm
                this.sendBodyData.is_blocked = response.data.IsBlocked
                this.sendBodyData.is_archive = response.data.IsArchive
            }
        },
        async submitForm() {
            this.isLoading = true
            this.warnings = []
            this.sendBodyData.id = this.itemId

            const response = await axios.patch('http://localhost:9000'+ReqUrls.user.update, this.sendBodyData)
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                if (response.data.status == 'success') {
                    this.$router.push({name: 'users', params: {section: 'list'}})
                } else {
                    this.warnings.push(response.data)
                }
            } else {
                alert('Что-то пошло не так, повторите попытку позже!')
            }
        },
        setInputField(event, field) {
            let value = event.target.value

            if (field === 'group_id' || field === 'age') value = Number(value)
            if (
                field === 'is_blocked' ||
                field === 'is_archive' ||
                field === 'is_confirm'
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
                <h4 class="mb-0 text-white d-inline">Сохранение пользователя</h4>
            </div>
        </div>
    </div>
    <div class="container bg-white rounded border pt-4 p-3">
        <template v-if="!renderIsLoading">
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
                                :selected="(group.Id == renderBodyData.GroupId) ? true : null"
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
                    <input type="text" @keyup="setInputField($event, 'username')" name="username" :value="renderBodyData.Username" class="form-control" id="valid-Username" required>
                    <div class="invalid-feedback">
                        Заполните Логин пользователя
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Password" class="form-label">Новый пароль</label>
                    <input type="text" @keyup="setInputField($event, 'password')" name="password" class="form-control" id="valid-Password">

                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Email" class="form-label">Email*</label>
                    <input type="text" @keyup="setInputField($event, 'email')" name="email" :value="renderBodyData.Email" class="form-control" id="valid-Email" required>
                    <div class="invalid-feedback">
                        Заполните Email пользователя
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Name" class="form-label">Имя</label>
                    <input type="text" @keyup="setInputField($event, 'name')" name="name" :value="renderBodyData.Name" class="form-control" id="valid-Name">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Surname" class="form-label">Фамилия</label>
                    <input type="text" @keyup="setInputField($event, 'surname')" name="surname" :value="renderBodyData.Surname" class="form-control" id="valid-Surname">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Age" class="form-label">Возраст</label>
                    <input type="number" @keyup="setInputField($event, 'age')" name="age" :value="renderBodyData.Age" min="14" max="200" class="form-control" id="valid-Age">
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <label for="valid-Gender" class="form-label">Пол</label>
                    <select name="gender" @change="setInputField($event, 'gender')" class="form-select" id="valid-Gender">
                        <option value="male" :selected="(renderBodyData.Gender == 'male') ? true : null">мужской</option>
                        <option value="female" :selected="(renderBodyData.Gender == 'female') ? true : null">женский</option>
                    </select>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" name="is_confirm" :checked="(renderBodyData.IsConfirm) ? true : null" @change="setInputField($event, 'is_confirm')" id="invalidCheck-IsConfirm">
                        <label class="form-check-label" for="invalidCheck-IsConfirm">
                            Подтвержден
                        </label>
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" name="is_blocked" :checked="(renderBodyData.IsBlocked) ? true : null" @change="setInputField($event, 'is_blocked')" id="invalidCheck-IsBlocked">
                        <label class="form-check-label" for="invalidCheck-IsBlocked">
                            Блокировка
                        </label>
                    </div>
                </div>
            </div>
            <div class="row mb-3">
                <div class="col">
                    <div class="form-check">
                        <input class="form-check-input" type="checkbox" name="is_archive" :checked="(renderBodyData.IsArchive) ? true : null" @change="setInputField($event, 'is_archive')" id="invalidCheck-IsArchive">
                        <label class="form-check-label" for="invalidCheck-IsArchive">
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