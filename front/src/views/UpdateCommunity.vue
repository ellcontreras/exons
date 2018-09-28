<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            Actualizando una comunidad
        </h1>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <div class="field">
            <input type="text" v-model="community.title" class="input" placeholder="title">
        </div>
        <div class="fiel">
            <textarea v-model="community.description" class="textarea" placeholder="description"></textarea>
        </div>
        <br>
        <button @click="onClick()" class="button is-info">Actualizar</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'UpdateCommunity',
    data() {
        return {
            community: [],
            error: ""
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/community/get/${this.$route.params.id}`).then(res => {
            this.community = res.data;
        }).catch(err => {
            console.log(err);
        })
    },
    methods: {
        onClick() {
            axios.put('http://localhost:8080/api/community/update', {
                _id: this.community._id,
                title: this.community.title,
                description: this.community.description
            }).then(res => {
                this.$router.push(`/community/${this.$route.params.id}`);
            }).catch(err => {
                console.log(err);
                this.error = "Ha ocurrido un error, intentalo después :("
            })
        }
    }
}
</script>
