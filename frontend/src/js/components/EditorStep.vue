<template>
  <li class="column is-full">
    <div class="box">
      <b-field grouped>
        <b-field expanded label="Text">
          <b-input type="textarea" required v-model="thisStep.Text" v-on:input="updateStepText" />
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

<script lang="ts">
import 'reflect-metadata';
import { Component, Prop } from 'vue-property-decorator';
import Vue from 'vue';

import Uploader from './Uploader.vue';

import Step from '../models/Step';
import Image from '../models/Image';

@Component({
  components: { Uploader },
})
export default class EditorStepComponent extends Vue {
  @Prop()
  step!: Step;

  thisStep: Step = this.step;

  updateStepText(): void {
    this.$emit('update:step', this.thisStep);
  }

  setImage(image: Image): void {
    this.thisStep.Image = image;
    this.$emit('update:step', this.thisStep);
  }
}
</script>
