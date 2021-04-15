import { createStore } from "vuex";
import axios from 'axios'

import router from '../router'

const store = createStore({
state:{
  title:"Vuex Store",
  users: [],
  login:false,
  user:{username:"",password:"",phone:""}
},
getters:{
  users(state){
    return state.users
  },
  user(state){
    return state.user
  },
  login(state){
    return state.login
  }
},
mutations:{
  SET_USERS(state,users){
    state.users = users
  },
  SET_ONE_USER(state,user){
    state.user = user
  },
  SET_LOGIN_USER(state,login){
    state.login =login
  }
},
actions:{
  setOneUser({commit},user){
    
    user = user.value
    console.log(user.username,user.password,user.phone)
    axios.post("/api/user/new",user,)
    .then(res => {
      if(res.data.login){
        commit("SET_LOGIN_USER",res.data.login)
        localStorage.setItem('token', res.data.token);
        router.push("/cp")
      }
    

    })
  },
  getUsers({commit}){
    axios.get("/api/users")
    .then(res =>{
      console.log(res.data)
      commit("SET_USERS",res.data)})},
      
  getOneUser({commit},id){
    axios.get(`/api/user/`+ id)
    .then(res => {
      console.log(res.data)
      commit("SET_ONE_USER",res.data)
    })
  },
  logout({commit}){
    commit("SET_LOGIN_USER",false)
    localStorage.removeItem("token")
    router.push("/auth/login")
  }
},

})

export default store 