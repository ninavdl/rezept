import Model from './Model';
import RecipeInfo from './RecipeInfo';
import API from './API';

export default class RecipeList extends Model {
  Recipes: RecipeInfo[] = [];

  Pages = 0;

  Page = 0;

  Results = 0;

  public static async getDrafts(): Promise<RecipeList> {
    const req: Response = await API.getInstance().GET(`drafts`);
    const data: RecipeList = await req.json();
    const r = new RecipeList();
    r.Pages = data.Pages;
    r.Page = data.Page;
    r.Results = data.Results;
    r.Recipes = data.Recipes.map((recipe) => {
      const recipeInfo = new RecipeInfo();
      recipeInfo.assign(recipe);
      return recipeInfo;
    });

    return r;
  }

  public static async getRecipes(page: number, query: any): Promise<RecipeList> {
    let queryString = '';
    if ('tags' in query) {
      queryString += query.tags.map((tag) => `&tag=${encodeURIComponent(tag)}`).join('');
    }
    if ('user' in query) {
      queryString += `&user=${encodeURIComponent(query.user)}`;
    }
    if ('keywords' in query) {
      queryString += query.keywords.map(
        (keyword) => `&keyword=${encodeURIComponent(keyword)}`,
      ).join('');
    }

    const req: Response = await API.getInstance().GET(`recipes?page=${page}${queryString}`);
    const data: RecipeList = await req.json();
    const r = new RecipeList();
    r.Pages = data.Pages;
    r.Page = data.Page;
    r.Results = data.Results;
    r.Recipes = data.Recipes.map((recipe) => {
      const recipeInfo = new RecipeInfo();
      recipeInfo.assign(recipe);
      return recipeInfo;
    });

    return r;
  }
}
