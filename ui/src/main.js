import Vue from 'vue';
import VueRouter from 'vue-router';
import Resource from 'vue-resource';

import App from './App.vue';
import Index from './components/Index.vue';
import NewThread from './components/NewThread.vue';
import Thread from './components/Thread.vue';
import Registration from './components/Registration.vue';
import Login from './components/Login.vue';
import Profile from './components/Profile.vue';
import UserIndex from './components/UserIndex.vue';

Vue.use(VueRouter);
Vue.use(Resource);

Vue.config.productionTip = false;

const routes = [{
  path: '/',
  component: Index,
},
{
  path: '/thread/new',
  component: NewThread,
},
{
  path: '/thread/',
  component: Thread,
},
{
  path: '/thread/:thread',
  component: Thread,
},
{
  path: '/register',
  component: Registration,
},
{
  path: '/login',
  component: Login,
},
{
  path: '/profile/:user',
  component: Profile,
},
{
  path: '/users',
  component: UserIndex,
},
{
  path: '*',
  component: {
    template: '<div>Not Found</div>',
  },
},
];

const router = new VueRouter({ routes, mode: 'history' });


new Vue({
  router,
  render: h => h(App),
}).$mount('#app');
