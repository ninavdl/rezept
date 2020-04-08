<script lang="ts">
import "reflect-metadata";
import { Component, Prop } from "vue-property-decorator";

import EditorBase from "./EditorBase.vue";

import Recipe from "../models/Recipe";

import { Metadata } from "../metadata.ts";

@Component({})
export default class EditRecipeComponent extends EditorBase {
  @Prop()
  recipeID!: number;

  heading = "Edit recipe";

  async created(): Promise<void> {
    this.isLoading = true;
    await this.getRecipe();
    await this.getRecipe();
    const metadata = new Metadata();
    metadata.title = "Edit: " + this.recipe.Name;
    metadata.description = this.recipe.ShortDescription;
    metadata.author =
      this.recipe.Creator != null ? this.recipe.Creator.DisplayName : null;
    metadata.date = this.recipe.UpdatedAt;
    metadata.imageURL =
      this.recipe.Image != null ? this.recipe.Image.URL : null;
    metadata.tags = this.recipe.Tags;
    this.$store.commit("setMetadata", metadata);

    this.isLoading = false;
  }

  async getRecipe(): Promise<void> {
    this.recipe = await Recipe.getRecipe(this.recipeID);
  }

  async submit(ev): Promise<void> {
    ev.preventDefault();
    this.isLoading = true;
    await this.$data.recipe.updateRecipe();
    this.isLoading = false;
    this.$router.push({
      name: "recipe",
      params: { recipeID: this.$data.recipe.ID }
    });
  }
}
</script>
