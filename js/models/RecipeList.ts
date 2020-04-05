import Model from './Model';
import RecipeInfo from './RecipeInfo';
import API from './API';

export default class RecipeList extends Model {
  Recipes: RecipeInfo[] = [];

  Pages = 0;

  Page = 0;

  Results = 0;

  public static async getRecipes(page: number, query: any): Promise<RecipeList> {
    let queryString = '';
    if ('tags' in query) {
      for (const tag of query.tags) {
        queryString += `&tag=${encodeURIComponent(tag)}`;
      }
    }
    if ('user' in query) {
      queryString += `&user=${encodeURIComponent(query.user)}`;
    }
    if ('keywords' in query) {
      for (const keyword of query.keywords) {
        queryString += `&keyword=${encodeURIComponent(keyword)}`;
      }
    }

    const req: Response = await API.getInstance().GET(`recipes?page=${page}${queryString}`);
    const data: RecipeList = await req.json();
    console.log(data);
    const r = new RecipeList();
    r.Pages = data.Pages;
    r.Page = data.Page;
    r.Recipes = new Array<RecipeInfo>(data.Recipes.length);
    r.Results = data.Results;
    for (const i in data.Recipes) {
      r.Recipes[i] = new RecipeInfo();
      r.Recipes[i].assign(data.Recipes[i]);
    }

    return r;
  }
}
