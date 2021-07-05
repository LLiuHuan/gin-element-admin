import {Mutations} from "./mutations";
import {ActionContext, ActionTree} from "vuex";
import {state, UserState} from "./state";
import {RootState} from "../index";
import {UserActionTypes} from "./action-types";
import {UserMutationTypes} from "./mutation-types";
import {login} from "../../apis/login";

type AugmentedActionContext = {
    commit<K extends keyof Mutations>(
        key: K,
        payload: Parameters<Mutations[K]>[1],
    ): ReturnType<Mutations[K]>
} & Omit<ActionContext<UserState, RootState>, 'commit'>

export interface Actions {
    [UserActionTypes.ACTION_LOGIN](
        { commit }: AugmentedActionContext,
        loginOption: { username: string, password: string, captcha: string, captchaId: string }
    ): void
    [UserActionTypes.ACTION_RESET_TOKEN](
        { commit }: AugmentedActionContext
    ): void
    [UserActionTypes.ACTION_GET_USER_INFO](
        { commit }: AugmentedActionContext
    ): void
}

export const actions: ActionTree<UserState, RootState> & Actions = {
    async [UserActionTypes.ACTION_LOGIN](
        { commit }: AugmentedActionContext,
        loginOption: { username: string, password: string, captcha: string, captchaId: string }
    ) {
        let { username, password, captcha, captchaId } = loginOption
        username = username.trim()
        await login({ username, password, captcha, captchaId }).then(async(res) => {
            if (res?.code === 0 && res.data.access_token) {
                commit(UserMutationTypes.SET_TOKEN, res.data.access_token)
            }
        }).catch((err) => {
            console.log(err)
        })
    },

    [UserActionTypes.ACTION_RESET_TOKEN](
        { commit }: AugmentedActionContext) {
        commit(UserMutationTypes.SET_TOKEN, '')
    },

    async [UserActionTypes.ACTION_GET_USER_INFO](
        { commit }: AugmentedActionContext
    ) {
        if (state.access_token === '') {
            throw Error('token is undefined!')
        }
        // await userInfoRequest().then((res) => {
        //     if (res?.code === 0) {
        //         commit(UserMutationTypes.SET_ROLES, res.data.roles)
        //         commit(UserMutationTypes.SET_NAME, res.data.name)
        //         commit(UserMutationTypes.SET_AVATAR, res.data.avatar)
        //         commit(UserMutationTypes.SET_INTRODUCTION, res.data.introduction)
        //         commit(UserMutationTypes.SET_EMAIL, res.data.email)
        //         return res
        //     } else {
        //         throw Error('Verification failed, please Login again.')
        //     }
        // })
    },
    [UserActionTypes.ACTION_LOGIN_OUT](
        { commit }: AugmentedActionContext
    ) {
        // console.log(commit)
        // removeToken()
        commit(UserMutationTypes.SET_TOKEN, '')
        // commit(UserMutationTypes.SET_ROLES, [])
        // resetRouter()
    }
}
