<template>
    <div>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <h1 class="title is-1 has-text-centered">
            {{ community.title }}
        </h1>
        <p>
            {{ community.description }}
        </p>
        <router-link :to="`${this.community._id}/update`">Actualizar</router-link>
        <button @click="clickDelete()" class="button is-danger">Borrar</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Community',
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
        clickDelete() {
            axios.delete("http://localhost:8080/api/community/delete", {data: {
                _id: this.$route.params.id
            }}).then(res => {
                this.$router.push('/');
            }).catch(error => {
                console.log(error);
                this.error = "Ocurrió un error al tratar de eliminar esta comunidad";
            });
        }
    }
}
</script>
