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
       <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2 text-right" for="phone">
         رقم الجوال
        </label>
        <input v-model="user.phone" class="shadow appearance-none border  rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" name="phone" id="phone" type="phone" placeholder="رقم الجوال">
        <p class="text-red-500 text-xs italic"></p>
      </div>
      <div class="flex h-8 justify-end">
     <span class="ml-2 px-2">موافق على <span class="underline">اتفاقية الاستخدام</span></span>
  <label class="flex items-center  ">   
    <input type="checkbox"  v-model="checked" class="form-checkbox">
  </label>
</div>
      <div class="flex items-center justify-between ">
         <button @click="submit" class="bg-blue-100 flex-grow h-10 px-4 py-2 text-xs font-semibold tracking-wider text-blue-600 rounded hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">تسجيل</button>
      </div>
    </form>

    
  </div>
</div>
</template>

<script>
import {useStore} from "vuex"
import { ref ,computed} from "vue"

export default{
    setup(){
        const name = ref("");
        const password = ref("");
        const phone = ref("");
        const checked = ref("");

        const store = useStore()
        const user =computed(() => store.getters.user)

        function submit(){
            store.dispatch("setOneUser",user);
        }
        return {
            name,
            user,
            password,
            phone,
            checked,
            submit
        }
    }

}
</script>
