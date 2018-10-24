<template>
    <div class="hero is-primary">
        <div class="hero-body">
            <h1 class="title is-1">{{ user.name }}</h1> 
            <h2 class="subtitle is-2">@{{ user.username }}</h2>
            <p>{{ user.email }}</p>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Profile',
    data() {
        return {
            user: []
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/user/${this.$route.params.id}`)
            .then(res => {
                this.user = res.data;
            }).catch(error => {
                console.log(error.response);
            })
    }
}
</script>
