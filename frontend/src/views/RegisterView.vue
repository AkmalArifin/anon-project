<template>
    <div class="register-view">
        <div class="logo-container">
            <h1 class="logo">Hnn?</h1>
        </div>
        <div class="card">
            <div class="header">
                <h3>Register</h3>
            </div>
            <div class="body">
                <div class="input">
                    <div class="input-container">
                        <input type="text" class="input-text p2" v-model="userInput.username" placeholder="Username">
                        <p>
                            <font-awesome-icon v-if="v$.username.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="v$.username.$error" :key="v$.username.$errors[0].$uid" class="error message">{{ v$.username.$errors[0].$message }} </span>
                        </p>
                        <p>
                            <font-awesome-icon v-if="serverError.field.includes('username')" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="serverError.field.includes('username')" class="error message">{{ serverError.value }} is already taken!</span>
                        </p>
                    </div>
                    <div class="input-container">
                        <input type="text" class="input-text p2" v-model="userInput.email" placeholder="Email">
                        <p>
                            <font-awesome-icon v-if="v$.email.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="v$.email.$error" :key="v$.email.$errors[0].$uid" class="error message">{{ v$.email.$errors[0].$message }} </span>
                        </p>
                        <p>
                            <font-awesome-icon v-if="serverError.field.includes('email')" icon="fa-solid fa-circle-exclamation" class="error" />
                            <span v-if="serverError.field.includes('email')" class="error message">{{ serverError.value }} is already taken!</span>
                        </p>
                    </div>
                    <div class="input-container">
                        <input type="password" class="input-text p2" v-model="userInput.password" placeholder="Password">
                        <font-awesome-icon v-if="v$.password.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                        <span v-if="v$.password.$error" :key="v$.password.$errors[0].$uid" class="error message">{{ v$.password.$errors[0].$message }} </span>
                    </div>
                    <div class="input-container">
                        <input type="password" class="input-text p2" v-model="userInput.confirmPassword" placeholder="Confirm Password">
                        <font-awesome-icon v-if="v$.confirmPassword.$error" icon="fa-solid fa-circle-exclamation" class="error" />
                        <span v-if="v$.confirmPassword.$error" :key="v$.confirmPassword.$errors[0].$uid" class="error message">{{ v$.confirmPassword.$errors[0].$message }} </span>
                    </div>
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
import { computed, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useVuelidate } from '@vuelidate/core'
import { required, email, minLength, maxLength, sameAs, helpers } from '@vuelidate/validators'
import passwordRule from '../validators/password'

const router = useRouter();

const userInput = reactive({
    username: "",
    email: "",
    password: "",
    confirmPassword: ""
});

const serverError = reactive({
    field: "",
    value: ""
})

const rules = computed(() => ({
    username: {required, minLength: minLength(6), maxLength: maxLength(30)},
    email: {required, email},
    password: {required, minLength: minLength(8), passwordRule: helpers.withMessage("Must contain at least 1 number", passwordRule)},
    confirmPassword: {required, sameAs: sameAs(userInput.password)}
}));

const v$ = useVuelidate(rules, userInput);

console.log(v$);

async function handleLogin() {
    router.push({name: "login"});
}

async function handleRegister(event: Event) {
    event.preventDefault();

    const result = await v$.value.$validate();
    if (!result) {
        console.log(v$.value.$errors);
        console.log(v$.value.confirmPassword.$errors[0].$message)
        return
    }

    const data = {
        "username": userInput.username,
        "email": userInput.email,
        "password": userInput.password
    }

    await axios.post("http://localhost:8082/signup", data)
        .then(response => {
            router.push({name: "login"});
        }).catch(error => {
            console.error(error.response);
            serverError.field = error.response.data.field
            serverError.value = error.response.data.value
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
    height: auto;
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

.register-view .card .body .input .input-text {
    width: min(320px, 60vw);
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.register-view .card .body .input .error {
   text-align: left;
   color: var(--color-error);
   white-space: inherit;
   margin-top: 8px;
}

.register-view .card .body .input .error.message {
   margin-left: 8px;
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