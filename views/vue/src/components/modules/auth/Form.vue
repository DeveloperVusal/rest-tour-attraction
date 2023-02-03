<script>
import './styles/form.scss'

import axios from 'axios'

export default {
    mounted() {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        const forms = document.querySelectorAll('.needs-validation')

        // Loop over them and prevent submission
        Array.from(forms).forEach(form => {
            form.addEventListener('submit', event => {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
    },
    data() {
        return {
            isLoading: false,
            bodyData: {
                username: '',
                passwd: ''
            },
            warnings: []
        }
    },
    methods: {
        async submitForm() {
            if (!this.bodyData.username.length || !this.bodyData.passwd.length) return false

            this.isLoading = true
            this.warnings = []

            const response = await axios.post('http://localhost:9000/api/auth', this.bodyData) // { withCredentials: true }
            
            setTimeout(() => {
                this.isLoading = false
            }, 500)

            if (response.status === 200) {
                if (response.data.status == 'warning') {
                    this.warnings.push(response.data.message)
                } else {
                    console.log(response.data)
                    localStorage.setItem('access_token', response.data.access_token)
                    localStorage.setItem('refresh_token', response.data.refresh_token)
                }
            }
        },
        setInputField(event, field) {
            this.bodyData[field] = event.target.value
        }
    }
}
</script>

<template>
    <form class="row bg-white p-5 shadow-sm rounded-2 needs-validation" novalidate @submit.prevent="submitForm()">
        <div class="text-center col mb-3 auth-title">Авторизация</div>
        <div class="mb-3">
            <label for="labelInputEmailLogin" class="form-label">Email или логин</label>
            <input type="text" class="form-control rounded-2" id="labelEmailLogin" required @keyup="setInputField($event, 'username')" />
            <div class="invalid-feedback">
                Пожалуйста введите Email или логин.
            </div>
        </div>
        <div class="mb-3">
            <label for="labelInputPasswd" class="form-label">Пароль</label>
            <input type="password" class="form-control rounded-2" id="labelInputPasswd" required @keyup="setInputField($event, 'passwd')" />
            <div class="invalid-feedback">
                Пожалуйста введите пароль.
            </div>
        </div>
        <div class="mb-3" v-if="warnings.length">
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
        <div class="col text-center">
            <button type="submit" class="btn btn-primary d-inline ps-4 pe-4">
                <span v-if="isLoading">
                    <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                    Loading...
                </span>
                <span v-else>Войти</span>
            </button>
        </div>
    </form>
</template>