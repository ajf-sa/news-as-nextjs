  
import { createRouter, createWebHistory } from "vue-router";
import Home from "./pages/Home.vue";
import Local from "./pages/Local.vue";
import Sport from "./pages/Sport.vue";
import Tech from "./pages/Tech.vue";
import About from "./pages/About.vue";
const routes = [
  { path: "/", component: Home },
  { path: "/local", component: Local },
  { path: "/sport", component: Sport },
  { path: "/tech", component: Tech },
  { path: "/about", component: About },
]
const router = createRouter({
    history: createWebHistory(),
    routes,
  });
export default router;