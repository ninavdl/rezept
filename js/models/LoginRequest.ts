import Model from "./Model";
import API from "./API";

export default class LoginRequest extends Model {
    Username: string;
    Password: string;

    async login(): Promise<string> {
        const login = await LoginRequest.buildModel<LoginResponse>(API.getInstance().PUT("login", this), LoginResponse);
        console.log(login);
        return login.SessionID;
    }
}

class LoginResponse extends Model {
    SessionID: string = "";
}