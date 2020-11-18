import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./assets/compiled.css";

const app = createApp(App)
app.use(router)
app.mount("#app")


// Hot Module Replacement (HMR) - Remove this snippet to remove HMR.
// Learn more: https://www.snowpack.dev/#hot-module-replacement
if (import.meta.hot) {
  import.meta.hot.accept();
  import.meta.hot.dispose(() => {
    app.unmount();
  });
}