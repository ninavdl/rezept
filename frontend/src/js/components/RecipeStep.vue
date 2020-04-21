<template>
  <li style="margin-left: 1em; margin-bottom: 1em">
    <div class="box" style="margin-left: 1em">
      <img v-if="step.Image" v-bind:src="step.Image.URL" style="max-height: 20em" />
      <RecipeStepText :text="step.Text" :factor="factor" />
    </div>
  </li>
</template>

<script lang="ts">
import "reflect-metadata";
import { Component, Prop } from "vue-property-decorator";
import Vue from "vue";

import RecipeStepText from "./RecipeStepText.vue";

import Step from "../models/Step";

@Component({
  components: { RecipeStepText }
})
export default class RecipeStepComponent extends Vue {
  @Prop()
  step!: Step;

  @Prop()
  factor!: number;
}
</script>
