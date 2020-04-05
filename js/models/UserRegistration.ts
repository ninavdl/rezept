import Model from './Model';
import User from './User';
import API from './API';

export default class UserRegistration extends Model {
  Username: string;

  DisplayName: string;

  Password: string;

  async signup(): Promise<User> {
    return UserRegistration.buildModel<User>(API.getInstance().PUT('users', this), User);
  }
}
