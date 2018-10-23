<template>
  <nav class="navbar is-info">
    <div class="navbar-brand">
      <router-link to="/">
        <h2 class="subtitle is-2 has-text-white">
          exons
        </h2>
      </router-link>
    </div>
    <div class="navbar-menu" v-if="!user">
      <div class="navbar-end">
        <router-link to="/sign-up" class="has-text-white navbar-item">
          Registrarme
        </router-link>
        <router-link to="/login" class="has-text-white navbar-item">
          Iniciar Sesión
        </router-link>
      </div>
    </div>
    <div class="navbar-menu" v-if="user">
      <div class="navbar-end">
        <router-link to="#" class="has-text-white navbar-item">
          Hola, {{ user.name }}!
        </router-link>
        <router-link to="/add-Community" class="has-text-white navbar-item">
          Agregar comunidad
        </router-link>
        <span @click="handleLogout()" class="has-text-white navbar-item">
          Cerrar Sesión
        </span>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  props: {
    msg: String
  },
  data() {
    return {
      user: null
    }
  },
  beforeMount() {
    if (localStorage.getItem("user")) {
      this.user = JSON.parse(localStorage.getItem("user")).user;
    }
  },
  methods: {
    handleLogout() {
      localStorage.removeItem("user");

      window.location.href = '/';
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
nav {
  margin-bottom: 3em;
}
</style>
