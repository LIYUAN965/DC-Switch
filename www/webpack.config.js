const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader/');
const ModuleFederationPlugin =
    require("webpack").container.ModuleFederationPlugin;

module.exports = {
    devtool: 'source-map',
    mode: 'development',
    entry: './src/main.js',
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, 'dist')
    },
    devServer: {
        port: 8000,
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
            "Access-Control-Allow-Headers": "X-Requested-With, content-type, Authorization"
        }
    },
    resolve: {
        alias: {
            '@': path.join(__dirname, 'src') // 绝对路径
        }
    },
    module: {
        rules: [{
                test: /\.mjs$/,
                resolve: { byDependency: { esm: { fullySpecified: false } } }
            },
            {
                test: /\.vue$/,
                use: [
                    'vue-loader'
                ]
            },
            {
                test: /\.(sa|sc|c)ss$/,
                use: [
                    // 将 JS 字符串生成为 style 节点
                    'style-loader',
                    // 将 CSS 转化成 CommonJS 模块
                    'css-loader',
                    // 将 Sass 编译成 CSS
                    'sass-loader',
                ],
            },
        ]
    },
    plugins: [
        new VueLoaderPlugin(),
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, './public/index.html'),
            filename: 'index.html',
            title: 'console-service'
        }),
        new ModuleFederationPlugin({
            name: "consoleback",
            filename: "remoteEntry.js",
            exposes: {
                "./RemotePage": "./src/App.vue",
                "./routes": "./src/router/index.js"
            },
            shared: {
                vue: {
                    singleton: true
                },
                'element-plus': {
                    singleton: true
                },
                'vue-router': {
                    singleton: true
                }
            }
        })
    ]
}
