<template>
  <main class="editor">
    <div>
      <b-loading :active="isLoading" />
    </div>

    <header class="level is-mobile">
      <div class="level-left">
        <h2 class="level-item title">{{ heading }}</h2>
      </div>
      <div class="level-right">
        <router-link
          v-if="recipe.ID != 0"
          :to="{ name: 'recipe', params: { recipeID: recipe.ID }}"
          class="level-item button"
        >Show recipe</router-link>
      </div>
    </header>

    <form v-on:submit="submit">
      <section class="section">
        <div class="container">
          <b-field>
            <b-input v-model="recipe.Name" required size="is-large" placeholder="Title" />
          </b-field>

          <b-field label="Tags">
            <b-taginput v-model="recipe.Tags" ellipsis icon="label" placeholder="Add a tag" />
          </b-field>

          <b-field label="Short description that is shown in the recipe list">
            <b-input v-model="recipe.ShortDescription" type="textarea" />
          </b-field>

          <b-field label="Long description that is shown at the recipe page">
            <b-input v-model="recipe.Description" type="textarea" />
          </b-field>

          <b-field label="Picture">
            <Uploader v-bind:image="recipe.Image" v-on:setImage="setImage"></Uploader>
          </b-field>

          <b-field label="Servings">
            <b-numberinput v-model="recipe.Servings" controls-position="compact" min="1" step="1" />
          </b-field>
        </div>
      </section>

      <section class="section">
        <div class="container">
          <h3 class="title">Ingredients</h3>
          <ul class="columns is-multiline">
            <EditorIngredient
              v-for="(ingredient, i) in recipe.Ingredients"
              v-bind:ingredient="ingredient"
              v-on:update:ingredient="updateIngredient(i, $event)"
              :key="'ingredient' + i"
              v-on:remove="recipe.Ingredients.splice(i, 1)"
            />
          </ul>
          <b-button
            v-on:click="addIngredient"
            type="is-success"
            icon-left="plus-circle"
          >Add ingredient</b-button>
        </div>
      </section>

      <section class="section">
        <div class="container">
          <h3 class="title">Steps</h3>
          <ol class="columns is-multiline">
            <EditorStep
              v-for="(step, i) in recipe.Steps"
              v-bind:step="step"
              v-on:update:step="updateStep(i, $event)"
              :key="'step' + i"
              v-on:remove="recipe.Steps.splice(i, 1)"
            />
          </ol>
          <b-button v-on:click="addStep" type="is-success" icon-left="plus-circle">
            Add step
          </b-button>
        </div>
      </section>

      <b-button tag="input" native-type="submit" type="is-primary" value="Save" expanded />
    </form>
  </main>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'vue-property-decorator';

import {
  Numberinput, Input, Taginput, Loading,
} from 'buefy';
import EditorStep from './EditorStep.vue';
import EditorIngredient from './EditorIngredient.vue';
import Uploader from './Uploader.vue';

import Recipe from '../models/Recipe';
import Ingredient from '../models/Ingredient';
import Step from '../models/Step';
import Image from '../models/Image';

Vue.use(Numberinput);
Vue.use(Input);
Vue.use(Taginput);
Vue.use(Loading);

@Component({
  components: {
    EditorStep,
    EditorIngredient,
    Uploader,
  },
})
export default class EditorBaseComponent extends Vue {
  recipe: Recipe = new Recipe();

  isLoading = false;

  heading = '';

  addIngredient(): void {
    this.recipe.Ingredients.push(new Ingredient());
  }

  addStep(): void {
    this.recipe.Steps.push(new Step());
  }

  setImage(image: Image): void {
    this.recipe.Image = image;
  }

  updateStep(i: number, step: Step): void {
    this.recipe.Steps[i] = step;
  }

  updateIngredient(i: number, ingredient: Ingredient): void {
    this.recipe.Ingredients[i] = ingredient;
  }
}
</script>
