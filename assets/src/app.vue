<template>
  <div class="container">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <a class="navbar-brand" href="#">
        <img v-bind:src="$store.state.staticPath + '/img/logo.jpg'" width="30" height="auto" class="d-inline-block align-top" alt="">
        Knowme
      </a>
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link class="nav-link" v-bind:to="{ name: 'Calendar' }">
            カレンダー
          </router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link" v-bind:to="{ name: 'User' }">
            ユーザー
          </router-link>
        </li>
      </ul>
      <ul class="navbar-nav">
        <li v-if="$store.getters.isSignIn" class="nav-item">
          {{ $store.state.user.id }}
          <!--
          <img v-bind:src="$store.state.user.photo" width="30" height="auto">
          {{ $store.state.user.name }}
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
  methods: {
    signIn: function() {
      var provider = new firebase.auth.GoogleAuthProvider();
      firebase.auth().signInWithPopup(provider).then((result) => {
        http.post('/signin', {
          email: result.user.email,
          name: result.user.displayName,
          photo: result.user.photoURL,
        }).then(() => {
          window.location.reload();
        });
      }).catch((error) => {
        console.log(error);
      });
    },
    signOut: function() {
      firebase.auth().signOut().then(() => {
        http.delete('/signout').then(() => {
          window.location.reload();
        });
      }).catch((error) => {
        console.log(error);
      });
    }
  }  
}
</script>
