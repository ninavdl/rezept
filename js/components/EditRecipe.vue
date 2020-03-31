<script>
import EditorBase from "./EditorBase.vue";
import Recipe from "../models/Recipe.ts";

export default EditorBase.extend({
  created: async function() {
    this.isLoading = true;
    await this.getRecipe();
    this.isLoading = false;
  },
  props: ["recipeID"],
  methods: {
    submit: async function(ev) {
      ev.preventDefault();
      this.loadingText = "Saving recipe";
      this.isLoading = true;
      this.$data.recipe.updateRecipe();
      this.isLoading = false;
      this.$router.push({ name: "recipe", params: { recipeID: this.$data.recipe.ID } });
    },
    getRecipe: async function() {
      this.recipe = await Recipe.getRecipe(this.$props.recipeID);
    }
  }
});
</script>