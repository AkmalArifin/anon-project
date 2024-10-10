<template>
    <div class="message-container" :class="color" >
        <p class="p3 date">{{ formattedDate }}</p>
        <p class="p2 message">{{ props.message }}</p>
        <div class="action-button">
            <div class="button-1">
                <font-awesome-icon icon="fa-solid fa-star" class="button star" :class="starYellow" @click="handleStar"/>
                <font-awesome-icon icon="fa-solid fa-paper-plane" class="button share"/>
            </div>
            <div class="button-2">
                <font-awesome-icon icon="fa-solid fa-trash" class="button trash" @click="handleDelete"/>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps({
    id: Number,
    index: Number,
    message: String,
    date: String,
    isStarred: Number
})

const emit = defineEmits<{
    star: [id: number]
    delete: [id: number]
}>();

const starYellow = computed(() => {
    if (props.isStarred) {
        return "yellow"
    } else {
        return ""
    }
});

const color = computed(() => {
    return (props.index ?? 0) % 2 == 0 ? 'even' : 'odd'
})


const dateObject = new Date(props.date ?? '');

const formattedDate = new Intl.DateTimeFormat('en-GB', {
    'day': '2-digit',
    'month': 'short',
    'year': 'numeric'
}).format(dateObject);

function handleStar() {
    emit('star', props.id ?? 0);
}

function handleDelete() {
    emit('delete', props.id ?? 0);
}

</script>

<script lang="ts">
export default {
    name: "Message"
}
</script>

<style scoped>
.message-container {
    width: 672px;
    height: auto;
    border-radius: 15px;

    padding: 18px 36px;
}

.message-container.even {
    border: 4px solid var(--color-primary);
}
.message-container.odd {
    border: 4px solid var(--color-secondary);
}

.message-container .date {
    color: var(--grey-dark);
}

.message-container .message {
    margin-top: 8px;
}

.message-container .action-button {
    margin-top: 16px;
    margin-bottom: 2px;
    font-size: 16px;

    display: flex;
    align-items: center;
    justify-content: space-between;
}

.message-container .action-button .button {
    cursor: pointer;
    transition: var(--transition);
}

.message-container .action-button .button:hover {
    scale: 1.2;
}

.message-container .action-button .share {
    margin-left: 16px;
}

.message-container .action-button .trash {
    color: #C7253E;
}

.message-container .action-button .yellow {
    color: #FCCD2A;
}


</style>