<template>
    <div class="box hero is-info">
        <div class="hero-body">
            <h1 class="title is-1 has-text-centered">
                Registro
            </h1>
            <div class="notification is-danger" v-if="error_response">
                {{ error_response }}
            </div>
            <div class="field">
                <label class="label">Nombre</label>
                <input type="text" v-model="name" class="input">
            </div>
            <div class="field">
                <label class="label">Username</label>
                <input type="text" v-model="username" class="input">
            </div>
            <div class="field">
                <label class="label">Email</label>
                <input type="email" v-model="email" class="input">
            </div>
            <div class="field">
                <label class="label">Contraseña</label>
                <input type="password" v-model="password" class="input">
            </div>
            <div class="field">
                <label class="label">Repite tu Contraseña</label>
                <input @input="handleChangePassword()" type="password" v-bind:class="{'is-danger' : error_password}"
                    v-model="repeat_password" class="input">
            </div>
            <small><span class="has-text-danger">*</span> todos tus datos para mandar el formulario</small>
            <br>
            <button @click="handleClick()" v-if="handleAllFilled() && !error_password"
                class="button is-info is-inverted is-rounded">Registrarme</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            name: "",
            username: "",
            email: "",
            password: "",
            repeat_password: "",
            error_password: "",
            error_response: ""
        }
    },
    beforeMount() {
        if (localStorage.getItem("user")) {
            this.$router.push('/');
        }
    },
    methods: {    
        handleClick() {
            if (this.handleAllFilled) {
                axios.post("http://localhost:8080/api/user/add", {
                    name: this.name,
                    username: this.username,
                    email: this.email,
                    password: this.password
                }).then(res => {
                    this.$router.push('/login');
                }).catch(error => {
                    if (error.response.status === 409) {
                        this.error_response = "Ese email o username ya se encuentra registrado";
                    } else {
                        this.error_response = "Ha ocurrido un error con el servidor intentalo más tarde";
                    }
                });
            }
        },
        handleAllFilled() {
            return this.name && this.username && this.email && 
                this.password && this.repeat_password;
        },
        handleChangePassword() {
            if (this.password != this.repeat_password) {
                this.error_password = "Las contraseñas no coinciden";
            } else {
                this.error_password = "";
            }
        }
    },
}
</script>
