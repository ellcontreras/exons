<template>
    <div>
        <h1 class="title is-1 has-text-centered">
            Actualizando una comunidad
        </h1>
        <div class="notification is-danger" v-if="error">
            {{ error }}            
        </div>
        <div class="field">
            <input type="text" v-model="exonsunity.title" class="input" placeholder="title">
        </div>
        <div class="fiel">
            <textarea v-model="exonsunity.description" class="textarea" placeholder="description"></textarea>
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
        onClick() {
            axios.put('http://localhost:8080/api/exonsunity/update', {
                _id: this.exonsunity._id,
                title: this.exonsunity.title,
                description: this.exonsunity.description
            }).then(res => {
                this.$router.push(`/exonsunity/${this.$route.params.id}`);
            }).catch(err => {
                console.log(err);
                this.error = "Ha ocurrido un error, intentalo después :("
            })
        }
    }
}
</script>
