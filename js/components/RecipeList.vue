<template>
  <main class="recipe-list section">
    <header class="level is-mobile">
      <div class="level-left">
        <h2 class="level-item title">Recipes</h2>
      </div>
      <div class="level-right">
        <router-link v-if="isLoggedIn" :to="{ name: 'addRecipe' }" class="level-item button">Add recipe</router-link>
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
    <section class="section">
      <Loading text="Loading recipes" v-if="isLoading"></Loading>
      <template v-else>
        <p v-if="recipeList.Results == 0">No recipe found :(</p>
        <p v-else-if="recipeList.Results == 1">Found one recipe:</p>
        <p v-else>Found {{ recipeList.Results }} recipes:</p>

        <div class="container">
          <RecipeListItem
            v-for="recipe in recipeList.Recipes"
            v-bind:recipe="recipe"
            :key="recipe.ID"
          />
        </div>

        <section class="section" v-if="recipeList.Pages > 1">
          <b-pagination
            v-bind:total="recipeList.Pages"
            v-bind:current.sync="pageNum == null ? 1 : pageNum"
            v-bind:per-page="1"
            v-on:change="setPage"
          />
        </section>
      </template>
    </section>
  </main>
</template>

<script>
import Vue from "vue";
import RecipeListItem from "./RecipeListItem.vue";
import Loading from "./Loading.vue";
import RecipeList from "../models/RecipeList.ts";

import { Menu, Pagination, Collapse } from "buefy";

Vue.use(Menu);
Vue.use(Pagination);
Vue.use(Collapse);

export default Vue.extend({
  created: async function() {
    await this.getRecipes();
    this.isLoading = false;
  },
  components: {
    RecipeListItem,
    Loading
  },
  async beforeRouteUpdate(to, from, next) {
    console.log(to, from, next);
    this.isLoading = true;
    this.query = to.query;
    await this.getRecipes();
    this.isLoading = false;
    next();
  },
  data: function() {
    console.log(this.$route.query);
    return {
      user: this.$controller.user,
      recipeList: new RecipeList(),
      isLoading: true,
      query: this.$route.query
    };
  },
  computed: {
    isLoggedIn() {
      return this.$store.state.isLoggedIn;
    },
    search() {
      return {
        tags: "tags" in this.query ? this.query.tags.split(",") : [],
        user: "user" in this.query ? this.query.user : "",
        keywords: "keywords" in this.query ? this.query.keywords.split(",") : ""
      };
    },
    pageNum() {
      return "page" in this.query ? this.query.page : 1;
    }
  },
  methods: {
    getQueryObject() {
      let queryObject = {};
      if (this.search.tags.length > 0)
        queryObject.tags = this.search.tags.join(",");
      if (this.search.user != "") queryObject.user = this.search.user;
      if (this.search.keywords.length > 0)
        queryObject.keywords = this.search.keywords.join(",");
      return queryObject;
    },
    async getRecipes() {
      this.recipeList = await RecipeList.getRecipes(this.pageNum, this.search);
    },
    async setPage(page) {
      let queryObject = this.getQueryObject();
      queryObject.page = page;
      this.$router.push({
        name: "list",
        query: queryObject
      });
    },
    updateSearch() {
      this.$router.push({ name: "list", query: this.getQueryObject() });
    }
  },
  props: ["page"]
});
</script>