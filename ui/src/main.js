import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

Vue.config.productionTip = false

import Index from './components/Index.vue';
import NewThread from './components/NewThread.vue';
import Thread from './components/Thread.vue';
import Registration from './components/Registration.vue';
import Login from './components/Login.vue';
import Profile from './components/Profile.vue';
import UserIndex from './components/UserIndex.vue';

const router = new VueRouter(
  [
    { path: '/', component: Index },
    { path: '/thread/new', component: NewThread },
    { path: '/thread/', component: Thread },
    { path: '/thread/:thread', component: Thread },
    { path: '/register', component: Registration },
    { path: '/login', component: Login },
    { path: '/profile', component: Profile },
    { path: '/users', component: UserIndex },
    { path: '*', component: { template: '<div>Not Found</div>' } },
  ]
);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
