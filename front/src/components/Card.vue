<template>
    <div class="card">
        <div class="card-heade">
            <div class="card-header-title">{{ title }}</div>
        </div>
        <div class="card-content">
            {{ description }}
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props: {
        title: String,
        description: String
    }
}
</script>

