// /*
//  * @Description: 权限
//  * @Author: LLiuHuan
//  * @Date: 2021-02-18 13:40:05
//  * @LastEditTime: 2021-03-02 18:36:14
//  * @LastEditors: LLiuHuan
//  */
//


// import router from './router'
//
// router.beforeEach(async(to: RouteLocationNormalized, _: RouteLocationNormalized, next: any) => {
//   console.log(to)
//   next()
// })


// @ts-ignore
import NProgress from 'nprogress'
import 'nprogress/css/nprogress.scss'
import router from './router'
import { RouteLocationNormalized } from 'vue-router'
import { useStore } from './store'
import { ElMessage } from 'element-plus'
import {getMenu} from "./apis/login";
import {getToken} from "./utils/cookies";
NProgress.configure({ showSpinner: false })

const whiteList = ['/login', '/redirect']

router.beforeEach(async(to: RouteLocationNormalized, _: RouteLocationNormalized, next: any) => {
  // Start progress bar
  NProgress.start()
  const store = useStore()
  // Determine whether the user has logged in
  console.log("token", store.state.user.access_token)
  const token = getToken()
  if (token) {
    if (to.path === '/login') {
      // If is logged in, redirect to the home page
      next({ path: '/' })
      NProgress.done()
    } else {
      getMenu().then((res: any) => {
        console.log(res)
      })
      // Check whether the user has obtained his permission roles
      // console.log(store.state.user.roles)
      // if (store.state.user.roles == '') {
      //   console.log(1)
      //   try {
      //     // Note: roles must be a object array! such as: ['admin'] or ['developer', 'editor']
      //     await store.dispatch(UserActionTypes.ACTION_GET_USER_INFO, undefined)
      //     const roles = store.state.user.roles
      //     console.log('roles')
      //     console.log(roles)
      //     await store.dispatch(PermissionActionType.ACTION_SET_ROUTES, roles)
      //     // Generate accessible routes map based on role
      //     // store.dispatch(PermissionActionType.ACTION_SET_ROUTES, roles)
      //     // Dynamically add accessible routes
      //     store.state.permission.dynamicRoutes.forEach((route) => {
      //       console.log("dynamicRoutes.forEach", route)
      //       if (route != undefined) {
      //         router.addRoute(route)
      //       }
      //     })
      //     // router.addRoute({ path: '*', redirect: '404' })
      //     // Hack: ensure addRoutes is complete
      //     // Set the replace: true, so the navigation will not leave a history record
      //     console.log("roles....")
      //     console.log(store.state.user.roles)
      //     console.log("准备跳转至原来路径", to.path)
      //     next({ ...to, replace: true })
      //   } catch (err) {
      //     console.log("跳转路由报错", err)
      //     // Remove token and redirect to login page
      //     store.dispatch(UserActionTypes.ACTION_RESET_TOKEN, undefined)
      //     ElMessage.error(err || 'Has Error')
      //     next(`/login?redirect=${to.path}`)
      //     NProgress.done()
      //   }
      // } else {
      //   console.log(2)
      //   next()
      // }
      NProgress.done()
      next()
    }
  } else {
    // Has no token
    if (whiteList.indexOf(to.path) !== -1) {
      // In the free login whitelist, go directly
      next()
    } else {
      // Other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`)
      NProgress.done()
    }
  }
})

router.afterEach((to: RouteLocationNormalized) => {
  // Finish progress bar
  // hack: https://github.com/PanJiaChen/vue-element-admin/pull/2939
  NProgress.done()

  // set page title
  document.title = to.meta.title || 'gea';
})
