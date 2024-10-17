<template>
    <div class="profile-page">
        <Navbar class="navbar"/>
        <div class="card data">

            <div class="profile-container" v-if="showData">
                <div class="photo">
                    <img id="photo" :src="photoUrl" alt="photo profile">
                </div>
                <div class="data">
                    <div class="data-container">
                        <span class="p2 title">Username</span>
                        <span class="p2 value">{{ user.username }}</span>
                    </div>
                    <div class="data-container">
                        <span class="p2 title">Email</span>
                        <span class="p2 value">{{ user.email }}</span>
                    </div>
                </div>
                <div class="button-container">
                    <button class="p2 button" @click="handleEdit">Edit</button>
                    <button class="p2 button" @click="handleChangePassword">Change Password</button>
                </div>
            </div>

            <div class="profile-container" v-if="showEdit">
                <div class="photo">
                    <img id="photo" :src="photoUrl" alt="photo profile">
                </div>
                <div class="data">
                    <div class="data-container">
                        <span class="p2 title">Username</span>
                        <div>
                            <input type="text" v-model="editUser.username" class="p2 value" @blur="v$.username.$touch">
                            <p>
                                <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="v$.username.$error"/>
                                <span class="error message" v-if="v$.username.$error">{{ v$.username.$errors[0].$message }}</span>
                            </p>
                            <p>
                                <font-awesome-icon class="error" icon="fa-solid fa-circle-exclamation" v-if="serverError.field.includes('username')"/>
                                <span class="error message" v-if="serverError.field.includes('username')">{{ serverError.value }} is already taken!</span>
                            </p>
                        </div>
                    </div>
                    <div class="data-container">
                        <span class="p2 title">Email</span>
                        <span class="p2 value">{{ user.email }}</span>
                    </div>
                </div>
                <div class="button-container">
                    <button class="p2 button" @click="handleCancel">Cancel</button>
                    <button class="p2 button" @click="handleSave">Save</button>
                </div>
            </div>
        </div>
        <div class="card change-password" v-if="showChangePassword"></div>
        <div class="card new-password" v-if="showNewPassword"></div>
    </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { parseJwt } from '../utils/jwt';
import axios from 'axios';
import Navbar from '../components/Navbar.vue';
import { useVuelidate } from '@vuelidate/core';
import { required, minLength, maxLength, helpers, sameAs } from '@vuelidate/validators';
import passwordRule from '../validators/password';
import { server } from 'typescript';

const photoUrl = ref("/profile-picture.png")
const showData = ref(true);
const showEdit = ref(false);
const showChangePassword = ref(false);
const showNewPassword = ref(false);
const user = ref({
    id: 0,
    username: "",
    email: "",
    password: "",
    photo_profile: "",
    created_at: "",
    deleted_at: ""
})
const editUser = reactive({
    id: 0,
    username: "",
    email: "",
    photo_profile: "",
    password: "",
    confirmPassword: ""
})

const serverError = reactive({
    field: "",
    value: ""
})
const rules = computed(() => ({
    username: {required, minLength: minLength(6), maxLength: maxLength(30)},
    password: {required, minLength: minLength(8), passwordRule: helpers.withMessage("Must contain at least 1 number", passwordRule)},
    confirmPassword: {required, sameAs: sameAs(editUser.password)}
}))

const v$ = useVuelidate(rules, editUser)

const token = sessionStorage.getItem('jwtToken');
const parsedToken = parseJwt(token);

/** Get Data */

axios.get(`http://localhost:8082/users/${parsedToken.id}`)
    .then(response => {
        user.value = response.data

        if (user.value.photo_profile !== null && user.value.photo_profile !== "") {
            photoUrl.value = user.value.photo_profile;
        }
    }).catch(error => {
        console.error(error);
    })

/** Function */
async function handleEdit() {
    editUser.id = user.value.id;
    editUser.username = user.value.username;
    editUser.email = user.value.email;
    editUser.photo_profile = user.value.photo_profile;

    showEdit.value = true;
    showData.value = false;
}

async function handleChangePassword() {

}

async function handleCancel() {
    showEdit.value = false;
    showData.value = true;
}

async function handleSave() {

    await v$.value.$validate;
    if (v$.value.username.$error) {
        return
    }
    const data = {
        "username": editUser.username,
        "email": editUser.email,
        "photo_profile": editUser.photo_profile,
    }

    await axios.put(`http://localhost:8082/users/${editUser.id}`, data, {
        headers: {
            "Authorization": `Bearer ${token}`
        }
    }).then(response => {
        user.value.username = editUser.username;
        user.value.email = editUser.email;
        user.value.photo_profile = editUser.photo_profile;
        serverError.field = "";
        serverError.value = "";

        showEdit.value = false;
        showData.value = true;
    }).catch(error => {
        console.error(error.response);
        serverError.field = error.response.data.field;
        serverError.value = error.response.data.value;
    })
}

</script>

<script lang="ts">
export default {
    name: "ProfileView"
}
</script>

<style scoped>
input {
    border: none;
    padding: 0px;
    width: auto;
}

input:focus {
  outline: none;         /* Remove the outline */
  box-shadow: none;      /* Remove any shadow */
}

.profile-page {
    padding-top: 1px;
}

.card {
    width: 672px;
    height: auto;
    border: 4px solid var(--color-primary);
    border-radius: 15px;

    margin: calc(72px + 60px) auto;
}

.profile-container {
    display: flex;
    flex-flow: column;
    align-items: center;
    justify-content: center;
    padding: 32px 0px;
}

.profile-container .photo {
    margin-bottom: 32px;
}

.profile-container #photo {
    height: 192px;
    width: 192px;
    border-radius: 100%;
}

.profile-container .data .data-container {
    display: grid;
    grid-template-columns: 150px 300px;
    grid-template-areas: 
    "title value";
    margin-bottom: 24px;
}

.profile-container .data .data-container .error {
    text-align: left;
    color: var(--color-error);
    white-space: inherit;
    margin-top: 8px;
}

.profile-container .data .data-container .error.message {
    margin-left: 8px;
}

.profile-container .data .data-container .title {
    grid-area: title;
    font-weight: 600;
}

.profile-container .data .data-container .value {
    grid-area: value;
}

.profile-container .button-container {
    display: flex;
    column-gap: 32px;
    margin-top: 12px;
    margin-bottom: 20px;
}

.profile-container .button-container .button {
    width: 160px;
    height: 48px;
    padding: 12px 0px;
    font-weight: 600;

    background: linear-gradient(var(--white),  var(--white)) padding-box,
    linear-gradient(143deg, rgba(131,153,217,1) 0%, rgba(211,114,137,1) 49%) border-box;
    border: solid 2px transparent;
    border-radius: 15px;

    cursor:pointer;
    transition: var(--transition);
}

.profile-container .button-container .button:hover {
    scale: 1.05;
}



</style>