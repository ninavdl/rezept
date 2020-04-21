import Model from './Model';
import API from './API';
import Image from './Image';

export default class RecipeInfo extends Model {
  ID = 0;

  Name = '';

  ShortDescription = '';

  CreatedAt: Date = new Date();

  UpdatedAt: Date = new Date();

  Image: Image = null;

  Published: boolean = true;

  public static async getRecipes(): Promise<RecipeInfo[]> {
    return this.buildModels<RecipeInfo>(API.getInstance().GET('recipes'), RecipeInfo);
  }
}
