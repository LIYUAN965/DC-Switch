import { createStore } from "vuex"
import router from "@/router";
const store = createStore({
    state() {
        return {
            count: 0,
            isCollapse: false,
            currentUser: '未登录',
            currentToken: '',
            userUm: '',
            userRole: '',
            routePath: '/'
        }
    },
    mutations: {
        setIsCollapse(state, payload) {
            state.isCollapse = payload
        },
        setCurrentUser(state, username) {
            state.currentUser = username
        },
        setCurrentToken(state, token) {
            state.currentToken = token
        },
        setUserUm(state, um) {
            state.userUm = um
        },
        setUserRole(state, role) {
            state.userRole = role
        },
        setRoutePath(state, url) {
            state.routePath = url
        },
    }
})

export default store
