<template>
  <main class="recipe">
    <div>
      <b-loading :active="isLoading" />
    </div>
    <header class="level is-mobile">
      <div class="level-left">
        <div class="level-item">
          <h2 class="title" v-if="recipe.Name == ''">Recipe</h2>
          <h2 class="title" v-else>
            {{recipe.Name}}
            <small v-if="recipe.Creator != null" class="subtitle">
              by
              <router-link
                :to="{ name: 'list', query: { user: recipe.Creator.Username } }"
              >{{ recipe.Creator.DisplayName }}</router-link>
            </small>
          </h2>
        </div>
      </div>
      <div class="level-right" v-if="canModify">
        <router-link
          :to="{ name: 'editRecipe', params: { recipeID: recipe.ID } }"
          class="level-item button"
        >Edit</router-link>
        <b-button type="button" v-on:click="deleteRecipe" class="level-item">Delete</b-button>
      </div>
    </header>

    <img v-if="recipe.Image != null" v-bind:src="recipe.Image.URL" style="max-height: 40em" />

    <section id="info" class="container">
      <p>{{ recipe.Description }}</p>

      <section id="tags" v-if="recipe.Tags.length != 0">
        Tags:
        <ul class="tags">
          <li class="tag" v-for="(tag, i) in recipe.Tags" :key="'tag' + i">
            <router-link :to="{ name: 'list', query: { tags: tag } }">{{tag}}</router-link>
          </li>
        </ul>
      </section>
    </section>

    <div class="columns">
      <section class="column is-one-third">
        <div class="container">
          <h3 class="title">Ingredients</h3>
          <div class="box">
            <b-field label="Servings">
              <b-numberinput
                v-model="recipe.Servings"
                controls-position="compact"
                size="is-small"
                min="1"
                step="1"
              ></b-numberinput>
            </b-field>

            <ul style="list-style: disc; margin-left: 1em">
              <RecipeIngredient
                v-for="ingredient in recipe.Ingredients"
                v-bind:ingredient="ingredient"
                v-bind:amount="ingredient.Amount / initialServings * recipe.Servings"
                :key="ingredient.ID"
              />
            </ul>
          </div>
        </div>
      </section>

      <section class="column">
        <div class="container">
          <h3 class="title">Steps</h3>
          <ol>
            <RecipeStep v-for="(step, i) in recipe.Steps" v-bind:step="step" :key="'step' + i" />
          </ol>
        </div>
      </section>
    </div>
  </main>
</template>

<script lang="ts">
import 'reflect-metadata';
import { Prop, Component } from 'vue-property-decorator';
import Vue from 'vue';

import {
  Numberinput, Field, Navbar, Loading,
} from 'buefy';
import Recipe from '../models/Recipe';
import User from '../models/User';

import RecipeStep from './RecipeStep.vue';
import RecipeIngredient from './RecipeIngredient.vue';


Vue.use(Numberinput);
Vue.use(Field);
Vue.use(Navbar);
Vue.use(Loading);

@Component({
  components: {
    RecipeIngredient,
    RecipeStep,
  },
})
export default class RecipeComponent extends Vue {
  @Prop()
  recipeID!: number;

  recipe: Recipe = new Recipe();

  isLoading = true;

  // required to calculate amounts of ingredients when the number of servings change
  initialServings = 0;

  async created(): Promise<void> {
    await this.getRecipe();
    this.isLoading = false;
  }

  get isLoggedIn(): boolean {
    return this.$store.state.isLoggedIn;
  }

  get user(): User {
    return this.$store.state.user;
  }

  get canModify(): boolean {
    if (!this.isLoggedIn) return false;
    if (this.recipe.Creator != null && this.recipe.Creator.ID === this.user.ID) return true;
    return this.user.IsAdmin;
  }

  async getRecipe(): Promise<void> {
    this.recipe = await Recipe.getRecipe(this.recipeID);
    this.initialServings = this.recipe.Servings;
  }

  async deleteRecipe(): Promise<void> {
    await this.recipe.deleteRecipe();
    this.$router.push({ name: 'list' });
  }
}
</script>
