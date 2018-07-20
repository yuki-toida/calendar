<template>
  <div class="container">
    <nav class="navbar navbar-expand-lg navbar-light">
      <router-link class="navbar-brand" v-bind:to="{ name: 'Calendar' }">
        <img v-bind:src="'/static/img/logo.jpg'" width="30" height="auto" class="d-inline-block align-top" alt="">
        Knowmeカレンダー<span class="text-muted">（β版)</span>
      </router-link>
      <ul class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link class="nav-link" v-bind:to="{ name: 'Images' }">
            飲み画像
          </router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link" v-if="$store.getters.isSignIn" v-bind:to="{ name: 'Search', params: { id: $store.state.user.id } }">
            飲み履歴
          </router-link>
        </li>
      </ul>
      <ul class="navbar-nav">
        <li v-if="$store.getters.isSignIn" class="nav-item">
          {{ $store.state.user.id }}
        <li v-else class="nav-item">
          <button class="btn btn-outline-secondary" v-on:click="signIn">
            {{ $store.state.domain }} でサインイン
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
        .then((data) => this.$store.commit('setUser', data.user))
        .catch(error => this.$toasted.show(error));
      })
      .catch((error) => console.log(error));
    },
    // signOut: function() {
    //   firebase.auth().signOut().then(() => {
    //     http.delete('/signout')
    //     .then(() => this.store.commit('setUser', null))
    //     .catch(error => this.$toasted.show(error));
    //   })
    //   .catch((error) => console.log(error));
    // }
  }  
}
</script>

<style>
.table-borderless td,
.table-borderless th {
  border: 0;
}
.bg-danger,
.badge-danger {
  background-color: #e72e5a !important;
}
.bg-success,
.badge-success {
  background-color: #3fbbb9 !important;
}
</style>
