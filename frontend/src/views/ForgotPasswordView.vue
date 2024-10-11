<template>
    <div class="forgot-password-view">
        <div class="logo-container">
            <h1 class="logo">Hnn?</h1>
        </div>
        <div class="card" v-if="showRequest">
            <div class="header">
                <h4>Forgot Password</h4>
            </div>
            <div class="body">
                <div class="input">
                    <input type="text" class="p2" v-model="email" placeholder="Email">
                </div>
                <div class="button-container">
                    <button class="forgot-password-button p2" @click="forgotPasswordClicked">Next</button>
                </div>
            </div>
            <div class="footer">
                <p class="p3"><span class="cancel-button" @click="handleCancel">Back to login</span></p>
            </div>
        </div>
        <div class="card" v-if="showNext">
            <div class="header">
                <h4>Check your email</h4>
            </div>
            <div class="body instruction">
                <div class="">
                    <p class="p2">We've send reset password instructions to your email. Please follow the instructions.</p>
                    <p class="gap"></p>
                    <p class="p2">If it doesn't arrive soon, check your spam folder or <span class="send-again-button" @click="forgotPasswordClicked">send the email again</span>.</p>
                </div>
                <div class="button-container">
                    <button class="forgot-password-button p2" @click="handleCancel">Back</button>
                </div>
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
const showRequest = ref(true);
const showNext = ref(false);


async function handleCancel() {
    router.push({name: 'login'})
}

async function forgotPasswordClicked(event: Event) {
    event.preventDefault();

    const data = {
        "email": email.value
    }
    
    axios.post("http://localhost:8082/reset-password", data)

    showRequest.value = false
    showNext.value = true
}

</script>

<script lang="ts">
export default {
    name: "ForgotPasswordView"
}
</script>

<style scoped>

.forgot-password-view {
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

.forgot-password-view .logo-container {
    margin-top: -116px;
}

.forgot-password-view .logo-container .logo {
    color: var(--white);
}

.forgot-password-view .card {
    background: var(--white);
    height: auto;
    width: min(510px, 80vw);
    padding: 48px min(72px, 10vw);
    border: none;
    border-radius: 15px;
    box-shadow: var(--shadow);
}

.forgot-password-view .card .header {
    text-align: center;
}

.forgot-password-view .card .body {
    margin-top: 48px;
}
.forgot-password-view .card .body.instruction {
    margin-top: 16px;
}

.forgot-password-view .card .body.instruction .send-again-button {
    font-weight: bold;
    cursor: pointer;
}

.forgot-password-view .card .body .input {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    row-gap: 24px;
}

.forgot-password-view .card .body .input > * {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.forgot-password-view .card .body .add {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 12px;
}

.forgot-password-view .card .body .add .checkbox {
    display: flex;
    align-items: center;
    column-gap: 8px;
}

.forgot-password-view .card .body .add .checkbox #remember {
    cursor: pointer;
}

.forgot-password-view .card .body .add .forgot-button {
    cursor: pointer;
}

.forgot-password-view .card .body .button-container {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
} 

.forgot-password-view .card .body .forgot-password-button {
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

.forgot-password-view .card .body .forgot-password-button:hover {
    scale: 1.05;
}

.forgot-password-view .card .footer {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 24px;
}

.forgot-password-view .card .footer .cancel-button {
    cursor: pointer;
}

</style>