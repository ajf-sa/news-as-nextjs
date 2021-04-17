<template>

  <div class="flex md:flex-row-reverse flex-wrap">
    <!--Main Content-->
    <div class="w-full md:w-4/5 bg-gray-100">
      <div class="container bg-gray-100 pt-16 px-6">
        admin page {{ username }}
        <router-view />
         <Switch
     v-model="enabled"
        :class='enabled ? "bg-blue-600" : "bg-gray-200"'
        class="relative inline-flex items-center h-6 transition-colors rounded-full w-11 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
 
  >
    <span class="sr-only">Enable notifications</span>
    <span
      :class="enabled ? 'translate-x-6' : 'translate-x-1'"
      class="inline-block w-4 h-4 transform bg-white rounded-full"
    />
  </Switch>
      </div>
    </div>

    <!--Sidebar-->
    <div
      class="w-full md:w-1/5 bg-gray-900 md:bg-gray-900 px-2 text-center fixed bottom-0 md:pt-8 md:top-0 md:right-0 h-16 md:h-screen md:border-r-4 md:border-gray-600"
    >
      <div class="md:relative mx-auto lg:float-right lg:px-6">
        <ul
          class="list-reset flex flex-row md:flex-col text-center md:text-right"
        >
          <li class="mr-3 flex-1">
            <router-link class="more" :to="{ name: 'cp-dashboard' }">
              <a
                class="block py-1 md:py-3 pl-1 align-middle text-gray-800 no-underline hover:text-pink-500 border-b-2 border-gray-800 md:border-gray-900 hover:border-pink-500"
              >
                <i class="fas fa-link pr-0 md:pr-3"></i
                ><span
                  class="pb-1 md:pb-0 text-xs md:text-base text-white md:font-bold block md:inline-block"
                >
                  لوحة التحكم
                </span>
              </a>
            </router-link>
          </li>
          <li class="mr-3 flex-1">
            <router-link class="more" :to="{ name: 'cp-post' }">
              <a
                class="block py-1 md:py-3 pl-1 align-middle text-gray-800 no-underline hover:text-pink-500 border-b-2 border-gray-800 md:border-gray-900 hover:border-pink-500"
              >
                <i class="fas fa-link pr-0 md:pr-3"></i
                ><span
                  class="pb-1 md:pb-0 text-xs md:text-base text-white md:font-bold block md:inline-block"
                >
                  الاخبار
                </span>
              </a>
            </router-link>
          </li>
          <li class="mr-3 flex-1">
            <router-link class="more" :to="{ name: 'cp-users' }">
              <a
                class="block py-1 md:py-3 pl-1 align-middle text-gray-800 no-underline hover:text-pink-500 border-b-2 border-gray-800 md:border-gray-900 hover:border-pink-500"
              >
                <i class="fas fa-link pr-0 md:pr-3"></i
                ><span
                  class="pb-1 md:pb-0 text-xs md:text-base text-white md:font-bold block md:inline-block"
                >
                  المستخدمين
                </span>
              </a>
            </router-link>
          </li>
          
          <li class="mr-3 flex-1">
            <a
              class="block py-1 md:py-3 pl-1 align-middle text-gray-800 no-underline hover:text-pink-500 border-b-2 border-gray-800 md:border-gray-900 hover:border-pink-500"
            >
              <i class="fas fa-link pr-0 md:pr-3"></i
              ><span
                class="pb-1 md:pb-0 text-xs md:text-base text-gray-600 md:text-gray-400 block md:inline-block"
              >
                <button
                  @click="logout"
                  class="bg-blue-100 flex-grow px-4 py-2 text-xs font-semibold tracking-wider text-blue-600 rounded hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                >
                  خروج
                </button>
              </span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { useStore } from "vuex";
import { ref, computed } from "vue";

 import { Switch } from "@headlessui/vue";

export default {
   components: { Switch },
  setup() {
    const store = useStore();
    const username = computed(() => store.getters.username);
     const enabled = ref(false);
    function logout() {
      store.dispatch("logout");
    }
    return {
      logout,
      username,
      enabled
    };
  },
};
</script>
<style></style>



