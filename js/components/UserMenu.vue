<template>
  <div v-if="isLoggedIn">
    <b-navbar-item tag="div" class="columns">
      <div class="column">Hi {{ user.DisplayName }}</div>
      <div class="column">
        <Loading v-if="isLoggingOut" text="Logging out" inline></Loading>
        <b-button v-on:click="logout" v-else>Logout</b-button>
      </div>
    </b-navbar-item>
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