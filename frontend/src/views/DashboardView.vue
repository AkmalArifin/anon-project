<template>
    <div>
        <div>
            <p v-for="q in questions" :key="q.id">{{q.question}}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { parseJwt } from '../utils/jwt';

interface Question {
    id: number;
    user_id: number;
    question: string;
    is_read: boolean;
    created_at: string;
};

const router = useRouter();

const token = sessionStorage.getItem('jwtToken');

var jwt

if (token) {
    const jwt = parseJwt(token)
} else {
    router.push({name: "login"})
}

const user = axios.get(`http://localhost:8082/users/${jwt.id}`)

const questions = ref<Question[]>();


axios.get("http://localhost:8082/ask-log", {
    headers: {
        'Authorization': `Bearer ${token}`
    }
}).then(response => {
    questions.value = response.data;
    console.log(questions.value);
}).catch(error => {
    console.error(error);

    // please login first
})
</script>

<script lang="ts">
export default {
    name: "DashboardView"
}
</script>

<style scoped>

</style>