<template>
  <div>
    <div class="container">
      <div class="starter-template">
        <h1>
          <router-link v-bind:to="{ name: 'Calendar' }">Hoge</router-link>
        </h1>
        <p class="lead">
          Use this document as a way to quickly start any new project.
        </p>
      </div>
    </div>
    <div class="container">
      <div v-if="$store.getters.isSignIn">
        <img v-bind:src="$store.state.user.photo" width="30" height="auto">
        <span>{{ $store.state.user.email }}</span>
        <span>{{ $store.state.user.name }}</span>
        <button class="button is-danger" v-on:click="signOut">サインアウト</button>
      </div>
      <div v-else>
        <button class="button is-primary" v-on:click="signIn">
          {{ $store.state.emailDomain }} でサインイン
        </button>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import firebase from 'firebase';
import http from './http';

export default {
  methods: {
    signIn: function() {
      var provider = new firebase.auth.GoogleAuthProvider();
      firebase.auth().signInWithPopup(provider).then((result) => {
        http.post('/signin', {
          email: result.user.email,
          name: result.user.displayName,
          photo: result.user.photoURL,
        }).then((data) => {
          this.$store.commit('setUser', data.user);
        });
      }).catch((error) => {
        console.log(error);
      });
    },
    signOut: function() {
      firebase.auth().signOut().then(() => {
        http.delete('/signout').then((data) => {
          this.$store.commit('setUser', data.user);
        });
      }).catch((error) => {
        console.log(error);
      });
    }
  }
}
</script>
