<template>
  <nav class="navbar is-link">
    <div class="navbar-brand">
      <router-link to="/">
        <h1 class="title is-1 has-text-white">
          Comm
        </h1>
      </router-link>
    </div>
    <div class="navbar-menu">
      <div class="navbar-end">
        <router-link to="/add-community" class="has-text-white navbar-item">
          Agregar comunidad
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
