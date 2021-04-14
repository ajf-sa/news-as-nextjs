<template>
admin page

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
<script></script>
<style></style>



