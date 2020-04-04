<template>
  <form v-on:submit="login">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <h2 class="modal-card-title">Login</h2>
        <p v-if="message != ''">{{ message }}</p>
      </header>
      <section class="modal-card-body" v-if="!isLoading">
        <template v-if="!isLoading">
          <b-field label="Username">
            <b-input type="string" v-model="loginRequest.Username" />
          </b-field>
          <b-field label="Password">
            <b-input type="password" v-model="loginRequest.Password" />
          </b-field>
        </template>
        <Loading v-else text="Logging in"></Loading>
      </section>
      <footer class="modal-card-foot">
        <b-button tag="input" native-type="submit" value="Login" />
      </footer>
    </div>
  </form>
</template>

<script lang="ts">
import "reflect-metadata";
import { Component } from "vue-property-decorator"
import Vue from "vue";

import Loading from "./Loading.vue";

import LoginRequest from "../models/LoginRequest.ts";
import API from "../models/API.ts";
import User from "../models/User.ts";

import Cookies from "cookies-js"

import { Modal, Button, Field, Input } from "buefy";

Vue.use(Modal);
Vue.use(Button);
Vue.use(Field);
Vue.use(Input);

@Component({
  components: { Loading }
})
export default class LoginMenuComponent extends Vue {
  isLoading: boolean = false;
  loginRequest: LoginRequest = new LoginRequest();
  message: String = "";

  async login(ev): Promise<void> {
      ev.preventDefault();
      this.isLoading = true;
      try {
        const sessionId = await this.loginRequest.login();
        console.log(sessionId);
        API.getInstance().setToken(sessionId);
        Cookies.set("token", sessionId);
        const user = await User.getLoggedInUser();
        console.log(user);
        this.$store.commit("setUser", user);
        this.$parent.close();
      } catch (e) {
        this.isLoading = false;
        this.message = e;
      }
    }
}
</script>