import Vue from 'vue';
import Router from 'vue-router';
import Calendar from './components/Calendar.vue';
import Search from './components/Search.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Calendar',
      component: Calendar,
    },
    {
      path: '/search/:id',
      name: 'Search',
      component: Search,
    },
  ],
});
