<template>
    <div>
        <div class="hero is-primary">
            <div class="hero-body">
                <h1 class="title is-1">{{ user.name }}</h1> 
                <h2 class="subtitle is-2">@{{ user.username }}</h2>
                <p>{{ user.email }}</p>
            </div>
        </div>
        <hr>
        <h2 class="subtitle is-2 has-text-centered">
            Comunidades de {{ user.username }}
        </h2>
        
        <div class="columns" v-for="community in communities" :key="community._id">
            <div v-for="c in community" :key="c._id" class="column">
                <Card :id="c._id" :title="c.title" 
                    :description="c.description"/>
            </div>
        </div>
    </div>
</template>

<script>
import Card from '../components/Card';
import axios from 'axios';

export default {
    name: 'Profile',
    components: {Card},
    data() {
        return {
            user: [],
            communities: []
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/user/${this.$route.params.id}`)
        .then(res => {
            this.user = res.data;
        }).catch(error => {
            console.log(error.response);
        })

        axios.get(`http://localhost:8080/api/community/get/all/user/${this.$route.params.id}`)
        .then(res => {
            this.communities = res.data;

            let arrays = [], size = 3;

            while (this.communities.length > 0)
                arrays.push(this.communities.splice(0, size));

            this.communities = arrays;
            console.log(this.communities);
        }).catch(err => {
            console.log(err.response);
        });
    }
}
</script>
