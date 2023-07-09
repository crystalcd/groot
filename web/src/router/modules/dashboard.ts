import { type RouteRecordRaw } from 'vue-router'
import Layout from '@/views/layout/Layout.vue'
const routes: RouteRecordRaw[] = [
    {
        path: '/dashboard',
        name: 'dashboard',
        component: Layout,

    }
]

export default routes
