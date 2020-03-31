<template>
  <form v-on:submit="login">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <h2 class="modal-card-title">Login</h2>
        <p v-if="Message != ''">{{ Message }}</p>
      </header>
      <section class="modal-card-body" v-if="!isLoading">
        <template v-if="!isLoading">
          <b-field label="Username">
            <b-input type="string" v-model="Username" />
          </b-field>
          <b-field label="Password">
            <b-input type="password" v-model="Password" />
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

<script>
import Vue from "vue";
import Loading from "./Loading.vue";

import {Modal, Button, Field, Input} from "buefy";

Vue.use(Modal);
Vue.use(Button);
Vue.use(Field);
Vue.use(Input);

export default Vue.extend({
  data: () => ({
    isLoading: false,
    Username: "",
    Password: "",
    Message: ""
  }),
  components: {
    Loading
  },
  methods: {
    login: async function(ev) {
      ev.preventDefault();
      this.isLoading = true;
      try {
        await this.$controller.login(this.Username, this.Password);
      } catch (e) {
        this.isLoading = false;
        this.Message = e;
      }
      this.$store.commit("setUser", this.$controller.user);
      this.$parent.close();
    }
  }
});
</script>