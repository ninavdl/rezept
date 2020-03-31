<script>
import EditorBase from "./EditorBase.vue";

export default EditorBase.extend({
  methods: {
    submit: async function(ev) {
      ev.preventDefault();
      this.isLoading = true;  
      const recipe = await this.$data.recipe.saveRecipe();
      this.isLoading = true;
      this.$router.push({ name: "recipe", params: { recipeID: recipe.ID } });
    }
  }
});
</script>