import { createApp } from 'vue'
import App from './App.vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// import nav from './components/nav.vue'
// import sidBar from './components/sideBar.vue'


// import todoList from './components/todoList.vue'
// import todoItem from './components/todoItem.vue'
// 加载全局样式
import '@/styles/index.scss'
import router from '@/router'
import store from '@/store/index'



const app = createApp(App)
    // app.component("app-nav", nav)
    // app.component("app-sidebar", sidBar)
    // app.component("todo-item", todoItem)


// app.component("todo-list", todoList)

app.use(ElementPlus).use(router).use(store)
app.mount('#root')
