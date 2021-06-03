import {store as app, AppState, AppStore} from './app'
import {createStore, createLogger} from 'vuex'

export interface RootState {
    app: AppState
}

export type Store = AppStore<Pick<RootState, 'app'>>


// const debug = process.env.NODE_ENV !== 'production'
const debug = false
const plugins = debug ? [createLogger({})]: []


export const store = createStore({
    plugins,
    modules: {
        app,
    }
})

export function useStore(): Store {
    return store as Store;
}