import { createStore } from "vuex";
import axios from "axios";

const api = axios.create({
});

export default createStore({
  state: {
    items: [],
  },
  mutations: {
    loadItems(state, item) {
      state.items.push(item);
    },
  },
  actions: {
    loadLatestTopItems(context) {
    fetch("/api/users")
    .then(response => response.json())
    .then(data => {
    console.log(data)
    return context.commit("loadItems", data)
    })
    },
  },
});