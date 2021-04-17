<template>
<div class="flex justify-center min-h-screen bg-gray-200">
  <div class="w-full max-w-xs m-auto ">
    <h1 class="text-center text-red-500 italic"></h1>
    <form  class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" v-on:submit.prevent >
      <input type="hidden" name="next" value="/cp">
      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2 text-right" for="username">
          اسم المستخدم
        </label>
        <input v-model="user.username" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" name="username" id="username" type="text" placeholder="اسم المستخدم">
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2 text-right" for="password">
          كلمة المرور
        </label>
        <input v-model="user.password" class="shadow appearance-none border  rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" name="password" id="password" type="password" placeholder="كلمة المرور">
        <p class="text-red-500 text-xs italic"></p>
      </div>
      <div class="flex items-center justify-between ">
         <button @click="submit" class="bg-blue-100 flex-grow h-10 px-4 py-2 text-xs font-semibold tracking-wider text-blue-600 rounded hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">دخول</button>
      </div>
    </form>
    <p class="text-center text-gray-500 text-xs">
      <router-link class="more" :to="{ name: 'register',}"  >تسجيل </router-link>
    </p>
  </div>
</div>
</template>

<script>
import {useStore} from "vuex"
import {onMounted, ref ,computed} from "vue"
export default{
    setup(){
        const name = ref("");
        const password = ref("")
        const store = useStore()
        const user =computed(() => store.getters.user)
        onMounted(()=>{
        store.dispatch("loginByToken")
    })
        function submit(){
           store.dispatch("setLoginUser",user)
        }
        return {
            user,
            name,
            password,
            submit
        }
    }

}
</script>
