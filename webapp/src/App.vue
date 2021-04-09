<template>
<div>
  <Header />
   <div class="w-full flex">
    <router-view></router-view>
  </div>
</div>
</template>

<script>
import Header from "./components/Header.vue";

export default {
  name: "app",
  components: {Header},
  data() {
    return {};
  },

 
}
</script>

<style>

html{
  direction: rtl;
}

</style>