import { createRouter, createWebHistory } from "vue-router";
import Home from "./pages/Home.vue";
import Post from "./pages/Post.vue";
import Config from "./pages/Config.vue";
import User from "./pages/User.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/post", component: Post },
  { path: "/config", component: Config },
  { path: "/user", component: User },
]
const router = createRouter({
    history: createWebHistory(),
    routes,
  });
export default router;