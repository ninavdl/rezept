<script lang="ts">
import "reflect-metadata";
import { Component, Prop } from "vue-property-decorator";

import EditorBase from "./EditorBase.vue";

import Recipe from "../models/Recipe";

@Component({})
export default class EditRecipeComponent extends EditorBase {
  @Prop()
  recipeID!: number;

  async created(): Promise<void> {
    this.isLoading = true;
    await this.getRecipe();
    this.isLoading = false;
  }

  async getRecipe(): Promise<void> {
    this.recipe = await Recipe.getRecipe(this.recipeID);
  }

  async submit(ev): Promise<void> {
    ev.preventDefault();
    this.loadingText = "Saving recipe";
    this.isLoading = true;
    this.$data.recipe.updateRecipe();
    this.isLoading = false;
    this.$router.push({
      name: "recipe",
      params: { recipeID: this.$data.recipe.ID }
    });
  }
}
</script>