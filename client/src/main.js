import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import axios from 'axios'
import Notifications from 'vue-notification'

Vue.config.productionTip = false

Vue.use(Antd)
Vue.use(Notifications)

axios.post('/api/token', JSON.stringify({
    key: 'matata'
})).then(res => {
    axios.defaults.headers.common['Authorization'] = "Bearer "+res.data.data.token;
    console.info("token", res.data.data.token)

    new Vue({
        router,
        render: function (h) {
            return h(App)
        }
    }).$mount('#app')
}).catch((error) => {
    console.error(error)
});

