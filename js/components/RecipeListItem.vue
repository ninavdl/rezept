<template>
  <div class="box">
    <router-link class="has-text-dark" :to="{name: 'recipe', params: { recipeID: recipe.ID}}">
      <div class="columns">
        <figure class="column is-narrow">
          <p class="image" style="width: 200px;">
            <img v-if="recipe.Image" v-bind:src="recipe.Image.ThumbnailURL" />
          </p>
        </figure>
        <div class="column">
          <header>
            <strong>{{ recipe.Name }}</strong>
          </header>
          <div class="content">{{ recipe.ShortDescription }}</div>
          <footer>{{ recipe.CreatedAt | formatDate }}</footer>
        </div>
      </div>
    </router-link>
  </div>
</template>

<script lang="ts">
import 'reflect-metadata';
import { Component, Prop } from 'vue-property-decorator';
import Vue from 'vue';

import { Menu } from 'buefy';
import Recipe from '../models/Recipe';

Vue.use(Menu);

@Component({})
export default class RecipeListItemComponent extends Vue {
  @Prop()
  recipe!: Recipe;
}
</script>
