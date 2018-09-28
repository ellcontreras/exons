<template>
  <div id="app">
    <Navbar/>
    <div class="container">
      <router-view/>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'

export default {
  components: {Navbar}
}
</script>


<style>
#app {
  color: #2c3e50 !important;
}

textarea { 
  resize: none !important;
}

</style>
