import { createRouter, createWebHistory, RouteRecordRaw, createWebHashHistory } from 'vue-router'
import Layout from '../layout/Index.vue'

export const routes: Array<RouteRecordRaw> = [
    {
      path: '/',
      component: Layout,
      redirect: '/dashboard',
      children: [
          {
              path: 'dashboard',
              component: () => import('../views/dashboard/Index.vue'),
              name: 'Dashboard',
              meta: {
                  title: 'dashboard',
              }
          }
      ],
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

export default router;