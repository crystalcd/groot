import { createRouter, createWebHashHistory, type RouteRecordRaw } from 'vue-router'

const modules = import.meta.glob('./modules/**/*.ts', { eager: true })

let routeModuleList: RouteRecordRaw[] = []

// 获取模块路由
Object.values(modules).forEach((key: any) => {
  const mod = key.default || []
  const modList = Array.isArray(mod) ? [...mod] : [mod]
  routeModuleList.push(...modList)
})

const constantRoutes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    redirect: '/dashboard',
    meta: {
      title: 'home'
    }
  },
]

let routes = constantRoutes

routes = [...routeModuleList, ...constantRoutes]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
export { constantRoutes, routeModuleList }
