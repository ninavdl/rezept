import Model from "./Model";
import API from "./API";

export default class PageData extends Model {
    Users: number = 0;
    SignupAllowed: boolean = false;

    static async getPageData(): Promise<PageData> {
        return PageData.buildModel<PageData>(API.getInstance().GET("data"), PageData);
    }
}
