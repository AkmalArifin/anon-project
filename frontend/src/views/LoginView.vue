<template>
    <div class="login-view">
        <div class="logo-container">
            <h1 class="logo">Hnn?</h1>
        </div>
        <div class="card">
            <div class="header">
                <h3>Login</h3>
            </div>
            <div class="body">
                <div class="input">
                    <div class="input-container">
                        <input type="text" class="input-text p2" v-model="userInput.email" placeholder="Email">
                        <font-awesome-icon v-if="v$.email.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                        <span v-if="v$.email.$error" :key="v$.email.$errors[0].$uid" class="error message">{{ v$.email.$errors[0].$message }} </span>
                    </div>
                    <div class="input-container">
                        <input type="password" class="input-text p2" v-model="userInput.password" placeholder="Password">
                        <p>
                            <font-awesome-icon v-if="v$.password.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="v$.password.$error" :key="v$.password.$errors[0].$uid" class="error message">{{ v$.password.$errors[0].$message }} </span>
                        </p>
                        <p>
                            <font-awesome-icon v-if="serverError !== ''" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="serverError !== ''" class="error message">{{ serverError }} </span>
                        </p>
                    </div>
                </div>
                <div class="add">
                    <div class="checkbox p3">
                        <input type="checkbox" name="remember" id="remember">
                        <label for="remember">  Remember me</label>
                    </div>
                    <p class="p3 forgot-button" @click="handelForgotPassword">Forgot password?</p>
                </div>
                <div class="button-container">
                    <button class="login-button p2" @click="loginClicked">Login</button>
                </div>
            </div>
            <div class="footer">
                <p class="p3">Don't have an account? <span class="register-button" @click="handleRegister">Register</span></p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { reactive, computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useVuelidate } from '@vuelidate/core';
import { required, email } from '@vuelidate/validators';
import { server } from 'typescript';

const router = useRouter();
const userInput = reactive({
    email: "",
    password: ""
});
const serverError = ref("")

const rules = computed(() => ({
    email: {required, email},
    password: {required}
}));


const v$ = useVuelidate(rules, userInput);

async function handleRegister() {
    router.push({name: 'register'});
}

async function handelForgotPassword() {
    router.push({name: "forgot-password"});
}

async function loginClicked(event: Event) {
    event.preventDefault();

    const result = await v$.value.$validate()

    if (!result) {
        console.log(v$.value.$errors)
        return
    }

    const data = {
        "email": userInput.email,
        "password": userInput.password
    };

    const response = axios.post("http://localhost:8082/login", data)
        .then(response => {
            if (response.status == 200) {
                const data : {
                    message: string,
                    token: string,
                } = response.data 
                sessionStorage.setItem('jwtToken', data.token);

                router.push({name: 'dashboard'});
            }   
        }).catch(error => {
            console.log(error.response.data);
            serverError.value = error.response.data.errors;
        });
}

</script>

<script lang="ts">
export default {
    name: "LoginView"
}
</script>

<style scoped>

.login-view {
    background: rgb(131,153,217);
    background: linear-gradient(337deg, rgba(131,153,217,1) 25%, rgba(211,114,137,1) 100%);
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-flow: column;
    align-items: center;
    justify-content: center;
    row-gap: 32px;
}

.login-view .logo-container {
    margin-top: -116px;
}

.login-view .logo-container .logo {
    color: var(--white);
}

.login-view .card {
    background: var(--white);
    height: 493px;
    width: min(510px, 80vw);
    padding: 48px min(96px, 10vw);
    border: none;
    border-radius: 15px;
    box-shadow: var(--shadow);
}

.login-view .card .header {
    text-align: center;
}

.login-view .card .body {
    margin-top: 48px;
}

.login-view .card .body .input {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    row-gap: 24px;
}

.login-view .card .body .input .input-text {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.login-view .card .body .input .error {
   text-align: left;
   color: var(--color-error);
   white-space: inherit;
   margin-top: 8px;
}

.login-view .card .body .input .error.message {
   margin-left: 8px;
}

.login-view .card .body .add {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 12px;
}

.login-view .card .body .add .checkbox {
    display: flex;
    align-items: center;
    column-gap: 8px;
}

.login-view .card .body .add .checkbox #remember {
    cursor: pointer;
}

.login-view .card .body .add .forgot-button {
    cursor: pointer;
}

.login-view .card .body .button-container {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
} 

.login-view .card .body .login-button {
    background: rgb(131,153,217);
    background: linear-gradient(157deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);
    width: 151px;
    height: 48px;
    padding: 12px 48px;
    font-weight: bold;
    border: none;
    border-radius: 15px;
    color: var(--white);
    cursor: pointer;
    transition: var(--transition);
}

.login-view .card .body .login-button:hover {
    scale: 1.05;
}

.login-view .card .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
}

.login-view .card .footer .register-button {
    font-weight: bold;
    cursor: pointer;
}

</style>