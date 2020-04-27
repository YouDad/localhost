import to from './to'
export {default as to} from './to'
export {default as copy} from './copy'
export {default as request} from './request'

export function setTitle(title) {
	if (!title) {
		title = ''
	}
	window.document.title = `${title} -- System`
}

export function parseObjsByMetas(metas, resArr, refArr) {
	if (!resArr) {
		return
	}
	for (var i of resArr) {
		let obj = {}
		for (var meta of metas) {
			obj[meta.key] = i[meta.key]
		}
		refArr.push(obj)
	}
}
