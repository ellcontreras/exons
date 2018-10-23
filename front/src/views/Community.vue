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
        <div v-if="userLoged">
            <router-link :to="`${this.Community._id}/update`">Actualizar</router-link>
            <button @click="clickDelete()" class="button is-danger">Borrar</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'exonsunity',
    data() {
        return {
            Community: [],
            error: "",
            userLoged: false
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/community/get/${this.$route.params.id}`).then(res => {
            this.Community = res.data;

            let user = JSON.parse(localStorage.getItem("user")).user;

            if (user._id === this.Community.user) {
                this.userLoged = true;
            }
        }).catch(err => {
            console.log(err.response);
        });
    },
    methods: {
        clickDelete() {
            var ls = localStorage.getItem('token');
            ls = JSON.parse(ls);
            
            axios.delete("http://localhost:8080/api/community/delete", {
                data: {
                    _id: this.$route.params.id
                },
                headers: {
                    Authorization: `Bearer ${ls.token}`
                }
            }).then(res => {
                this.$router.push('/');
            }).catch(error => {
                console.log(error);
                this.error = "Ocurrió un error al tratar de eliminar esta comunidad";
            });
        }
    },
}
</script>
