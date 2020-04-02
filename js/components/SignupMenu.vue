<template>
  <form v-on:submit="login">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <h2 class="modal-card-title">Signup</h2>
        <p v-if="message != ''">{{ message }}</p>
      </header>
      <section class="modal-card-body" v-if="!isLoading">
        <template v-if="!isLoading">
          <b-field label="Display name">
            <b-input type="password" v-model="registration.DisplayName" required />
          </b-field>
          <b-field label="Username">
            <b-input type="string" v-model="registration.Username" required />
          </b-field>
          <b-field label="Password" message="Minimum 8 characters">
            <b-input type="password" v-model="registration.Password" required pattern=".{8,}" />
          </b-field>
        </template>
        <Loading v-else text="Logging in"></Loading>
      </section>
      <footer class="modal-card-foot">
        <b-button tag="input" native-type="submit" value="Sign up" />
      </footer>
    </div>
  </form>
</template>

<script>
import Vue from "vue";
import Loading from "./Loading.vue";
import UserRegistration from "../models/UserRegistration.ts";

import { Modal, Button, Field, Input } from "buefy";
import LoginRequest from "../models/LoginRequest";
import API from "../models/API.ts";
import User from "../models/User.ts";
import Cookies from "cookies-js";

Vue.use(Modal);
Vue.use(Button);
Vue.use(Field);
Vue.use(Input);

export default Vue.extend({
  data: () => ({
    isLoading: false,
    message: "",
    registration: new UserRegistration()
  }),
  components: {
    Loading
  },
  methods: {
    login: async function(ev) {
      ev.preventDefault();
      this.isLoading = true;
      try {
        await this.registration.signup();
        const loginRequest = new LoginRequest();
        loginRequest.Username = this.registration.Username;
        loginRequest.Password = this.registration.Password;
        const sessionId = await loginRequest.login();
        API.getInstance().setToken(sessionId);
        Cookies.set("token", sessionId);
        const user = await User.getLoggedInUser();
        this.$store.commit("setUser", user);
        this.$parent.close();
      } catch (e) {
        this.isLoading = false;
        this.message = e;
      }
    }
  }
});
</script>