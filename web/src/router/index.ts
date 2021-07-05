import {createRouter, createWebHistory, RouteRecordRaw, createWebHashHistory} from 'vue-router'
import Layout from '../layout/Index.vue'

// export const routes: Array<RouteRecordRaw> = [
//     {
//       path: '/',
//       component: Layout,
//       redirect: '/dashboard',
//       children: [
//           {
//               path: 'dashboard',
//               component: () => import('../views/dashboard/Index.vue'),
//               name: 'Dashboard',
//               meta: {
//                   title: 'dashboard',
//               }
//           }
//       ],
//     },
// ]


const baseRouters: Array<RouteRecordRaw> = [
    {
        path: '/redirect',
        component: Layout,
        meta: {hidden: true},
        children: [
            {
                path: '/redirect/:path(.*)',
                component: () => import(/* webpackChunkName: "redirect" */ '../views/redirect/Index.vue')
            }
        ]
    },
    {
        path: '/',
        redirect: '/dashboard1'
    },
    {
        path: '/login',
        name: 'login',
        component: () => import (/* webpackChunkName: "redirect" */ '../views/login/Index.vue')
    },
    {
        path: '/dashboard1',
        component: () => import(/* webpackChunkName: "redirect" */ '../views/dashboard/Index.vue'),
        name: 'dashboard1',
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes: baseRouters,
})

// router.beforeEach(async(to: RouteLocationNormalized, _: RouteLocationNormalized, next: any) => {
//     console.log(to)
//     next()
// })

export default router
