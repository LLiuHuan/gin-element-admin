import https from "../utils/https";
import {ContentType, Method} from "axios-mapper";
import {RootObject} from "../model/globalModel";
import {Captcha, UserInfo} from "../model/response/loginModel";
import {LoginOptions} from "../model/request/loginModel";


export const getCaptcha = () => {
    return https().request<RootObject<Captcha>>('/v1/base/captcha', Method.POST, {}, ContentType.json)
}

export const login = (opt: LoginOptions) => {
    return https().request<RootObject<UserInfo>>('/v1/base/login', Method.POST, opt, ContentType.json)
}