<template>
    <div class="register-view">
        <div class="logo-container">
            <h1 class="logo">Hnn?</h1>
        </div>
        <div class="card">
            <div class="header">
                <h3>Login</h3>
            </div>
            <div class="body">
                <div class="input">
                    <input type="text" class="p2" v-model="username" placeholder="Username">
                    <input type="text" class="p2" v-model="email" placeholder="Email">
                    <input type="password" class="p2" v-model="password" placeholder="Password">
                </div>
                <div class="button-container">
                    <button class="register-button p2" @click="handleRegister">Register</button>
                </div>
            </div>
            <div class="footer">
                <p class="p3">Already have an account? <span class="login-button" @click="handleLogin">Login</span></p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const username = ref("");
const email = ref("");
const password = ref("");

async function handleLogin() {
    router.push({name: "login"});
}

async function handleRegister(event: Event) {
    event.preventDefault();
    const data = {
        "username": username.value,
        "email": email.value,
        "password": password.value
    }

    await axios.post("http://localhost:8082/signup", data)
        .then(response => {
            router.push({name: "login"});
        }).catch(error => {
            console.error(error.response);
        });
}

</script>

<script lang="ts">
export default {
    name: "LoginView"
}
</script>

<style scoped>

.register-view {
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

.register-view .logo-container {
    margin-top: -116px;
}

.register-view .logo-container .logo {
    color: var(--white);
}

.register-view .card {
    background: var(--white);
    height: 558px;
    width: min(510px, 80vw);
    padding: 48px min(96px, 10vw);
    border: none;
    border-radius: 15px;
    box-shadow: var(--shadow);
}

.register-view .card .header {
    text-align: center;
}

.register-view .card .body {
    margin-top: 48px;
}

.register-view .card .body .input {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    row-gap: 24px;
}

.register-view .card .body .input > * {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.register-view .card .body .button-container {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
} 

.register-view .card .body .register-button {
    background: rgb(131,153,217);
    background: linear-gradient(157deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);
    width: 151px;
    height: 48px;
    padding: 12px 34px;
    font-weight: bold;
    border: none;
    border-radius: 15px;
    color: var(--white);
    cursor: pointer;
    transition: var(--transition);
}

.register-view .card .body .register-button:hover {
    scale: 1.05;
}

.register-view .card .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
}

.register-view .card .footer .login-button {
    font-weight: bold;
    cursor: pointer;
}

</style>