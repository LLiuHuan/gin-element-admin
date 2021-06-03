import {AppState} from "./state";
import {AppMutationTypes} from "./mutation-types";

export type Mutations<S = AppState> = {
    [AppMutationTypes.TOGGLE_SIDEBAR](state: AppState): void
    [AppMutationTypes.CLOSE_SIDEBAR](state:AppState): void
}

export const mutations: Mutations<AppState> & Mutations = {
    [AppMutationTypes.TOGGLE_SIDEBAR](state: AppState) {
        state.sidebar.opened = !state.sidebar.opened
    },
    [AppMutationTypes.CLOSE_SIDEBAR](state: AppState) {
        state.sidebar.opened = false
    }
}