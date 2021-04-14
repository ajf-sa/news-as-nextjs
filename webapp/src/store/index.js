import { createStore } from "vuex";
import axios from 'axios'

const store = createStore({
state:{
  title:"Vuex Store",
  users: [],
  login:false,
  user:null
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
  }
},
actions:{
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
  }
}
})

export default store 