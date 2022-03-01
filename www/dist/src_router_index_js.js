"use strict";
(self["webpackChunkconsole_service"] = self["webpackChunkconsole_service"] || []).push([["src_router_index_js"],{

/***/ "./src/router/index.js":
/*!*****************************!*\
  !*** ./src/router/index.js ***!
  \*****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "routes": () => (/* binding */ routes),
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var vue_router__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vue-router */ "webpack/sharing/consume/default/vue-router/vue-router");
/* harmony import */ var vue_router__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(vue_router__WEBPACK_IMPORTED_MODULE_0__);


const routes = [{
        path: '/',
        component: () =>
            Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_css-loader_dist_runtime_api_js-node_modules_css-loader_dist_runtime_sour-865ddd"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_71be"), __webpack_require__.e("src_layout_AppLayout_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/layout/AppLayout.vue */ "./src/layout/AppLayout.vue")),
        meta: {
            title: '主页'
        },
        children: [{
            name: 'home',
            path: '', // 默认子路由
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("src_views_home_HomeIndex_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/views/home/HomeIndex.vue */ "./src/views/home/HomeIndex.vue"))
        }, {
            name: 'prepareList',
            path: '/switch/preparation',
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_axios_index_js-node_modules_css-loader_dist_runtime_api_js-node_modules_-bd071c"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("webpack_sharing_consume_default_element-plus_element-plus"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_2abd"), __webpack_require__.e("src_components_prepareList_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/prepareList.vue */ "./src/components/prepareList.vue"))
        }, {
            name: 'prepareItem',
            path: '/switch/preparation/detail/:id',
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_axios_index_js-node_modules_css-loader_dist_runtime_api_js-node_modules_-bd071c"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("webpack_sharing_consume_default_element-plus_element-plus"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_2abd"), __webpack_require__.e("src_components_prepareItem_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/prepareItem.vue */ "./src/components/prepareItem.vue"))
        }, {
            name: 'drillList',
            path: '/switch/drill',
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_axios_index_js-node_modules_css-loader_dist_runtime_api_js-node_modules_-bd071c"), __webpack_require__.e("vendors-node_modules_echarts_lib_chart_pie_install_js-node_modules_echarts_lib_component_grid-1c924e"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("src_components_drillList_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/drillList.vue */ "./src/components/drillList.vue"))
        }, {
            name: 'drillItem',
            path: '/switch/drill/detail/:id',
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_axios_index_js-node_modules_css-loader_dist_runtime_api_js-node_modules_-bd071c"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("webpack_sharing_consume_default_element-plus_element-plus"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_2abd"), __webpack_require__.e("src_components_drillItem_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/drillItem.vue */ "./src/components/drillItem.vue"))
        }, {
            name: 'rangeItem',
            path: '/switch/range/:id',
            component: () =>
                Promise.all(/*! import() */[__webpack_require__.e("vendors-node_modules_axios_index_js-node_modules_css-loader_dist_runtime_api_js-node_modules_-bd071c"), __webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("webpack_sharing_consume_default_element-plus_element-plus"), __webpack_require__.e("src_components_rangeItem_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/rangeItem.vue */ "./src/components/rangeItem.vue"))
        }]
    },
    , {
        path: '/content1',
        component: () =>
            Promise.all(/*! import() */[__webpack_require__.e("webpack_sharing_consume_default_vue_vue-_f05f"), __webpack_require__.e("src_components_content1_vue")]).then(__webpack_require__.bind(__webpack_require__, /*! @/components/content1.vue */ "./src/components/content1.vue"))
    }
]

const router = (0,vue_router__WEBPACK_IMPORTED_MODULE_0__.createRouter)({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: (0,vue_router__WEBPACK_IMPORTED_MODULE_0__.createWebHashHistory)(),
    routes, // `routes: routes` 的缩写
})


/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (router);


/***/ })

}]);
//# sourceMappingURL=src_router_index_js.js.map