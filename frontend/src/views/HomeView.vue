<template>
    <div class="container">
        <h1>Hnn<span style="color: var(--color-primary)">?</span></h1>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();

router.push({name: "login"})

</script>

<script lang="ts">
export default {
    name: "HomeView"
}
</script>

<style scoped>
.container {
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
}

h1 {
    font-family: var(--ff-logo);
}

</style>