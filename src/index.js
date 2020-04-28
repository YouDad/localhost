import Vue from 'vue'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)
globalThis.ELEMENT = ElementUI

import App from '@/app'
import router from '@/router'
import store from '@/store'
import '@/router/permission_checker'
import '@/css/main.css'

new Vue({
	el: '#app',
	router,
	store,
	render: (h) => h(App)
})
