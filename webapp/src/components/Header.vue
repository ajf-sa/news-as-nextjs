<template>
<nav class="bg-gray-800">
  <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8 right-0">
    <div class="relative flex items-center justify-between h-16">
      <div class="absolute inset-y-0 right-0 flex items-center sm:hidden">
        <!-- Mobile menu button-->
        <button @click="show" class="inline-flex items-right justify-right p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:bg-gray-700 focus:text-white transition duration-150 ease-in-out" aria-label="Main menu" aria-expanded="false">
          <!-- Icon when menu is closed. -->
          <!--
            Heroicon name: menu

            Menu open: "hidden", Menu closed: "block"
          -->
          <svg  class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <!-- Icon when menu is open. -->
          <!--
            Heroicon name: x

            Menu open: "block", Menu closed: "hidden"
          -->
          <svg class="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      <div class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
        <div class="flex-shrink-0">
          <h3  class="block lg:hidden h-8 w-auto px-6 font-medium leading-5 text-white py-2">موجز اخباري</h3>
          <h1  class="hidden lg:block h-8 w-auto px-6 font-medium leading-5 text-white py-2">موجز اخباري</h1>
        </div>
        <div class="hidden sm:block sm:ml-6">
          <div class="flex">
            <router-link to="/" class="ml-4 px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">عناوين</router-link>
            <router-link to="/local" class="ml-4 px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">محليات</router-link>
            <router-link to="/sport" class="ml-4 px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">رياضة</router-link>
            <router-link to="/about" class="ml-4 px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">تقنية</router-link>
          </div>
        </div>
      </div>

    </div>
  </div>

  <!--
    Mobile menu, toggle classes based on menu state.

    Menu open: "block", Menu closed: "hidden"
  -->
  <div :class="showMenu">
    <div class="px-2 pt-2 pb-3">
      <router-link to="/" class="block px-3 py-2 rounded-md text-base font-medium text-white  hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">عناوين</router-link>
      <router-link to="/local" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">محليات</router-link>
      <router-link to="/sport" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">رياضة</router-link>
      <router-link to="/tech" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">تقنية</router-link>
    </div>
  </div>
</nav>
  <!-- <nav
    class="w-full bg-gradient-to-l text-blue px-4 py-5"
  >
    <router-link  class="ml-3 inline-flex rounded-md shadow inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-indigo-600 bg-white hover:text-indigo-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out" to="/">Home</router-link>
    <router-link
      v-for="item in list"
      :key="item.to"
      class="ml-3 inline-flex rounded-md shadow inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-indigo-600 bg-white hover:text-indigo-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out"
      :to="item.to"
      >{{ item.title }}</router-link>
  </nav> -->
</template>

<script>

export default {
  data() {
   
    return {
      list: [
        { title: "About", to: "/about" },
      ],
      showMenu:"hidden sm:hidden",
      isActive:false,
    };
  },
   methods: {
     show(){
       this.isActive = !this.isActive
      if (this.isActive){
        this.showMenu ="block sm:hidden"
        
      }
      else{
         this.showMenu ="hidden sm:hidden"
      }
     }
   },
 
};
</script>

<style></style>