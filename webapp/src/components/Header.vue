<template>
  <nav
    class="w-full bg-gradient-to-l text-blue px-4 py-5"
  >
    <router-link class="ml-3 inline-flex rounded-md shadow inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-indigo-600 bg-white hover:text-indigo-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out" to="/">Home</router-link>
    <router-link
      v-for="item in list"
      :key="item.to"
      class="ml-3 inline-flex rounded-md shadow inline-flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-indigo-600 bg-white hover:text-indigo-500 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out"
      :to="item.to"
      >{{ item.title }}</router-link>
  </nav>
</template>

<script>

export default {
  data() {
    return {
      list: [
        { title: "About", to: "/about" },
      ],
    };
  },
 
};
</script>

<style></style>