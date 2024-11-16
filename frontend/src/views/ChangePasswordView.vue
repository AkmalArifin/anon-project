<template>
    <div class="change-password-page">
        <Navbar class="navbar"/>
        <div class="card" v-if="showVerify">
            <div class="header">
                <h4>Change Password</h4>
            </div>
            <div class="body">
                <div class="input">
                    <div class="input-container">
                        <input type="password" class="p2 input-text" v-model="userInput.password" placeholder="Password">
                        <p>
                            <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="serverError.field.includes('password')"/>
                            <span class="error message" v-if="serverError.field.includes('password')">{{ serverError.message }}</span>
                        </p>
                    </div>
                </div>
            </div>
            <div class="footer">
                <button class="p2 button" @click="handleCancel">Cancel</button>
                <button class="p2 button" @click="handleNext">Next</button>
            </div>
        </div>
        <div class="card" v-if="showNewPassword">
            <div class="header">
                <h4>Change Password</h4>
            </div>
            <div class="body">
                <div class="input">
                    <div class="input-container">
                        <input type="password" class="p2 input-text" v-model="userInput.newPassword" placeholder="Password" @blur="v$.newPassword.$touch()">
                        <p>
                            <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="v$.newPassword.$error"/>
                            <span class="error message" v-if="v$.newPassword.$error">{{ v$.newPassword.$errors[0].$message }}</span>
                        </p>
                    </div>
                    <div class="input-container">
                        <input type="password" class="p2 input-text" v-model="userInput.confirmPassword" placeholder="Confirm Password" @blur="v$.confirmPassword.$touch()" @keyup.enter="handleSave">
                        <p>
                            <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="v$.confirmPassword.$error"/>
                            <span class="error message" v-if="v$.confirmPassword.$error">{{ v$.confirmPassword.$errors[0].$message }}</span>
                        </p>
                        <p>
                            <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="serverError.field.includes('newPassword')"/>
                            <span class="error message" v-if="serverError.field.includes('newPassword')">{{ serverError.message }}</span>
                        </p>
                    </div>
                </div>
            </div>
            <div class="footer">
                <button class="p2 button" @click="handleSave">Save</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { reactive, computed, ref } from 'vue';
import Navbar from '../components/Navbar.vue';
import { useRouter } from 'vue-router';
import { useVuelidate } from '@vuelidate/core'
import { required, minLength, maxLength, sameAs, helpers } from '@vuelidate/validators'
import passwordRule from '../validators/password';
import { parseJwt } from '../utils/jwt';
import { server } from 'typescript';

const router = useRouter();
const props = defineProps<{
    email?: number
}>()

const email = ref("")
const showVerify = ref(true)
const showNewPassword = ref(false)
const userInput = reactive({
    password: "",
    newPassword: "",
    confirmPassword: ""
})
const serverError = reactive({
    field: "",
    message: ""
})

const rules = computed(() => ({
    newPassword: {required, minLength: minLength(8), passwordRule: helpers.withMessage("Must contain at least 1 number", passwordRule)},
    confirmPassword: {required, sameAs: sameAs(userInput.newPassword)}
}))

const v$ = useVuelidate(rules, userInput);

/** Get Data */

const token = sessionStorage.getItem("jwtToken");
const parsedToken = parseJwt(token)

axios.get(`http://localhost:8082/users/${parsedToken.id}`)
    .then(response => {
        email.value = response.data.email
    }).catch(error => {
        console.error(error);
    })

async function handleCancel() {
    router.push({name: "profile"})
}

async function handleNext() {
    const data = {
        email: email.value,
        password: userInput.password
    };

    await axios.post("http://localhost:8082/login", data)
        .then(response => {
            showVerify.value = false;
            showNewPassword.value = true;
        }).catch(error => {
            console.error(error.response)
            serverError.field = "password";
            serverError.message = error.response.data.errors;
        });
}

async function handleSave() {
    await v$.value.$validate();
    if (v$.value.newPassword.$error || v$.value.confirmPassword.$error) {
        return;
    }

    const data = {
        password: userInput.newPassword
    }

    await axios.put(`http://localhost:8082/users/password/${parsedToken.id}`, data, {
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then(response => {
        router.push({name: "profile"})
    }).catch(error => {
        console.error(error.response)
        serverError.field = "newPassword"
        serverError.message = error.response.data.message
    })
}


</script>

<script lang="ts">
export default {
    name: "ChangePasswordView"
}
</script>

<style scoped>
.change-password-page {
    padding-top: 1px;
}

.error {
    text-align: left;
    color: var(--color-error);
    white-space: inherit;
    margin-top: 8px;
}

.error.message {
    margin-left: 8px;
}

.card {
    width: min(510px, 80vw);
    height: auto;
    border: 4px solid var(--color-primary);
    border-radius: 15px;

    margin: calc(72px + 60px) auto;
    padding: 58px min(96px, 10vw);
}

.card .header {
    text-align: center;
}

.card .body {
    margin-top: 48px;
}

.card .body .input {
    display: flex;
    flex-flow: column;
    align-content: center;
    justify-content: center;
    row-gap: 24px;
}

.card .body .input .input-container .input-text {
    width: 100%;
    height: 58px;
    padding: 13px 24px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
}

.card .footer {
    margin-top: 36px;
    display: flex;
    align-items: center;
    justify-content: center;
    column-gap: 32px;
}

.card .footer .button {
    width: 160px;
    height: 48px;
    padding: 12px 0px;
    font-weight: 600;

    background: linear-gradient(var(--white),  var(--white)) padding-box,
    linear-gradient(143deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 49%) border-box;
    border: solid 2px transparent;
    border-radius: 15px;

    cursor:pointer;
    transition: var(--transition);
}

.card .footer .button:hover {
    scale: 1.05
}

</style>