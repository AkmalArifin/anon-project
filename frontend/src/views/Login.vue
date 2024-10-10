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
import { useRouter } from 'vue-router';

const router = useRouter();

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

</style>