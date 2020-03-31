<template>
  <main class="editor">
    <Loading :text="loadingText" v-if="isLoading"></Loading>
    <template v-else>
      <header class="level is-mobile">
        <div class="level-left">
              <h2 class="level-item title" v-if="recipe.ID != 0">Edit recipe</h2>
              <h2 class="level-item title" v-else>Add recipe</h2>
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
              <b-taginput v-model="recipe.Tags" ellipsis icon="label" placeholder="Add a tag"></b-taginput>
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
              <b-numberinput
                v-model="recipe.Servings"
                controls-position="compact"
                min="1"
                step="1"
              />
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
            <b-button v-on:click="addStep" type="is-success" icon-left="plus-circle">Add step</b-button>
          </div>
        </section>

        <b-button tag="input" native-type="submit" type="is-primary" value="Save" expanded />
      </form>
    </template>
  </main>
</template>

<script>
import EditorStep from "./EditorStep.vue";
import EditorIngredient from "./EditorIngredient.vue";
import Vue from "vue";
import Loading from "./Loading.vue";
import Uploader from "./Uploader.vue";
import Recipe from "../models/Recipe.ts";
import Ingredient from "../models/Ingredient.ts";

import { Numberinput, Input, Taginput } from "buefy";

Vue.use(Numberinput);
Vue.use(Input);
Vue.use(Taginput);

export default Vue.extend({
  methods: {
    addIngredient: async function(ev) {
      this.recipe.Ingredients.push(new Ingredient());
    },
    addStep: async function(ev) {
      this.recipe.Steps.push("");
    },
    setImage: function(image) {
      this.recipe.Image = image;
    },
    updateStep(i, step) {
      this.recipe.Steps[i] = step;
      console.log(step);
    },
    updateIngredient(i, ingredient) {
      this.recipe.Ingredients[i] = ingredient;
    }
  },
  data: () => ({
    recipe: new Recipe(),
    isLoading: false,
    loadingText: "Loading recipe"
  }),
  components: {
    EditorStep,
    Loading,
    EditorIngredient,
    Uploader
  }
});
</script>
