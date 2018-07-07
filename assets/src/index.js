import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';
import Vue from 'vue';
import Vuex from 'vuex';
import firebase from 'firebase';
import router from './router';
import app from './app.vue';
import http from './http';

Vue.use(Vuex);

// Initialize Firebase
firebase.initializeApp({
  apiKey: 'AIzaSyCMtszkhNgnTODhCKTw9cz5hDVPaOdkv68',
  authDomain: 'planet-pluto-dev.firebaseapp.com',
  databaseURL: 'https://planet-pluto-dev.firebaseio.com',
  projectId: 'planet-pluto-dev',
  storageBucket: '',
  messagingSenderId: '631172645333',
});

firebase.auth().onAuthStateChanged((user) => {
  const googleUser = user || null;
  const email = googleUser ? googleUser.email : null;
  const name = googleUser ? googleUser.displayName : null;
  const photo = googleUser ? googleUser.photoURL : null;
  http.post('/init', { email, name, photo }).then((data) => {
    console.log(data);
    const store = new Vuex.Store({
      strict: false,
      state: {
        user: data.user,
        staticUrl: data.staticUrl,
        events: data.events,
      },
      mutations: {
        addEvent(state, payload) {
          state.events.push(payload);
        },
        removeEvent(state, payload) {
          const ev = state.events.find(x => x.id === payload.id);
          state.events.splice(state.events.indexOf(ev), 1);
        },
      },
    });
    new Vue({
      el: '#app',
      router,
      components: { app },
      render: h => h(app),
      store,
    });
  });
});
