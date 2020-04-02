import Model from "./Model";
import API from "./API";

export default class User extends Model {
    ID: number = 0;
    Username: string = "";
    DisplayName: string = "";
    IsAdmin: boolean = false;

    static async getLoggedInUser(): Promise<User> {
        return User.buildModel<User>(API.getInstance().GET("login"), User);
    }

    static async logout(): Promise<void> {
        await API.getInstance().DELETE("login");
    }
}