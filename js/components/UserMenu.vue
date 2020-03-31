<template>
  <div v-if="isLoggedIn">
    <b-navbar-dropdown :label="'Hi ' + user.DisplayName" right boxed>
      <b-navbar-item v-on:click="logout" v-if="!isLoggingOut">
        Logout
      </b-navbar-item>
      <b-navbar-item v-else>
        Logging out…
      </b-navbar-item>
      <b-navbar-item
        tag="router-link"
        :to="{ name: 'list', query: { user: user.Username } }"
      >My recipes</b-navbar-item>
    </b-navbar-dropdown>
  </div>
  <b-navbar-item v-else>
    <b-button type="is-primary" v-on:click="$emit('showLogin')">Login</b-button>
  </b-navbar-item>
</template>

<script>
import Vue from "vue";
import Loading from "./Loading.vue";

import { Button, Tag } from "buefy";

Vue.use(Button);
Vue.use(Tag);

export default Vue.extend({
  methods: {
    logout: async function() {
      this.isLoggingOut = true;
      await this.$controller.logout();
      this.$store.commit("setUser", {});
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
  }),
  components: {
    Loading
  }
});
</script>