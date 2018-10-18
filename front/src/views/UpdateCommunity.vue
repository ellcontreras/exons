<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            Actualizando una comunidad
        </h1>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <div class="field">
            <input type="text" v-model="Community.title" class="input" placeholder="title">
        </div>
        <div class="fiel">
            <textarea v-model="Community.description" class="textarea" placeholder="description"></textarea>
        </div>
        <br>
        <button @click="onClick()" class="button is-info">Actualizar</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Updateexonsunity',
    data() {
        return {
            Community: [],
            error: ""
        }
    },
    beforeMount() {
        axios.get(`http://localhost:8080/api/community/get/${this.$route.params.id}`).then(res => {
            this.Community = res.data;
        }).catch(err => {
            console.log(err);
        })
    },
    methods: {
        onClick() {
            var ls = localStorage.getItem('user');
            ls = JSON.parse(ls);

            let config = {
                headers: {
                    Authorization: `Bearer ${ls.token}`
                }
            };

            axios.put('http://localhost:8080/api/community/update', {
                _id: this.Community._id,
                title: this.Community.title,
                description: this.Community.description
            }, config).then(res => {
                this.$router.push(`/Community/${this.$route.params.id}`);
            }).catch(err => {
                console.log(err.response);
                this.error = "Ha ocurrido un error, intentalo después :("
            })
        }
    }
}
</script>
