<template>
  <div class="container">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <router-link class="navbar-brand" v-bind:to="{ name: 'Calendar' }">
        <img v-bind:src="$store.state.staticPath + '/img/logo.jpg'" width="30" height="auto" class="d-inline-block align-top" alt="">
        Knowmeカレンダー
      </router-link>
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link class="nav-link" v-bind:to="{ name: 'User', params: { id: $store.state.user.id } }">
            検索
          </router-link>
        </li>
      </ul>
      <ul class="navbar-nav">
        <li v-if="$store.getters.isSignIn" class="nav-item">
          {{ $store.state.user.id }}
          <!--
          <button class="btn btn-danger" v-on:click="signOut">サインアウト</button>
          -->
        <li v-else class="nav-item">
          <button class="btn btn-primary" v-on:click="signIn">
            {{ $store.state.emailDomain }} でサインイン
          </button>
        </li>
      </ul>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script>
import firebase from 'firebase';
import http from './http';

export default {
  name: 'App',
  methods: {
    signIn: function() {
      var provider = new firebase.auth.GoogleAuthProvider();
      firebase.auth().signInWithPopup(provider).then((result) => {
        http.post('/signin', {
          email: result.user.email,
          name: result.user.displayName,
          photo: result.user.photoURL,
        })
        .then(() => window.location.reload())
        .catch(error => this.$toasted.show(error));
      }).catch((error) => console.log(error));
    },
    signOut: function() {
      firebase.auth().signOut().then(() => {
        http.delete('/signout')
        .then(() => window.location.reload())
        .catch(error => this.$toasted.show(error));
      }).catch((error) => console.log(error));
    }
  }  
}
</script>
