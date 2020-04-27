import Vue from 'vue'

import store from '@/store'
import {request} from '@/utils'

// you can set only in production env show the error-log
// if (process.env.NODE_ENV === 'production') {

Vue.config.errorHandler = function(err, vm, info) {
	Vue.nextTick(() => {
		store.dispatch('addErrorLog', {
			err, vm, info,
			url: window.location.href
		})
		console.error('[src/error_log_handler.js 15]\n', err, info)
		request({url: '/log', method: 'post', data: {
			message: `${window.location.href} throw err\ninfo: ${info}\nerr:\n${JSON.stringify(err, null, 4)}`
		}})
	})
}
