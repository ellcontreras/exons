<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            {{ community.title }}
        </h1>
        <p>
            {{ community.description }}
        </p>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Community',
    data() {
        return {
            community: []
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/community/get/${this.$route.params.id}`).then(res => {
            this.community = res.data.community;
        }).catch(err => {
            console.log(err);
        })
    }
}
</script>

