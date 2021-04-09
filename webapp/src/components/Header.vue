<template>
 <div class="px-2 pt-2 pb-3">
      <router-link to="/" class="block px-3 py-2 rounded-md text-base font-medium text-white  hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">عناوين</router-link>
      <router-link to="/user" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">محليات</router-link>
      <router-link to="/config" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">رياضة</router-link>
      <router-link to="/post" class="mt-1 block px-3 py-2 rounded-md text-base font-medium text-gray-300 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out">تقنية</router-link>
    </div>
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
 
}
</script>

<style></style>