
<template>
 <nav class="bg-white shadow dark:bg-gray-800">
        <div class="container px-6 py-3 mx-auto md:flex">
            <div class="flex items-center justify-between">
                <div>
                    <a class="text-xl font-bold text-gray-800 dark:text-white md:text-2xl hover:text-gray-700 dark:hover:text-gray-300" href="#">موجز</a>
                </div>
                
                <!-- Mobile menu button -->
                <div class="flex md:hidden">
                    <button  @click="toggile" type="button" class="text-gray-500 dark:text-gray-200 hover:text-gray-600 dark:hover:text-gray-400 focus:outline-none focus:text-gray-600 dark:focus:text-gray-400" aria-label="toggle menu">
                        <svg viewBox="0 0 24 24" class="w-6 h-6 fill-current">
                            <path fill-rule="evenodd" d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"></path>
                        </svg>
                    </button>
                </div>
            </div>

            <!-- Mobile Menu open: "block", Menu closed: "hidden" -->
            <div class="w-full md:flex md:items-center md:justify-between" v-bind:class="{hidden:isActive}">
                <div class="flex flex-col px-2  py-3 -mx-4 md:flex-row md:mx-0 md:py-0">
                    <a href="#" class="px-2 py-1 text-sm font-medium text-gray-700 transition-colors duration-200 transform rounded dark:text-gray-200 hover:bg-gray-900 hover:text-gray-100 md:mx-2">محليات</a>
                    <a href="#" class="px-2 py-1 text-sm font-medium text-gray-700 transition-colors duration-200 transform rounded dark:text-gray-200 hover:bg-gray-900 hover:text-gray-100 md:mx-2">رياضة</a>
                    <a href="#" class="px-2 py-1 text-sm font-medium text-gray-700 transition-colors duration-200 transform rounded dark:text-gray-200 hover:bg-gray-900 hover:text-gray-100 md:mx-2">تقنية</a>
                </div>
                
                 <div class="flex items-center py-2 -mx-1 md:mx-0">
                        <router-link  :to="{name:'login'}" class="block w-1/2 px-3 py-2 mx-auto text-sm font-medium leading-5 text-center text-gray-400 transition-colors duration-200 transform bg-white-500 rounded-md hover:text-white hover:bg-gray-600 md:mx-2 md:w-auto" >دخول</router-link>
                 </div>

            </div>
        </div>
    </nav>

<section class="text-gray-600 body-font text-right">
  <div class="container px-5 py-24 mx-auto">
    <div class="flex flex-wrap  flex-row -m-4" >
     <div v-for="item in 50" :key="item" class="p-4 md:w-1/3">
        <div class="h-full border-2 border-gray-200 border-opacity-60 rounded-lg overflow-hidden">
          <img class="lg:h-48 md:h-36 w-full object-cover object-center" v-lazy ='image' alt="blog">
          <div class="p-6">
            <h2 class="tracking-widest text-xs title-font font-medium text-gray-400 mb-1">التصنيف</h2>
            <h1 class="title-font text-lg font-medium text-gray-900 mb-3">عنوان الخبر</h1>
            <p class="leading-relaxed mb-3">نبذه قصير عن الخبر من عدد كلمات تشرح الخبر بشكل سريع ومبسط</p>
            <div class="flex items-center flex-rap">
             
                <a class="text-pink-500 ml-auto inline-flex items-center md:mb-2 lg:mb-0">
                   <svg class="w-4 h-4 ml-2" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M5 12h14"></path>
                  <path d="M12 5l7 7-7 7"></path>
                </svg>
                  قراءة المزيد
               
              </a>
            

            <span class="text-gray-400 inline-flex items-center  leading-none text-sm">
                <svg class="w-4 h-4 mr-1" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                  <circle cx="12" cy="12" r="3"></circle>
                </svg>1.2K
              </span>
              <span class="text-gray-400 inline-flex items-center  leading-none text-sm">
                <svg class="w-4 h-4 mr-1" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round" viewBox="0 0 24 24">
                  <path d="M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z"></path>
                </svg>6
              </span>
              
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</section>
</template>

<script>
import {useStore} from "vuex"
import {ref,onMounted,computed} from "vue"
export default{
  setup(){
    const store = useStore()
    const isActive = ref(true)
    const users =computed(() => store.getters.users)
    function toggile(){
      console.log("click")
      isActive.value = !isActive.value
    }
    onMounted(()=>{
      store.dispatch("loginByToken")
        // store.dispatch("getUsers");
        // store.commit("SET_USERS")
    })
    return{
      users,
      toggile,
      isActive,
      image :"https://static.dw.com/image/49972010_403.jpg"
    }
  }
}
</script>


<style>

</style>