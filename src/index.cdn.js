import App from '@/app'
import router from '@/router'
import store from '@/store'
import '@/router/permission_checker'
import '@/error_log_handler'
import '@/css/main.css'
import PopConfirm from '@c/popconfirm'

PopConfirm.install(Vue)

new Vue({
	el: '#app',
	router,
	store,
	render: (h) => h(App)
})
