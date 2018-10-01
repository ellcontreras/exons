<template>
    <div>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <h1 class="title is-1 has-text-centered">
            {{ exonsunity.title }}
        </h1>
        <p>
            {{ exonsunity.description }}
        </p>
        <router-link :to="`${this.exonsunity._id}/update`">Actualizar</router-link>
        <button @click="clickDelete()" class="button is-danger">Borrar</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'exonsunity',
    data() {
        return {
            exonsunity: [],
            error: ""
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/exonsunity/get/${this.$route.params.id}`).then(res => {
            this.exonsunity = res.data;
        }).catch(err => {
            console.log(err);
        })
    },
    methods: {
        clickDelete() {
            axios.delete("http://localhost:8080/api/exonsunity/delete", {data: {
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
