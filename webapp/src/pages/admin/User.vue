
<template>
  <div v-if="user">
    <h1>
    {{user.username}}
    </h1>
    </div>
</template>

<script>
import {useStore} from "vuex"
import {onMounted,computed} from "vue"
import {useRoute} from "vue-router";
export default{
  setup(){
    const store = useStore()
    const route = useRoute()
    const user =computed(() => store.getters.user)

    onMounted(()=>{
        // store.dispatch("getOneUser",route.params.id)
        // store.commit("SET_ONE_USER")
    })
    
    return{
      user
    }
  }
}
</script>


<style>

</style>