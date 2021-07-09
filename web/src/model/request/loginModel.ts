export interface LoginOptions {
    username: string;
    password: string;
    captcha: string;   // 验证码
    captchaId: string; // 验证码ID
}


export interface MenuByIdOptions {
    AuthorityId: string;
}