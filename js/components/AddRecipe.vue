<script lang="ts">
import EditorBase from "./EditorBase.vue";

import { Component } from "vue-property-decorator";

@Component({})
export default class AddRecipeComponent extends EditorBase {
  async submit(ev) {
    ev.preventDefault();
    this.isLoading = true;
    const recipe = await this.$data.recipe.saveRecipe();
    this.isLoading = true;
    this.$router.push({ name: "recipe", params: { recipeID: recipe.ID } });
  }
}
</script>