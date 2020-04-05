<template>
  <div class="container">
    <header role="banner">
      <b-navbar shadow spaced style="z-index: 10">
        <template slot="brand">
          <b-navbar-item tag="router-link" :to="{ name: 'list' }">
            <h1 class="title">rezept</h1>
          </b-navbar-item>
        </template>
        <template slot="end">
          <UserMenu v-on:showLogin="showLogin = true" v-on:showSignup="showSignup = true" />
        </template>
      </b-navbar>
    </header>
    <router-view v-on:loading="loading" class="section container"></router-view>

    <b-modal v-bind:active.sync="showLogin" aria-role="dialog" aria-modal>
      <LoginMenu />
    </b-modal>
    <b-modal v-bind:active.sync="showSignup" aria-role="dialog" aria-modal>
      <SignupMenu />
    </b-modal>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'vue-property-decorator';

import { Navbar, Button, Modal } from 'buefy';
import UserMenu from './UserMenu.vue';
import LoginMenu from './LoginMenu.vue';
import SignupMenu from './SignupMenu.vue';


Vue.use(Navbar);
Vue.use(Button);
Vue.use(Modal);

@Component({
  components: {
    LoginMenu,
    SignupMenu,
    UserMenu,
  },
})
export default class MainComponent extends Vue {
  showLogin = false;

  showSignup = false;

  isLoading = false;

  get isLoggedIn(): boolean {
    return this.$store.state.isLoggedIn;
  }

  loading(show): void {
    this.isLoading = show;
  }
}
</script>
