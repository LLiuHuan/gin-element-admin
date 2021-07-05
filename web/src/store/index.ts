import {store as app, AppState, AppStore} from './app'
import {createStore, createLogger} from 'vuex'
import {store as user, UserState, UserStore} from "./user";

export interface RootState {
    app: AppState
    user: UserState
}

export type Store = AppStore<Pick<RootState, 'app'>> & UserStore<Pick<RootState, 'user'>>


// const debug = process.env.NODE_ENV !== 'production'
const debug = false
const plugins = debug ? [createLogger({})]: []


export const store = createStore({
    plugins,
    modules: {
        app,
        user
    }
})

export function useStore(): Store {
    return store as Store;
}