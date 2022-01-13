import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
        path: '/',
        component: () =>
            import ("@/layout/AppLayout.vue"),
        meta: {
            title: '主页'
        },
        children: [{
            name: 'home',
            path: '', // 默认子路由
            component: () =>
                import ('@/views/home/HomeIndex.vue')
        }, {
            name: 'prepareList',
            path: '/switch/preparation',
            component: () =>
                import ("@/components/prepareList.vue")
        }, {
            name: 'prepareItem',
            path: '/switch/preparation/detail/:id',
            component: () =>
                import ("@/components/prepareItem.vue")
        }]
    },
    , {
        path: '/content1',
        component: () =>
            import ("@/components/content1.vue")
    }
]

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})

export {
    routes
}
export default router
