
<template>
  <ul id="example-1">
   <li v-for="user in $store.state.users" :key="user.id">
  
   <router-link :to="user.id">{{user.username}}</router-link>

  <h1></h1>
   </li>
</ul>
</template>

<script>
import {useStore} from "vuex"
import {reactive,computed} from "vue"
export default{
  setup(){
    const store = useStore()
    const state = reactive({
      test:"hello ",
    })

    return{
      state,
    }
  }
}
</script>


<style>

</style>