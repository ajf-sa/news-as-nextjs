
<template>
{{morwelcome()}}
  <h1>{{message}}</h1>
</template>

<script>


export default{
  
  data(){
    return {
      message:"hello"
    }
  },computed:{
     greeting() {
      return this.message + '!'
    }
  },methods:{
    morwelcome(){

    fetch("/api/users")
    .then(response => response.json())
    .then(data => (this.message = data.test))
     
    }
  }

}
</script>


<style>

</style>