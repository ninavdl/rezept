<template>
  <li style="margin-left: 1em; margin-bottom: 1em">
    <div class="box" style="margin-left: 1em">
      <img v-if="step.Image" v-bind:src="step.Image.URL" style="max-height: 20em" />
      <p>{{ step.Text }}</p>
    </div>
  </li>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  props: ["step"]
});
</script>