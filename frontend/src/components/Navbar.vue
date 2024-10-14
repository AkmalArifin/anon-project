<template>
    <div class="navbar-container">
        <div class="logo-container">
            <h3 class="logo" @click="logoClicked">Hnn<span style="color: var(--color-primary); font-weight: bold;">?</span></h3>
        </div>
        <div class="nav-container">
            <font-awesome-icon icon="fa-solid fa-bars" class="bars-icon" @click="handleBars" />
            <div class="modal" v-if="showModal">
                <ul class="modal-content">
                    <li class="item" @click="handleDashboard">Dashboard</li>
                    <li class="item" @click="handleProfile">Profile</li>
                    <li class="item" @click="handleLogOut">Log out</li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const showModal = ref(false);

async function logoClicked() {
    router.push({name: "home"});
}

async function handleBars() {
    showModal.value = !showModal.value;
}

async function handleDashboard() {
    router.push({name: "dashboard"});
}

async function handleProfile() {
    router.push({name: "profile"});
}

async function handleLogOut() {
    sessionStorage.removeItem('jwtToken');
    router.push({name: "login"})
}

</script>

<script lang="ts">
export default {
    name: "NavbarVue"
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
    z-index: 1;
    border-bottom: 1px solid var(--grey-light);
    box-shadow: var(--shadow);
    /* border-bottom: 1px solid var(--black-1); */
}

.logo-container {
    padding-top: 10px;
}

.logo {
    cursor: pointer;
}

.bars-icon {
    font-size: 24px;
    color: var(--black-2);
    transition: var(--transition);

    cursor: pointer;
}

.bars-icon:hover {
    color: var(--color-primary);
}

.modal {
    position: fixed;
    z-index: 1;
    background-color: var(--white);
    width: 160px;
    height: auto;
    border: 1px solid var(--grey-light);
    box-shadow: var(--shadow);
    margin: 23px 0px 0px -125px;
}

.modal .modal-content {
    display: flex;
    flex-flow: column;
    row-gap: 10px;
    padding: 24px 24px;
}

.modal .modal-content .item {
    list-style: none;
    cursor: pointer;   
}

.modal .modal-content .item:hover { 
    font-weight: 550;
}

</style>