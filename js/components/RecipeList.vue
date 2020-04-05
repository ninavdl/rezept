<template>
  <main class="recipe-list section">
    <header class="level is-mobile">
      <div class="level-left">
        <h2 class="level-item title">Recipes</h2>
      </div>
      <div class="level-right">
        <router-link
          v-if="isLoggedIn"
          :to="{ name: 'addRecipe' }"
          class="level-item button"
        >Add recipe</router-link>
      </div>
    </header>

    <section class="section">
      <b-collapse
        class="card"
        animation="slide"
        aria-id="search"
        :open="search.tags.length != 0 || search.keywords.length != 0 || search.user != ''"
      >
        <div
          slot="trigger"
          slot-scope="props"
          class="card-header"
          role="button"
          aria-controls="search"
        >
          <p class="card-header-title">Search recipes</p>
          <a class="card-header-icon">
            <b-icon :icon="props.open ? 'menu-down' : 'menu-up'"></b-icon>
          </a>
        </div>
        <div class="card-content">
          <div class="content">
            <b-field label="Search by tags">
              <b-taginput
                v-model="search.tags"
                ellipsis
                icon="label"
                placeholder="Tag"
                maxtags="10"
                v-on:add="updateSearch"
                v-on:remove="updateSearch"
              ></b-taginput>
            </b-field>
            <b-field label="Search by keywords">
              <b-taginput
                v-model="search.keywords"
                ellipsis
                icon="label"
                placeholder="Keyword"
                maxtags="10"
                v-on:add="updateSearch"
                v-on:remove="updateSearch"
              ></b-taginput>
            </b-field>
            <b-field label="Search by username">
              <b-input v-model="search.user" v-on:input="updateSearch" placeholder="Username"></b-input>
            </b-field>
          </div>
        </div>
      </b-collapse>
    </section>
    <div>
      <div><b-loading :active="isLoading" :isFullPage="false" /></div>
      <section class="section">
        <b-message>
          <p v-if="recipeList.Results == 0">No recipe found :(</p>
          <p v-else-if="recipeList.Results == 1">Found one recipe:</p>
          <p v-else>Found {{ recipeList.Results }} recipes:</p>
        </b-message>

        <div class="container">
          <RecipeListItem
            v-for="recipe in recipeList.Recipes"
            v-bind:recipe="recipe"
            :key="recipe.ID"
          />
        </div>
      </section>

      <section class="section" v-if="recipeList.Pages > 1">
        <b-pagination
          v-bind:total="recipeList.Pages"
          v-bind:current.sync="pageNum"
          v-bind:per-page="1"
          v-on:change="setPage"
        />
      </section>
    </div>
  </main>
</template>

<script lang="ts">
import Vue from 'vue';
import { Component } from 'vue-property-decorator';
import {
  Menu, Pagination, Collapse, Message, Loading,
} from 'buefy';
import { Route } from 'vue-router';
import RecipeListItem from './RecipeListItem.vue';
import RecipeList from '../models/RecipeList';
import User from '../models/User';


[Menu, Pagination, Collapse, Message, Loading].forEach((c) => Vue.use(c));

class SearchObject {
  tags: string[];

  keywords: string[];

  user: string;

  page: number;
}

class QueryObject {
  tags: string;

  keywords: string;

  user: string;

  page: string;
}

@Component({
  components: {
    RecipeListItem,
  },
  async beforeRouteUpdate(to: Route, from: Route, next: Function) {
    this.isLoading = true;
    this.query = to.query;
    this.search = this.getSearchObject();
    await this.getRecipes();
    this.isLoading = false;
    next();
  },
})
export default class RecipeListComponent extends Vue {
  recipeList: RecipeList = new RecipeList();

  isLoading = true;

  query: any;

  pageNum: number;

  search: SearchObject;

  get isLoggedIn(): boolean {
    return this.$store.state.isLoggedIn;
  }

  get user(): User {
    return this.$store.state.user;
  }

  getSearchObject(): SearchObject {
    const q = new SearchObject();
    q.tags = 'tags' in this.query ? this.query.tags.split(',') : [];
    q.user = 'user' in this.query ? this.query.user : '';
    q.keywords = 'keywords' in this.query ? this.query.keywords.split(',') : [];
    return q;
  }

  async created(): Promise<void> {
    this.query = this.$route.query;
    this.search = this.getSearchObject();
    this.pageNum = 'page' in this.query ? parseInt(this.query.page) : 1;
    await this.getRecipes();
    this.isLoading = false;
  }

  getQueryObject(): QueryObject {
    const queryObject = new QueryObject();
    if (this.search.tags.length > 0) queryObject.tags = this.search.tags.join(',');
    if (this.search.user != '') queryObject.user = this.search.user;
    if (this.search.keywords.length > 0) queryObject.keywords = this.search.keywords.join(',');
    return queryObject;
  }

  async getRecipes(): Promise<void> {
    this.recipeList = await RecipeList.getRecipes(this.pageNum, this.search);
  }

  async setPage(page: number): Promise<void> {
    this.pageNum = page;
    const queryObject = this.getQueryObject();
    queryObject.page = page.toString();
    this.$router.push({
      name: 'list',
      query: <any>queryObject,
    });
  }

  updateSearch(): void {
    this.$router.push({ name: 'list', query: <any> this.getQueryObject() });
  }
}
</script>
