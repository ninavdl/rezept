import Ingredient from "./Ingredient";
import User from "./User";
import RecipeInfo from "./RecipeInfo";
import Step from "./Step";
import API from "./API";

export default class Recipe extends RecipeInfo {
    Servings: number = 0;
    Steps: Step[] = [];
    Tags: string[] = [];
    Ingredients: Ingredient[] = [];
    Description: string = "";
    Creator: User = null;

    public static async getRecipe(id: number): Promise<Recipe> {
        return this.buildModel<Recipe>(API.getInstance().GET(`recipes/${id}`), Recipe);
    }

    public async updateRecipe(): Promise<Recipe> {
        return Recipe.buildModel<Recipe>(API.getInstance().POST(`recipes/${this.ID}`, this), Recipe)
    }

    public async saveRecipe(): Promise<Recipe> {
        return Recipe.buildModel<Recipe>(API.getInstance().PUT(`recipes`, this), Recipe);
    }
}