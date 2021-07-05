export interface UserState {
    access_token: string;
    expiresAt: string;
    userName: string;
    nickName: string;
    headerImg: string;
}

export const state: UserState = {
    access_token: '',
    expiresAt: '',
    userName: '',
    nickName: '',
    headerImg: '',
}