<script setup lang="ts">
import { ref, computed } from 'vue'
defineProps({
    text: String,
});
const dataTest = "これはデータです！"
const count = ref<number>(0)
</script>


<template>
    <p>{{ text }}</p>
    <p>{{ dataTest }}</p>
    <div>
        <button @click="count++">ここをクリック</button>
        {{ count }}
    </div>
</template>