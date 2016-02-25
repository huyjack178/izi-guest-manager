module Application.Interfaces {
    export interface IUserService {
        getUser: () => Array<IUser>;
    }

    export interface IUser {
        name: string;
    }
}