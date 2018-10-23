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
        if (!localStorage.getItem("user")) {
            this.$router.push('/');
        }

        axios.get(`http://localhost:8080/api/community/get/${this.$route.params.id}`).then(res => {
            this.Community = res.data;

            let user = JSON.parse(localStorage.getItem("user")).user;

            if (user._id !== this.Community.user) {
                this.$router.push('/');
            }
        }).catch(err => {
            console.log(err);
        })
    },
    methods: {
        onClick() {
            var ls = localStorage.getItem('token');
            ls = JSON.parse(ls);

            var user = localStorage.getItem("user");
            user = JSON.parse(user);

            let config = {
                headers: {
                    Authorization: `Bearer ${ls.token}`
                }
            };

            axios.put('http://localhost:8080/api/community/update', {
                _id: this.Community._id,
                title: this.Community.title,
                description: this.Community.description,
                user: user._id
            }, config).then(res => {
                this.$router.push(`/Community/${this.$route.params.id}`);
            }).catch(err => {
                console.log(err);
                this.error = "Ha ocurrido un error, intentalo después :("
            })
        }
    }
}
</script>
