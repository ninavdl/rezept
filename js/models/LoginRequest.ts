import Model from './Model';
import API from './API';
import LoginResponse from './LoginResponse';

export default class LoginRequest extends Model {
  Username: string;

  Password: string;

  async login(): Promise<string> {
    const login = await LoginRequest.buildModel<LoginResponse>(API.getInstance().PUT('login', this), LoginResponse);
    return login.SessionID;
  }
}
