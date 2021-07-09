import Cookies from 'js-cookie'
import Keys from '../model/keys'

export const getToken = () => Cookies.get(Keys.tokenKey)
export const setToken = (token: string) => Cookies.set(Keys.tokenKey, token)
export const removeToken = () => Cookies.remove(Keys.tokenKey)