import Model from "./Model";

export default class User extends Model {
    ID: number;
    Username: string;
    DisplayName: string;
    IsAdmin: boolean;
}