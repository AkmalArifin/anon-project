<template>
    <div class="modal">
        <div class="modal-container">
            <div class="header">
                <h4>Delete Confirmation</h4>
            </div>
            <div class="body">
                <p class="p2">Are you sure you want to delete this message?</p>
            </div>
            <div class="footer">
                <button class="p2 button cancel" @click="handleCancel">Cancel</button>
                <button class="p2 button delete" @click="handleDelete">Delete</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const props = defineProps<{id: number}>()
const emit = defineEmits<{cancel: [], delete: []}>()
const token = sessionStorage.getItem('jwtToken');

function handleCancel() {
    emit('cancel');
}

async function handleDelete() {
    await axios.delete(`http://localhost:8082/log/delete/${props.id}`, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    }).then(response => {
        router.push({name: 'dashboard'});
        emit('delete');
    }).catch(error => {
        console.error(error.response);
    })

}

</script>

<script lang="ts">
export default {
    name: "DeleteModal"
}
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 1; 
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  overflow: auto;
  background-color: rgb(0,0,0);
  background-color: rgba(0,0,0,0.4);
  align-items: center;
  justify-content: center;
}

.modal.show {
    display: flex;
}

.modal .modal-container {
    background-color: white;
    width: 420px;
    padding: 48px 72px;
    border: 1px solid var(--black-1);
    border-radius: 15px;
    animation: fade-in .5s;
}

.modal .modal-container .body {
    margin-top: 12px;
}

.modal .modal-container .footer {
    margin-top: 24px;

    display: flex;
    column-gap: 32px;
    align-items: center;
    justify-content: center;
}

.button {
    font-weight: bold;
    padding: 12px 24px;
    border: none;
    border-radius: 15px;
    transition: var(--transition);
    cursor: pointer;
}

.button:hover {
    scale: 1.05;
}

.button.cancel {
    background-color: var(--grey-light);
}

.button.delete {
    background-color: #d63e55;
}

</style>