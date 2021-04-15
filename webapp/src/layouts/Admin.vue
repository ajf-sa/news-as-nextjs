<template>
admin page <button @click="logout" class="bg-blue-100 flex-grow px-4 py-2 text-xs font-semibold tracking-wider text-blue-600 rounded hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"> logout</button>

<h1>
<router-link class="more" :to="{ name: 'cp-dashboard',}"  >dashboard </router-link>
 </h1>
 <h1>
<router-link class="more" :to="{ name: 'cp-post',}"  >post </router-link>
 </h1>
 <h1>
<router-link class="more" :to="{ name: 'cp-users',}"  >users </router-link>
 </h1>
 <router-view />
</template>

<script>
import {useStore} from "vuex"
import { ref ,computed} from "vue"

export default{
    setup(){
        const store = useStore()
        // const user =computed(() => store.getters.user)

        function logout(){
            store.dispatch("logout");
        }
        return {
            logout
        }
    }

}
</script>
<style></style>



