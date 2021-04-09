<template>
  <div v-if="items.length">
    <div class="rounded-md bg-blue-50 px-3 py-1 my-1 h-auto" v-for="item in items" :key="item.id">
      <div class="flex flex-wrap">
        <a class="flex-auto font-sans text-black" target="blank" :href="item.url">{{ item.title }}</a>
        <div class="text-sm font-medium text-gray-500">{{ item.time }}</div>
        <div class="w-full flex-nonefont-semibold text-yellow-500 mt-2">{{ item.by }}</div>
      </div>
    </div>
  </div>
  <div v-else>
    <svg class="text-red-500 animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
    </svg>
  </div>
</template>
<script>
export default {
  name: "Stories",
  props: {
    items: Array,
  }
</script>
<style>

</style>