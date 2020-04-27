import Vue from 'vue'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from '@/app'
import router from '@/router'
import store from '@/store'
import '@/router/permission_checker'
import '@/error_log_handler'
import '@/css/main.css'

Vue.use(ElementUI)

globalThis.ELEMENT = ElementUI

new Vue({
	el: '#app',
	router,
	store,
	render: (h) => h(App)
})
