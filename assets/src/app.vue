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
      <div v-if="isSignIn">
        <img v-bind:src="$store.state.user.photo" width="30" height="auto">
        <span>{{ $store.state.user.email }}</span>
        <span>{{ $store.state.user.name }}</span>
        <button class="button is-danger" v-on:click="signOut">SignOut</button>
      </div>
      <div v-else>
        <button class="button is-primary" v-on:click="signIn">SignIn</button>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
import firebase from 'firebase';

export default {
  computed: {
    isSignIn() {
      return this.$store.state.user.id;
    }
  },
  methods: {
    signIn: function() {
      var provider = new firebase.auth.GoogleAuthProvider();
      firebase.auth().signInWithPopup(provider).then(function(result) {
        window.location.reload();
      }).catch(function(error) {
        console.log(error);
      });
    },
    signOut: function() {
      firebase.auth().signOut().then(function() {
        window.location.reload();
      }).catch(function(error) {
        console.log(error);
      });
    }
  }
}
</script>
