<template>
  <main class="recipe">
    <Loading text="Loading recipe" v-if="isLoading"></Loading>
    <template v-else>
      <header class="level is-mobile">
        <div class="level-left">
          <div class="level-item">
            <h2 class="title">
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
        <div class="level-right">
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
                  v-bind="ingredient"
                  v-bind:Amount="ingredient.Amount / initialServings * recipe.Servings"
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
    </template>
  </main>
</template>

<script>
import Recipe from "../models/Recipe.ts";
import RecipeStep from "./RecipeStep.vue";
import RecipeIngredient from "./RecipeIngredient.vue";
import Loading from "./Loading.vue";
import Vue from "vue";

import { Numberinput, Field, Navbar } from "buefy";

Vue.use(Numberinput);
Vue.use(Field);
Vue.use(Navbar);

export default Vue.extend({
  created: async function() {
    await this.getRecipe();
    this.isLoading = false;
  },
  props: ["recipeID"],
  data: () => ({
    recipe: null,
    isLoading: true,
    initialServings: 0 // required to calculate amounts of ingredients when the number of servings changes
  }),
  computed: {
    isLoggedIn() {
      return this.$store.state.isLoggedIn;
    }
  },
  methods: {
    getRecipe: async function() {
      this.recipe = await Recipe.getRecipe(this.$props.recipeID);
      console.log(this.recipe);
      this.initialServings = this.recipe.Servings;
    },
    deleteRecipe: async function(ev) {
      this.$controller.showLoadingScreen();
      let res = await this.$controller.delete(
        `recipes/${this.$props.recipeID}`
      );
      if (!res.ok) {
        let data = await res.json();
        throw data.Error;
      }
      this.$router.push({ name: "index" });
    },
    canModify: function() {
      if (this.$controller.user == null) return false;
      if (
        this.recipe.Creator != null &&
        this.recipe.Creator.ID == this.$controller.user.ID
      )
        return true;
      return this.$controller.user.IsAdmin;
    }
  },
  components: {
    RecipeIngredient,
    RecipeStep,
    Loading
  }
});
</script>
