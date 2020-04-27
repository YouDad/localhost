import axios from 'axios'

import store from '@/store'

const service = axios.create({
	baseURL: process.env.BASE_API,
	timeout: 5000
})

service.interceptors.request.use(request => {
	console.debug('[src/utils/request.js 15]\n', 'service.interceptors.request', request)
	if (!request.data) {
		request.data = {}
	}
	return request
}, error => {
	console.error('[src/utils/request.js 17]\n', 'service.interceptors.request', error)
	Promise.reject(error)
})

service.interceptors.response.use(response => {
	let res = response.data
	if (res) {
		console.debug('[src/utils/request.js 24]\n', 'service.interceptors.response', response)
	} else {
		console.error('[src/utils/request.js 26]\n', 'service.interceptors.response', response, response.config)
		return res
	}

	if (res.message) {
		ELEMENT.Message({
			message: res.message,
			type: res.code === 0 ? 'success' : 'error',
			duration: 5 * 1000,
		})
	}

	if (res.code === 0) {
		delete res.code
		delete res.message
		return res
	}
	return Promise.reject(res)
}, error => {
	console.error('[src/utils/request.js 45]\n', 'service.interceptors.response', error)
	ELEMENT.Message({
		message: error.message,
		type: 'error',
		duration: 5 * 1000,
	})
	return Promise.reject(error)
})

export default service
