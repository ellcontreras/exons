<template>
    <div class="box hero is-info">
        <div class="hero-body">
            <h1 class="title is-1 has-text-centered">
                Login
            </h1>
            <div class="notification is-danger" v-if="error_response">
                {{ error_response }}
            </div>
            <div class="field">
                <label class="label">Email</label>
                <input type="email" v-model="email" class="input">
            </div>
            <div class="field">
                <label class="label">Contraseña</label>
                <input type="password" v-model="password" class="input">
            </div>
            <small><span class="has-text-danger">*</span> todos tus datos para mandar el formulario</small>
            <br>
            <button @click="handleClick()" v-if="handleAllFilled()"
                class="button is-info is-inverted is-rounded">Iniciar Sesión</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            email: "",
            password: "",
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
                axios.post("http://localhost:8080/api/user/login", {
                    email: this.email,
                    password: this.password
                }).then(res => {
                    localStorage.setItem("user", JSON.stringify({
                        token: res.data.token
                    }));

                    this.$router.push('/');

                    this.$root.$emit("send", "hola");

                }).catch(error => {
                    if (error.response.status === 401) {
                        this.error_response = "Los datos ingresados no son correctos, revisalos por favor";
                    } else {
                        this.error_response = "Ha ocurrido un error con el servidor intentalo más tarde";
                    }
                });
            }
        },
        handleAllFilled() {
            return this.email && this.password;
        }
    },
}
</script>
