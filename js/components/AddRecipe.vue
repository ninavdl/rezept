<script lang="ts">
import { Component } from 'vue-property-decorator';
import EditorBase from './EditorBase.vue';


@Component({})
export default class AddRecipeComponent extends EditorBase {
  heading = 'Add recipe';

  async submit(ev) {
    ev.preventDefault();
    this.isLoading = true;
    const recipe = await this.$data.recipe.saveRecipe();
    this.isLoading = true;
    this.$router.push({ name: 'recipe', params: { recipeID: recipe.ID } });
  }
}
</script>
