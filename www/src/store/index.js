import { createStore } from "vuex"
const store = createStore({
    state() {
        return {
            count: 0,
            isCollapse: false
        }
    },
    mutations: {
        setIsCollapse(state, payload) {
            state.isCollapse = payload
        },
    }
})
export default store