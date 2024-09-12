<template>
    <div>
        <form action="">
            <input type="text" v-model="email">
            <input type="text" v-model="password">
            <button type="submit" @click="loginClicked">Login</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';

const email = ref("");
const password = ref("");

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
                    const token = response.data.token   
                    sessionStorage.setItem('jwtToken', token);
                }
            });
    } catch (e) {
        console.error(e)
    }
}


</script>

<script lang="ts">
export default {
    name: "LoginView"
}
</script>

<style scoped>

</style>