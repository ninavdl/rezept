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

<script lang="ts">
import "reflect-metadata";
import { Component } from "vue-property-decorator";
import Vue from "vue";

import User from "../models/User";
import API from "../models/API";

import Cookies from "cookies-js";

import { Button, Tag } from "buefy";

Vue.use(Button);
Vue.use(Tag);

@Component({})
export default class UserMenuComponent extends Vue {
  isLoggingOut: boolean = false;

  get user(): User {
    return this.$store.state.user;
  }

  get isLoggedIn(): boolean {
    return this.$store.state.isLoggedIn;
  }

  async logout(): Promise<void> {
    this.isLoggingOut = true;
    await User.logout();
    this.$store.commit("setUser", {});
    API.getInstance().setToken(null);
    Cookies.expire("token");
    this.isLoggingOut = false;
  }
}
</script>