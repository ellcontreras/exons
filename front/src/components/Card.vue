<template>
    <router-link :to="`/community/${id}`">
        <div class="card">
            <div class="card-header">
                <div class="card-header-title">{{ title }}</div>
            </div>
            <div class="card-content">
                {{ description }}
            </div>
        </div>
    </router-link>
</template>

<script>
export default {
    name: 'Card',
    props: {
        id: String,
        title: String,
        description: String
    }
}
</script>

<style scoped>
    .card {
        margin-bottom: 1em;
    }
</style>

