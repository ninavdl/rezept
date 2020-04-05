<template>
  <form v-on:submit="signup">
    <div class="modal-card" style="width: auto">
      <div>
        <b-loading :active="isLoading" :isFullPage="false" />
      </div>
      <header class="modal-card-head">
        <h2 class="modal-card-title">Signup</h2>
        <p v-if="message != ''">{{ message }}</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Display name">
          <b-input type="string" v-model="registration.DisplayName" required />
        </b-field>
        <b-field label="Username">
          <b-input type="string" v-model="registration.Username" required />
        </b-field>
        <b-field label="Password" message="Minimum 8 characters">
          <b-input type="password" v-model="registration.Password" required pattern=".{8,}" />
        </b-field>
      </section>
      <footer class="modal-card-foot">
        <b-button tag="input" native-type="submit" value="Sign up" />
      </footer>
    </div>
  </form>
</template>

<script lang="ts">
import "reflect-metadata";
import { Component } from "vue-property-decorator";
import Vue from "vue";

import LoginRequest from "../models/LoginRequest";
import UserRegistration from "../models/UserRegistration";
import API from "../models/API";
import User from "../models/User";

import Cookies from "cookies-js";

import { Modal, Button, Field, Input, Loading } from "buefy";

Vue.use(Modal);
Vue.use(Button);
Vue.use(Field);
Vue.use(Input);
Vue.use(Loading);

@Component({})
export default class SignupMenuComponent extends Vue {
  isLoading: boolean = false;
  message: String = "";
  registration: UserRegistration = new UserRegistration();

  async signup(ev): Promise<void> {
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
</script>