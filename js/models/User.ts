import Model from './Model';
import API from './API';

export default class User extends Model {
  ID = 0;

  Username = '';

  DisplayName = '';

  IsAdmin = false;

  static async getLoggedInUser(): Promise<User> {
    return User.buildModel<User>(API.getInstance().GET('login'), User);
  }

  static async logout(): Promise<void> {
    await API.getInstance().DELETE('login');
  }
}
