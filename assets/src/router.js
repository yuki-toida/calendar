import Vue from 'vue';
import Router from 'vue-router';
import Calendar from './components/calendar.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Calendar',
      component: Calendar,
    },
  ],
});
