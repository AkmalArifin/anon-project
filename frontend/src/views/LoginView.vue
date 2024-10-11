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
                    <input type="text" class="p2" v-model="email" placeholder="Email">
                    <input type="password" class="p2" v-model="password" placeholder="Password">
                </div>
                <div class="add">
                    <div class="checkbox p3">
                        <input type="checkbox" name="remember" id="remember">
                        <label for="remember">  Remember me</label>
                    </div>
                    <p class="p3 forgot-button">Forgot password?</p>
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
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const email = ref("");
const password = ref("");


async function handleRegister() {
    router.push({name: 'register'})
}

async function loginClicked(event: Event) {
    event.preventDefault();
    try {
        const data = {
            "email": email.value,
            "password": password.value
        }

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
            });
    } catch (e) {
        console.error(e);
    }
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

.login-view .card .body .input > * {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
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