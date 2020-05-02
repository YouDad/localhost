import {request} from '@/utils'

export function List() {
	return request({
		url: '/db/list',
		method: 'get',
		data: {},
	})
}

export function GetBlocks(db) {
	return request({
		url: `/db/${db}`,
		method: 'get',
		data: {},
	})
}

export function GetObject(db, id) {
	return request({
		url: `/db/${db}/${id}`,
		method: 'get',
		data: {},
	})
}
