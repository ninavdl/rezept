<template>
  <li class="column is-full">
    <div class="box">
      <b-field grouped>
        <b-field expanded label="Text">
          <b-input type="textarea" required v-model="thisStep.Text" v-on:change="updateStepText" />
        </b-field>
        <b-field label="Picture">
          <Uploader v-bind:image="thisStep.Image" v-on:setImage="setImage"></Uploader>
        </b-field>
      </b-field>
      <p class="control">
        <b-button v-on:click="$emit('remove')" type="is-danger" icon-left="delete">Delete</b-button>
      </p>
    </div>
  </li>
</template>

<style lang="scss" scoped>
@import "../../sass/_variables.scss";

#steps {
  padding-left: 1.5em;
}

.step {
  margin-left: 1.5em;
}

.step > div {
  display: flex;
  width: 100%;
  padding-left: 1em;
  padding-right: 0.75em;

  textarea {
    width: 100%;
    margin-right: 0.75em;
  }

  margin-bottom: 1em;
}
</style>

<script>
import Vue from "vue";
import Uploader from "./Uploader.vue";

export default Vue.extend({
  data: function() {
    return {
      thisStep: this.step
    };
  },
  methods: {
    updateStepText() {
      this.$emit("update:step", this.thisStep);
    },
    setImage(image) {
      console.log(image);
      this.thisStep.Image = image;
      this.$emit("update:step", this.thisStep);
    }
  },
  props: ["step"],
  components: { Uploader }
});
</script>