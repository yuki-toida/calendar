import Vue from 'vue';
import Router from 'vue-router';
import calendar from './components/calendar.vue';
import user from './components/user.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Calendar',
      component: calendar,
    },
    {
      path: '/user',
      name: 'User',
      component: user,
    },
  ],
});
