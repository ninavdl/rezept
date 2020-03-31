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
          <UserMenu v-on:showLogin="showLogin = true" />
        </template>
      </b-navbar>
    </header>
    <router-view v-on:loading="loading" class="section container"></router-view>

    <b-modal v-bind:active.sync="showLogin" aria-role="dialog" aria-modal>
      <LoginMenu />
    </b-modal>
  </div>
</template>

<script>
import Vue from "vue";

import UserMenu from "./UserMenu.vue";
import LoginMenu from "./LoginMenu.vue";

import { Navbar, Button, Modal } from "buefy";

Vue.use(Navbar);
Vue.use(Button);
Vue.use(Modal);

export default Vue.extend({
  methods: {
    loading: function(show) {
      this.isLoading = show;
    }
  },
  computed: {
    isLoggedIn() {
      return this.$store.state.isLoggedIn;
    }
  },
  components: {
    LoginMenu,
    UserMenu
  },
  data: function() {
    return {
      showLogin: false,
      isLoading: false,
      childComponent: this.$controller.mainComponent
    };
  }
});
</script>
