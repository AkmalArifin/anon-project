<template>
    <div>
        <input type="text" placeholder="Ask a question" v-model="question" @keyup.enter="sendClicked">
        <button type="submit" @click="sendClicked">Send</button>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const question = ref("")

async function sendClicked() {
    console.log(question.value)
    try {
        const data = {
            "username": route.params.username,
            "question": question.value
        };

        const response = await axios.post("http://localhost:8082/ask-log", data);
        console.log(response);
    } catch (e) {
        console.error(e)
    }
}

</script>

<script lang="ts">
export default {
    name: "AskView"
}
</script>

<style scoped>

</style>