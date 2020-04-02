<template>
  <div v-if="isLoggedIn">
    <b-navbar-dropdown :label="'Hi ' + user.DisplayName" right boxed>
      <b-navbar-item v-on:click="logout" v-if="!isLoggingOut">Logout</b-navbar-item>
      <b-navbar-item v-else>Logging out…</b-navbar-item>
      <b-navbar-item
        tag="router-link"
        :to="{ name: 'list', query: { user: user.Username } }"
      >My recipes</b-navbar-item>
    </b-navbar-dropdown>
  </div>
  <b-navbar-item tag="div" v-else>
    <b-navbar-item>
      <b-button type="is-primary" v-on:click="$emit('showLogin')">Login</b-button>
    </b-navbar-item>
    <b-navbar-item>
      <b-button type="is-primary" v-on:click="$emit('showSignup')">Sign up</b-button>
    </b-navbar-item>
  </b-navbar-item>
</template>

<script>
import Vue from "vue";

import { Button, Tag } from "buefy";

import User from "../models/User.ts";
import API from "../models/API.ts";
import Cookies from "cookies-js";

Vue.use(Button);
Vue.use(Tag);

export default Vue.extend({
  methods: {
    logout: async function() {
      this.isLoggingOut = true;
      await User.logout();
      this.$store.commit("setUser", {});
      API.getInstance().setToken(null);
      Cookies.expire("token");

      this.isLoggingOut = false;
    }
  },
  computed: {
    user: function() {
      return this.$store.state.user;
    },
    isLoggedIn() {
      return this.$store.state.isLoggedIn;
    }
  },
  data: () => ({
    isLoggingOut: false
  })
});
</script>