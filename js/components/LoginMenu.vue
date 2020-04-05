<template>
  <form v-on:submit="login">
    <div class="modal-card" style="width: auto">
      <div>
        <b-loading :active="isLoading" :isFullPage="false" />
      </div>
      <header class="modal-card-head">
        <h2 class="modal-card-title">Login</h2>
        <p v-if="message != ''">{{ message }}</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Username">
          <b-input type="string" v-model="loginRequest.Username" />
        </b-field>
        <b-field label="Password">
          <b-input type="password" v-model="loginRequest.Password" />
        </b-field>
      </section>
      <footer class="modal-card-foot">
        <b-button tag="input" native-type="submit" value="Login" />
      </footer>
    </div>
  </form>
</template>

<script lang="ts">
import 'reflect-metadata';
import { Component } from 'vue-property-decorator';
import Vue from 'vue';

import Cookies from 'cookies-js';
import {
  Modal, Button, Field, Input, Loading,
} from 'buefy';
import LoginRequest from '../models/LoginRequest';
import API from '../models/API';
import User from '../models/User';


Vue.use(Modal);
Vue.use(Button);
Vue.use(Field);
Vue.use(Input);
Vue.use(Loading);

@Component({})
export default class LoginMenuComponent extends Vue {
  isLoading = false;

  loginRequest: LoginRequest = new LoginRequest();

  message = '';

  async login(ev): Promise<void> {
    ev.preventDefault();
    this.isLoading = true;
    try {
      const sessionId = await this.loginRequest.login();
      API.getInstance().setToken(sessionId);
      Cookies.set('token', sessionId);
      const user = await User.getLoggedInUser();
      this.$store.commit('setUser', user);
      this.$parent.close();
    } catch (e) {
      this.isLoading = false;
      this.message = e;
    }
  }
}
</script>
