
<template>

  <h1>{{message}}</h1>
</template>

<script>


export default{
  
  data(){
    return {
      message:"Post"
    }
  },computed:{
     greeting() {
      return this.message + '!'
    }
  },methods:{
    morwelcome(){

    fetch("http://localhost:3000/api/users")
    .then(response => response.json())
    .then(data => (this.message = data.test))
     
    }
  }

}
</script>


<style>

</style>