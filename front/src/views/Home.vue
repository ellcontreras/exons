<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            Comunidades
        </h1>

        <hr>

        <div class="columns" v-for="community in communities" :key="community._id">
            <div v-for="c in community" :key="c._id" class="column">
                <Card :id="c._id" :title="c.title" 
                    :description="c.description"/>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import Card from '@/components/Card'

export default {
    name: 'Home',
    components: { Card },
    data() {
        return {
            communities: []
        }
    },
    beforeMount() {
        axios.get('http://localhost:8080/api/community/get/all').then(res => {
            this.communities = res.data;

            let arrays = [], size = 3;

            while (this.communities.length > 0)
                arrays.push(this.communities.splice(0, size));

            this.communities = arrays;
        }).catch(err => {
            console.log(err.response);
        });
    }
}
</script>
