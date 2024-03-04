<template>
  <div id="app">
      <navbar />
      <b-container>
    <router-view/>
      </b-container>
  </div>
</template>

<script>
import Navbar from './views/layout/layout_Navbar.vue';

export default ({
  components: { Navbar },
  data() {
  },
});
</script>

<style lang="scss">

</style>
