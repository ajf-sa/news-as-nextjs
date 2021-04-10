import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Post from "../pages/Post.vue";
import Config from "../pages/Config.vue";
import User from "../pages/User.vue";

const routes = [
  { path: "/cp", component: Home },
  { path: "/cp/post", component: Post },
  { path: "/cp/config", component: Config },
  { path: "/cp/user", component: User },
  { path: "/cp/user/:id", component: User },
]
const router = createRouter({
    history: createWebHistory(),
    routes,
  });
export default router;