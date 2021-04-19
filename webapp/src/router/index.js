import { createRouter, createWebHistory } from "vue-router"
import store from "../store"
import Admin from "../layouts/Admin.vue"
import Auth from "../layouts/Auth.vue"


import Dashboard from '../pages/admin/Dashboard.vue'
import User from "../pages/admin/User.vue"
import Post from "../pages/admin/Post.vue"
import Image from "../pages/admin/Image.vue"

import Login from '../pages/auth/Login.vue'
import Register from '../pages/auth/Register.vue'

import Home from "../pages/Home.vue"
import PageNotFound from "../pages/PageNotFound.vue"


const routes = [
    {
    redirect: "/cp/dashboard",
    path:"/cp", component:Admin,
    beforeEnter: [(to,from ,next) =>{
      console.log(to.path)
      if (!store.getters.login){
       return next({
         name:"login",
         params: {next:to.path}
       })
      }
      next()
    }],
    children:[
      { path: "/cp/dashboard", component: Dashboard,name:"cp-dashboard" },
      { path: "/cp/post", component: Post,name:"cp-post" },
      { path: "/cp/image", component: Image,name:"cp-image" },
      { path: "/cp/users", component: User,name:"cp-users" },
      { path: "user/:id", component: User ,name:"cp-user"},
    ]
},
{
  redirect:"/auth/login",
  path:"/auth", component: Auth,
  children:[
    {path:"/auth/login",component:Login,name:'login-c'},
    {path:"/auth/login/:next",component:Login,name:'login'},
    {path:"/auth/register",component:Register,name:'register'},
  ]

},
{ path: "/", component: Home,name:"index" },
  
  { path: '/:pathMatch(.*)*', component: PageNotFound }
]
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router