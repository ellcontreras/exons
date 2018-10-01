<template>
    <div>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <h1 class="title is-1 has-text-centered">
            {{ Community.title }}
        </h1>
        <p>
            {{ Community.description }}
        </p>
        <router-link :to="`${this.Community._id}/update`">Actualizar</router-link>
        <button @click="clickDelete()" class="button is-danger">Borrar</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Community',
    data() {
        return {
            Community: [],
            error: ""
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/Community/get/${this.$route.params.id}`).then(res => {
            this.Community = res.data;
        }).catch(err => {
            console.log(err);
        })
    },
    methods: {
        clickDelete() {
            axios.delete("http://localhost:8080/api/Community/delete", {data: {
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
