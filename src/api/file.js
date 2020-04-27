import {request} from '@/utils'

export function List() {
	return request({
		url: '/file/list',
		method: 'get',
		data: {},
	})
}
