<template>
    <div class="container">
        <div class="logo-container">
            <h1 class="logo">Hnn<span style="color: var(--color-primary); font-weight: bold;">?</span></h1>
        </div>
        <div class="input-container">
            <div class="top"></div>
            <img class="profile-picture" :src="photoUrl" alt="profile picture">
            <p class="p3 counter" :class="countColor">{{ count }}/250</p>
            <textarea id="message" class="p2" name="message" v-model="message"></textarea>
        </div>
        <div class="button-container" @click="sendClicked">
            <font-awesome-icon icon="fa-solid fa-paper-plane" class="button-icon"/>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { computed, ref } from 'vue';
import { useRoute } from 'vue-router';
import Swal from 'sweetalert2';

const route = useRoute();

const message = ref("");
const autoHeight = ref("");
const count = computed(() => {
    return message.value.length
})
const countColor = computed(() => {
    if (count.value > 250) {
        return "red"
    } else {
        return ""
    }
})


// Photo Profile
var photoUrl = "/profile-picture.png"

axios.get(`http://localhost:8082/users/username/${route.params.username}`)
    .then(response => {
        if (response.data.photo_profile !== null) {
            photoUrl = response.data.photo_profile
        }
    }).catch(error => {
        console.error(error)
    })

// Setting Textarea Auto Height
document.addEventListener('DOMContentLoaded', function () {
    const textarea = document.getElementById("message");

    if (textarea){
        textarea.addEventListener("input", function () {
            this.style.height = 'auto';
            this.style.height = Math.max(this.scrollHeight, 230) + 'px';
        })
    }
})

async function sendClicked() {
    if (count.value <= 0 || count.value > 250) {
        Swal.fire({
            title: "Fill this field and keep it under the limit.",
            icon: "error",
            timer: 2000,
            confirmButtonText: 'Close',
            confirmButtonColor: '#8399D9',
        });
    } else {
        const data = {
            "username": route.params.username,
            "message": message.value
        };

        await axios.post("http://localhost:8082/log", data)
            .then(response => {
                Swal.fire({
                    title: "Message sent!",
                    icon: "success",
                    timer: 2000,
                    confirmButtonText: 'Close',
                    confirmButtonColor: '#8399D9',
                    willClose: () => {
                        message.value = "";
                    }
                });
            }).catch(error => {
                Swal.fire({
                    title: "Please try again.",
                    text: error.response.data.message,
                    icon: "error",
                    timer: 2000,
                    confirmButtonText: 'Close',
                    confirmButtonColor: '#8399D9',
                });
                console.error(error)
            })
    }
}

</script>

<script lang="ts">
export default {
    name: "AskView"
}
</script>

<style scoped>
.container {
    height: 100vh;
    display: flex;
    flex-flow: column;
    justify-content: center;
    align-items: center;
}

.input-container {
    position: relative;

    margin-top: 48px;
}

.input-container .top {
    height: 66px;
    width: min(80vw, 672px);

    border-radius: 15px 15px 0px 0px;
    position: absolute;

    background: rgb(131,153,217);
    background: linear-gradient(113deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);
}

.input-container .profile-picture {
    position: absolute;
    height: 96px;
    width: 96px;
    border-radius: 100%;

    left: 50%;
    transform: translate(-50%, -50%);
}

.input-container .counter {
    position: absolute;

    bottom:12px;
    right: 36px;
}

.input-container .counter.red {
    color: red;
}

.input-container #message {
    width: min(80vw, 672px);
    min-height: 230px;
    padding: 82px 36px 44px 36px;

    overflow: hidden;
    resize: none;
    outline: none;

    border-color: var(--white);
    border-radius: 15px;

    box-shadow: var(--shadow);

    height: auto;
}

.button-container {
    height: 64px;
    width: 64px;

    background: rgb(131,153,217);
    background: linear-gradient(157deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);

    border-radius: 15px;

    position: relative;
    margin-top: 24px;

    transition: var(--transition);

    cursor: pointer;
}

.button-container:hover {
    scale: 1.05;
}

.button-container .button-icon {
    color: var(--white);
    font-size: 28px;

    position: absolute;

    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);

}

</style>