module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:8081/', //对应自己的接口
                changeOrigin: true,
                ws: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        }
    }
}