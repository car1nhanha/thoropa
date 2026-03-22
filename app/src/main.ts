import { createApp, h } from "vue";
import { createRouter, createWebHistory, RouterView } from "vue-router";
import Home from "./components/pages/home/Home.vue";
import Redirect from "./components/pages/redirect/redirect.vue";
import "./reset.css";
import "./style.css";

const routes = [
  { path: "/", component: Home },
  { path: "/redirect/:id", component: Redirect },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp({ render: () => h(RouterView) })
  .use(router)
  .mount("#app");
