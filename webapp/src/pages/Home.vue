
<template>
<Stories :items="items"></Stories>
</template>

<script>

import { mapState } from "vuex";
import Stories from "../components/Stories.vue";

export default{
  
  name: "Home",
  components: { Stories },
  data(){
    return {
      message:"hello"
    }
  },computed:{
     ...mapState(["items"]),
  },methods:{
  },
   created() {
    this.$store.dispatch("loadLatestTopItems");
  },

}
</script>


<style>

</style>