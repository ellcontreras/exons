<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            exonsunities
        </h1>
<<<<<<< HEAD
        <Card v-for="Community in exonsunities" :key="Community._id" 
            :id="Community._id"
            :title="Community.title" :description="Community.description"/>
=======
        <Card v-for="exonsunity in exonsunities" :key="exonsunity._id" 
            :id="exonsunity._id"
            :title="exonsunity.title" :description="exonsunity.description"/>
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de
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
        axios.get('http://localhost:8080/api/community/get/all').then(res => {
            this.exonsunities = res.data;
        }).catch(err => {
            console.log(err);
        });
    }
}
</script>
