<template>
  <div class="App">
    <header class="App-header">
      <img src="/sa.svg" class="App-logo" alt="logo" />

      <a class="App-link" href="https://ajf.sa" target="_blank" rel="noopener noreferrer">{{
        state.message
      }}</a>
    </header>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive} from 'vue';
interface State {
  message: string;
}
export default defineComponent({
  components: {},
  setup() {
    const state = reactive({
      message: 'سارعي للمجد والعلياء',
    });
    return {
      state,
    };
  },
});
</script>

<style>
.App {
  text-align: center;
}
.App-header {
  background-color: #f9f6f6;
  color: #32485f;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 2vmin);
}
.App-link {
  color: #00c185;
}
.App-logo {
  height: 40vmin;
  pointer-events: none;
  margin-bottom: 1rem;
  animation: App-logo-spin infinite 1.6s ease-in-out alternate;
}
.App-tsx {
  display: flex;
}
.App-tsx > div {
  margin-left: 30px;
  font-size: 16px;
}
@keyframes App-logo-spin {
  from {
    transform: scale(1);
  }
  to {
    transform: scale(1.03);
  }
}
</style>