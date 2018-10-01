<template>
  <nav class="navbar is-info">
    <div class="navbar-brand">
      <router-link to="/">
        <h2 class="subtitle is-2 has-text-white">
          exons
        </h2>
      </router-link>
    </div>
    <div class="navbar-menu">
      <div class="navbar-end">
        <router-link to="/add-Community" class="has-text-white navbar-item">
          Agregar comunidad
        </router-link>
        <router-link to="/sign-up" class="has-text-white navbar-item">
          Registrarme
        </router-link>
        <router-link to="/login" class="has-text-white navbar-item">
          Iniciar Sesión
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
nav {
  margin-bottom: 3em;
}
</style>
