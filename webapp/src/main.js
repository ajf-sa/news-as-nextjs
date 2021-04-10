import { createApp } from 'vue'
import App from './App.vue'
import store from "./store";
import router from "./router"
import "./assets/css/main.css";

const app = createApp(App)
app.config.productionTip = true;
app.config.devtools = true;
app.use(store)
app.use(router)
app.mount("#app")


if (import.meta.hot) {
    import.meta.hot.accept();
    import.meta.hot.dispose(() => {
      app.unmount();
    });
  }
