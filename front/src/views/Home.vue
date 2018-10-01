<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            exonsunities
        </h1>
        <Card v-for="exonsunity in exonsunities" :key="exonsunity._id" 
            :id="exonsunity._id"
            :title="exonsunity.title" :description="exonsunity.description"/>
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
            exonsunities: []
        }
    },
    beforeMount() {
        axios.get('http://localhost:8080/api/exonsunity/get/all').then(res => {
            this.exonsunities = res.data;
        }).catch(err => {
            console.log(err);
        });
    }
}
</script>
