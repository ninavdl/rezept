import Controller from "./controller.js";

import VueRouter from 'vue-router';
import Vue from 'vue';
import Vuex from 'vuex';

import RecipeList from './components/RecipeList.vue';
import Recipe from './components/Recipe.vue';
import AddRecipe from './components/AddRecipe.vue';
import EditRecipe from './components/EditRecipe.vue';
import API from "./models/API.ts";


export default function (config) {
    API.init(config.APIPrefix);

    Vue.use(VueRouter);
    Vue.use(Vuex);

    const router = new VueRouter({
        mode: "history",
        routes: [
            { path: '/', name: "list", component: RecipeList, props: true },
            { path: '/recipe/:recipeID/edit', name: 'editRecipe', component: EditRecipe, props: true },
            { path: '/recipe/add', name: 'addRecipe', component: AddRecipe },
            { path: '/recipe/:recipeID', name: "recipe", component: Recipe, props: true },
        ]
    })


    new Controller(config, router);
}
