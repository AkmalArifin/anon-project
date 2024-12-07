<template>
    <div>
        <NavbarVue/>
        <DeleteModal v-if="showDeleteModal" :id="deleteID" @delete="deleteMessage" @cancel="showDeleteModal = false"/>
        <div class="profile-container">
            <img id="profile-picture" :src="photoUrl" alt="photo profile">
            <div class="profile-name">
                <div class="username">
                    <h4>{{ user.username }}</h4>
                </div>
                <div class="count">
                    <p class="p1">{{ messages.length }} <font-awesome-icon icon="fa-solid fa-envelope"/></p>
                    
                </div>
            </div>
            <div class="button-container">
                <font-awesome-icon icon="fa-solid fa-share-nodes" class="button-icon" @click="shareClicked"/>
            </div>
            <div>
                <p class="p3 alert" v-if="showAlert">Link Copied!</p>
            </div>
        </div>
        <div class="messages-container">
            <div class="messages-navbar">
                <div class="nav">
                    <p class="p1 nav-button" :class="allFilter" @click="filter = 'all'">All</p>
                    <p class="p1 nav-button" :class="starFilter" @click="filter = 'star'">Starred</p>
                </div>
                <div class="sort">
                    <font-awesome-icon :icon="sortIcon" class="button-sort" @click="sortHandle"/>
                </div>
            </div>
            <div class="messages">
                <!-- <p v-for="m in messages" :key="m.id">{{m.message}}</p> -->
                <Message v-for="m, i in showedMessages" :key="m.id" :id="m.id" :index="i" :message="m.message" :date="m.created_at" :isStarred="m.is_starred" @star="updateStar" @delete="deleteClicked"/>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
// Component
import Message from '../components/Message.vue';
import NavbarVue from '../components/Navbar.vue';
import DeleteModal from '../components/DeleteModal.vue';

import axios from 'axios';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { parseJwt } from '../utils/jwt';

interface Message {
    id: number;
    user_id: number;
    message: string;
    is_starred: number;
    created_at: string;
};

const router = useRouter();

const messages = ref<Message[]>([]);
const photoUrl = ref("/profile-picture.png")
const showDeleteModal = ref(false);
const deleteID = ref(0);
const filter = ref("all");
const sortIcon = ref("fa-solid fa-sort-up");
const showAlert = ref(false);

const user = ref({
    id: 0,
    username: "",
    email: "",
    password: "",
    photo_profile: "",
    created_at: "",
    deleted_at: ""
})

const token = sessionStorage.getItem('jwtToken');
const parsedToken = parseJwt(token);

/** Get Data */

// get data user
axios.get(`http://localhost:8082/users/${parsedToken.id}`)
    .then(response => {
        user.value = response.data;
        
        if (user.value.photo_profile !== null && user.value.photo_profile !== "") {
            photoUrl.value = user.value.photo_profile;
        }
    }).catch(error => {
        console.error(error);
    });

// get messages log
axios.get("http://localhost:8082/log", {
    headers: {
        'Authorization': `Bearer ${token}`
    }
    }).then(response => {
        messages.value = response.data || [];
        messages.value = messages.value.reverse();
    }).catch(error => {
        console.error(error);

        router.push({name: "login"})
    })


/** Dynamic  Variables  */

const showedMessages = computed(() => {
    if (filter.value === 'all') {
        return messages.value
    } else if (filter.value === 'star')  {
        return messages.value.filter(message => message.is_starred === 1)
    }
});

const allFilter = computed(() => {
    if (filter.value === 'all') {
        return "active"
    } else {
        return ""
    }
})

const starFilter = computed(() => {
    if (filter.value === 'star') {
        return "active"
    } else {
        return ""
    }
})

/** Dynamic Function */

// make choice to share to other app
async function shareClicked() {
    const shareRoute = router.resolve({ name: "ask", params: {username: user.value.username}});
    const fullLink = window.location.origin + shareRoute.fullPath;

    navigator.clipboard.writeText(fullLink)
        .then(() => {
            showAlert.value = true;
            setTimeout(() => {
                showAlert.value = false;
            }, 2000)
        }).catch(error => {
            console.error(error)
        })
}

// send post to update star
async function updateStar(id: number) {
    let message = messages.value.find(message => message.id === id);
    
    if (!message) {
        console.error('message not found');
        return;
    }

    message.is_starred = message.is_starred === 1 ? 0 : 1; 

    axios.put(`http://localhost:8082/log/star/${id}`, "", {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }).then(response => {

    }).catch(error => {
        console.error(error.response);
    })
}

async function deleteClicked(id: number) {
    showDeleteModal.value = true;
    deleteID.value = id;
}

async function deleteMessage() {
    messages.value = messages.value.filter(message => message.id !== deleteID.value);
    showDeleteModal.value = false;
}

async function sortHandle() {
    messages.value = messages.value.reverse();

    if (sortIcon.value === 'fa-solid fa-sort-up') {
        sortIcon.value = 'fa-solid fa-sort-down';
    } else {
        sortIcon.value = 'fa-solid fa-sort-up';
    }
}

</script>

<script lang="ts">
export default {
    name: "DashboardView"
}
</script>

<style scoped>

.profile-container {
    padding-top: calc(70px + 60px);
    text-align: center;

    /* background-color: var(--black-2); */
}

.profile-container #profile-picture {
    border-radius: 100%;
    height: min(80vw, 200px);
    width: min(80vw, 200    px);
}

.profile-container .profile-name {
    display: flex;
    justify-content: center;
    align-items: center;
    column-gap: 12px;
}

.profile-container .count {
    letter-spacing: 2px;
    padding-top: 2px;
}

.profile-container .button-container {
    height: 64px;
    width: 64px;

    background: rgb(131,153,217);
    background: linear-gradient(157deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 100%);

    border-radius: 15px;

    position: relative;
    margin-top: 12px;
    left: 50%;
    transform: translateX(-50%);

    transition: var(--transition);

    cursor: pointer;
}

.profile-container .button-container:hover {
    scale: 1.05;
    transform: translateX(-48%);
}

.profile-container .button-container .button-icon {
    color: var(--white);
    font-size: 32px;

    position: absolute;

    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
}

.profile-container .alert {
    background-color: var(--black-2);
    color: var(--white);
    width: 120px;
    padding: 2px 4px;
    border-radius: 15px;
    
    position: absolute;
    left: calc(50% + 32px + 12px);
    top: 410px;
    animation: fade-in-out 2s linear;
}

/* Messages */

.messages-navbar {
    width: min(80vw, 672px);
    margin: auto;   
    padding: 0px 30px;

    display: flex;
    align-items: center;
    justify-content: space-between;
    column-gap: 160px;

    border-bottom: 1px solid var(--black-1);
}

.messages-navbar .nav {
    display: flex;
    column-gap: 50px;
}

.messages-navbar .nav > * {
    padding-bottom: 2px;
    position: relative;

    cursor: pointer;
}

.messages-navbar .nav .active:after {
    content: '';
    position: absolute;
    width: calc(100% + 10px);
    bottom: 0;
    left: -5px;
    right: 0;
    background: var(--color-secondary);
    height: 4px;
    border-radius: 2px 2px 2px 2px;
    animation: zoom-in .15s;
}

.messages-navbar .sort .button-sort {
    font-size: 20px;
    cursor: pointer;
}

.messages {
    display: flex;
    flex-flow: column;
    align-items: center;
    row-gap: 16px;

    margin-top: 24px;
}

</style>