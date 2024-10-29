<template>
    <div class="reset-password-view">
        <div class="logo-container">
            <h1 class="logo">Hnn?</h1>
        </div>
        <div class="card" v-if="showReset">
            <div class="header">
                <h4>Reset Password</h4>
            </div>
            <div class="body">
                <div class="input">
                    <input type="password" class="p2" v-model="password" placeholder="Password">
                    <input type="password" class="p2" v-model="confirmPassword" placeholder="Confirm Password" @keyup.enter="handleResetPassword">
                </div>
                <div class="button-container">
                    <button class="reset-password-button p2" @click="handleResetPassword">Reset Password</button>
                </div>
            </div>
        </div>
        <div class="card" v-if="showUnauthorized">
            <div class="header">
                <h4>Unauthorized</h4>
            </div>
            <div class="body instruction">
                <div class="">
                    <p class="p2">You're unauthorized, please check again your url.</p>
                </div>
                <div class="button-container">
                    <button class="reset-password-button p2" @click="handleBack">Back</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { parseJwt } from '../utils/jwt';
import axios from 'axios';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();

const password = ref("");
const confirmPassword = ref("");
const showReset = ref(false);
const showUnauthorized = ref(true);

const id = route.query.id as string
const token = route.query.token as string

const data = {
    "id": parseInt(id ,10),
    "token": token
}

console.log(data)

axios.post("http://localhost:8082/reset-password/verify", data)
    .then(response => {
        console.log(response.data)
        sessionStorage.setItem('jwtToken', response.data);
    })
    .catch(error => {
        showUnauthorized.value = true
        showReset.value = false

        console.error(error.response)
    })

const jwtToken = sessionStorage.getItem('jwtToken')
const parsedToken = parseJwt(jwtToken);

async function handleResetPassword(event: Event) {
    event.preventDefault();
    
    if (password.value != confirmPassword.value) {
        alert("Passwords do not match!")
        return
    }

    const data = {
        "password": password.value
    }

    await axios.put(`http://localhost:8082/users/password/${parsedToken.id}`, data, {
        headers: {
            'Authorization': `Bearer ${jwtToken}`
        }
    }).then(response => {
        router.push({name: 'login'})
    }).catch(error => {
        alert("error")
    })
}

async function handleBack() {
    router.push({name: "login"})
}

</script>

<script lang="ts">
export default {
    name: "ResetPasswordView"
}
</script>

<style scoped>

.reset-password-view {
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

.reset-password-view .logo-container {
    margin-top: -116px;
}

.reset-password-view .logo-container .logo {
    color: var(--white);
}

.reset-password-view .card {
    background: var(--white);
    height: auto;
    width: min(510px, 80vw);
    padding: 48px min(96px, 10vw);
    border: none;
    border-radius: 15px;
    box-shadow: var(--shadow);
}

.reset-password-view .card .header {
    text-align: center;
}

.reset-password-view .card .body {
    margin-top: 40px;
}

.reset-password-view .card .body .input {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    row-gap: 24px;
}

.reset-password-view .card .body .input > * {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.reset-password-view .card .body .button-container {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
    margin-bottom: 12px;
} 

.reset-password-view .card .body .reset-password-button {
    background: rgb(131,153,217);
    background: linear-gradient(157deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);
    width: auto;
    height: 48px;
    padding: 12px 48px;
    font-weight: bold;
    border: none;
    border-radius: 15px;
    color: var(--white);
    cursor: pointer;
    transition: var(--transition);
}

.reset-password-view .card .body .reset-password-button:hover {
    scale: 1.05;
}

</style>