import Model from "./Model";
import API from "./API";
import Image from "./Image";

export default class RecipeInfo extends Model {
    ID: number = 0;
    Name: string = "";
    ShortDescription: string = "";
    CreatedAt: Date = new Date();
    UpdatedAt: Date = new Date();
    Image: Image = null;

    constructor() {
        super();
    }

    public static async getRecipes(): Promise<RecipeInfo[]> {
        return this.buildModels<RecipeInfo>(API.getInstance().GET("recipes"), RecipeInfo);
    }
}