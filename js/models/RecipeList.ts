import Model from "./Model";
import RecipeInfo from "./RecipeInfo";
import API from "./API";

export default class RecipeList extends Model {
    Recipes: RecipeInfo[] = [];
    Pages: number = 0;
    Page: number = 0;
    Results: number = 0;

    public static async getRecipes(page: number, query: any): Promise<RecipeList> {
        let queryString = "";
        if ("tags" in query) {
            for (let tag of query.tags) {
                queryString += "&tag=" + encodeURIComponent(tag);
            }
        }
        if ("user" in query) {
            queryString += "&user=" + encodeURIComponent(query.user);
        }
        if ("keywords" in query) {
            for (let keyword of query.keywords) {
                queryString += "&keyword=" + encodeURIComponent(keyword);
            }
        }

        let req : Response = await API.getInstance().GET("recipes?page=" + page + queryString);
        let data : RecipeList = await req.json();
        console.log(data);
        let r = new RecipeList();
        r.Pages = data.Pages;
        r.Page = data.Page;
        r.Recipes = new Array<RecipeInfo>(data.Recipes.length);
        r.Results = data.Results;
        for (let i in data.Recipes) {
            r.Recipes[i] = new RecipeInfo();
            r.Recipes[i].assign(data.Recipes[i])
        }

        return r;
    }
}