import { createApp } from 'vue'


import App from './App.vue'
import store from "./store";
import "./assets/css/main.css";
import "./axios"
import VueLazyLoad from 'vue3-lazyload'
import router from './router'


const app = createApp(App)
// app.config.productionTip = true;
app.config.devtools = true;
app.use(router)
app.use(VueLazyLoad)
app.use(store)
app.mount("#app")


// if (import.meta.hot) {
//     import.meta.hot.accept();
//     import.meta.hot.dispose(() => {
//       app.unmount();
//     });
//   }
