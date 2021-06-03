import {Mutations} from "./mutations";
import {AppState} from "./state";
import {RootState} from "../index";
import {AppActionTypes} from "./action-types";
import { ActionTree, ActionContext } from 'vuex'
import {AppMutationTypes} from "./mutation-types";

type AugmentedActionContext = {
    commit<K extends keyof Mutations>(
        key: K,
        payload?: Parameters<Mutations[K]>[1],
    ): ReturnType<Mutations[K]>
} & Omit<ActionContext<AppState, RootState>, 'commit'>


export interface Actions {
    [AppActionTypes.ACTION_TOGGLE_SIDEBAR](
        {commit}: AugmentedActionContext
    ): void

    [AppActionTypes.ACTION_CLOSE_SIDEBAR](
        {commit}: AugmentedActionContext
    ): void
}

export const actions: ActionTree<AppState, RootState> & Actions = {
    [AppActionTypes.ACTION_TOGGLE_SIDEBAR]({commit}) {
        commit(AppMutationTypes.TOGGLE_SIDEBAR)
    },
    [AppActionTypes.ACTION_CLOSE_SIDEBAR]({commit}) {
        commit(AppMutationTypes.CLOSE_SIDEBAR)
    }
}
