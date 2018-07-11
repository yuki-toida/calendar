import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';
import Vue from 'vue';
import Vuex from 'vuex';
import Toasted from 'vue-toasted';
import firebase from 'firebase';
import router from './router';
import App from './App.vue';
import http from './http';

Vue.use(Vuex);
Vue.use(Toasted, {
  theme: 'outline',
  position: 'top-center',
  duration: 4000,
  type: 'error',
  singleton: true,
});

function createStore(data) {
  return new Vuex.Store({
    strict: false,
    state: {
      emailDomain: data.emailDomain,
      events: data.events,
      user: data.user,
      myEvent: data.myEvent,
      dayEventRest: data.dayEventRest,
      nightEventRest: data.nightEventRest,
    },
    getters: {
      isSignIn: state => state.user != null,
    },
    mutations: {
      addEvent(state, payload) {
        state.events.push(payload);
      },
      removeEvent(state, payload) {
        const event = state.events.find(x => x.id === payload.id);
        state.events.splice(state.events.indexOf(event), 1);
      },
    },
  });
}

// Initialize Firebase
firebase.initializeApp({
  apiKey: 'AIzaSyCMtszkhNgnTODhCKTw9cz5hDVPaOdkv68',
  authDomain: 'planet-pluto-dev.firebaseapp.com',
  databaseURL: 'https://planet-pluto-dev.firebaseio.com',
  projectId: 'planet-pluto-dev',
  storageBucket: '',
  messagingSenderId: '631172645333',
});

http.get('/initial').then((data) => {
  console.log(data);
  new Vue({
    el: '#app',
    router,
    components: { App },
    render: h => h(App),
    store: createStore(data),
  });
}).catch(error => this.$toasted.show(error));

// firebase.auth().onAuthStateChanged((user) => {
//   const googleUser = user || null;
//   const email = googleUser ? googleUser.email : null;
//   const name = googleUser ? googleUser.displayName : null;
//   const photo = googleUser ? googleUser.photoURL : null;
// });
