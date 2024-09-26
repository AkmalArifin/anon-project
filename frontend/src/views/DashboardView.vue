<template>
    <div>
        <div>
            <p v-for="q in questions" :key="q.id">{{q.question}}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'

interface Question {
    id: number;
    user_id: number;
    question: string;
    is_read: boolean;
    created_at: string;
}

const questions = ref<Question[]>()

const token = sessionStorage.getItem('jwtToken');

axios.get("http://localhost:8082/ask-log", {
    headers: {
        'Authorization': `Bearer ${token}`
    }
}).then(response => {
    questions.value = response.data
    console.log(questions.value)
}).catch(error => {
    console.error(error)
})

async function getData(event: Event) {
    
}

</script>

<script lang="ts">
export default {
    name: "DashboardView"
}
</script>

<style scoped>

</style>