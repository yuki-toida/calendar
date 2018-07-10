import Vue from 'vue';
import Router from 'vue-router';
import Calendar from './components/Calendar.vue';
import User from './components/User.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Calendar',
      component: Calendar,
    },
    {
      path: '/users/:id',
      name: 'User',
      component: User,
    },
  ],
});
