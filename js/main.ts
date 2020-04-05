import VueRouter from 'vue-router';
import Vue from 'vue';
import Vuex from 'vuex';

import RecipeList from './components/RecipeList.vue';
import Recipe from './components/Recipe.vue';
import AddRecipe from './components/AddRecipe.vue';
import EditRecipe from './components/EditRecipe.vue';
import App from "./components/App.vue";

import API from "./models/API.ts";

import Cookies from "cookies-js";
import User from "./models/User.ts";

export default async function (config) {
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
        ],
        base: config.PathPrefix
    });

    const sessionId = Cookies.get("token");
    let user = null;
    if (sessionId) {
        API.getInstance().setToken(sessionId);
        user = await User.getLoggedInUser();
    }

    Vue.filter('round', (value, decimals) => Math.round(value * Math.pow(10, decimals)) / Math.pow(10, decimals))
    Vue.filter('formatDate', date => date.toLocaleDateString())

    let store = new Vuex.Store({
        state: {
            user: user == null ? {} : user,
            isLoggedIn: user != null
        },
        mutations: {
            setUser(state, newUser) {
                state.user = Object.assign({}, state.user, newUser);
                state.isLoggedIn = "ID" in newUser
            }
        }
    })

    const page = new Vue({
        router: router,
        store,
        render: createElement => createElement(App)
    });

    page.$mount("#app");
}
