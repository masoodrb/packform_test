import { createApp } from "vue";
import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import { library } from '@fortawesome/fontawesome-svg-core'
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons'
import "element-plus/dist/index.css";

import App from "./App.vue";
import router from "./router";

const app = createApp(App);

library.add(faMagnifyingGlass)

app.use(ElementPlus);
app.use(createPinia());
app.use(router);

app.mount("#app");
