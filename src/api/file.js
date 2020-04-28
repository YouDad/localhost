import {request} from '@/utils'

export function List() {
	return request({
		url: '/file/list',
		method: 'get',
		data: {},
	})
}

export function Line(filename, line) {
	return request({
		url: `/file/line/${filename}/${line}`,
		method: 'get',
		data: {},
	})
}
