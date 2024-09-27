<template>
    <div class="navbar-container">
        <div class="logo-container">
            <h3 class="logo" @click="logoClicked">Hnn<span style="color: var(--color-primary);">?</span></h3>
        </div>
        <div class="nav-container">
            <font-awesome-icon icon="fa-solid fa-gear" class="setting-icon" @click="settingClicked"/>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";

const router = useRouter();

async function logoClicked() {
    router.push({name: "home"})
}

async function settingClicked() {
    // router.push({name: "profile"})
}
</script>

<style scoped>
.navbar-container {
    display: flex;
    justify-content: space-around;
    align-items: center;
    column-gap: 160px;

    height: 70px;
    width: 100vw;
    background-color: var(--white);
    position: fixed;
    /* border-bottom: 1px solid var(--black-1); */
}

.logo-container {
    padding-top: 10px;
}

.logo {
    cursor: pointer;
}

.setting-icon {
    font-size: 24px;
    color: var(--black-2);
    transition: var(--transition);

    cursor: pointer;
}

.setting-icon:hover {
    color: var(--color-primary);
}
</style>