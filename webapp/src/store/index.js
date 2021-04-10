import { createStore } from "vuex";

const store = createStore({
state:{
  title:"Vuex Store",
  users: [
    {id:1,avatar:null,username:"north",password:"$2a$08$RVtfPN5ra1QslcX5GZcyFefvLJ./peWSvW0LjwFXoZsc7guw2gjem",phone:"0597417580",email:null,createAt:"2021-04-08T20:35:51.69Z",updateAt:"2021-04-08T20:35:51.69Z",isActive:true,isStaff:false,isAdmin:false},
    {id:8,avatar:null,username:"aziz",password:"$2a$08$Jo5zs6iEEpDSQQ11dF6XCunqMDTb76yx6r9ijAUItF3Py.3yPqav6",phone:"owerifjwoef",email:null,createAt:"2021-04-09T13:07:40.472Z",updateAt:"2021-04-09T13:07:40.472Z",isActive:false,isStaff:false,isAdmin:false}
  ],
},
getters:{},
mutations:{

},
actions:{}
})

export default store 