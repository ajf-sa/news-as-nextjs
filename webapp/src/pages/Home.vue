
<template>
  <ul id="example-1">
    <div v-if="users.length > 0 ">
   <li v-for="user in users" :key='user.id'>

  
   <router-link :to="{path:'/cp/user/'+user.id}" > <h1>{{user.username}}</h1></router-link>

   </li>
   </div>
<div v-else>
 No Records found
</div>
</ul>
</template>

<script>
import {useStore} from "vuex"
import {onMounted,computed} from "vue"
export default{
  setup(){
    const store = useStore()
    const users =computed(() => store.getters.users)
    onMounted(()=>{
        store.dispatch("getUsers");
        store.commit("SET_USERS")
    })
    
    return{
      users
    }
  }
}
</script>


<style>

</style>