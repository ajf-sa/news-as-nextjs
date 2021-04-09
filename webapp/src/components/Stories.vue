
<template>
	<div v-if="items.length">
		<div class="rounded-md bg-gray-100 px-3 py-1 my-1 h-auto dark:bg-gray-800" v-for="item in items" :key="item.id">
			<h1>ddd {{item.username}}</h1>
		</div>
	</div>
	<div class="flex" v-else>
		<div class="flex-auto text-center">
			<span class="animate-ping text-red-500 dark:text-pink-200">Loading ...</span>
		</div>
	</div>
</template>
<script>
export default {
	name: "Stories",
	props: {
		items: Array,
	},

};
</script>