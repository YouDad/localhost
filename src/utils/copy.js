let visited = new Set()

export default function copy(obj, level) {
	if (level === undefined) {
		visited.clear()
		level = 0
	}

	if (!visited.has(obj)) {
		visited.add(obj)
	} else {
		return
	}
	if (typeof obj !== 'object' || !obj) {
		return obj
	}

	let newobj
	switch (obj.constructor) {
		case Date: {
			newobj = new Date(obj)
		} break
		case Array: {
			newobj = []
			for (var i in obj) {
				newobj[i] = typeof obj[i] === 'object' ? copy(obj[i], level+1) : obj[i]
			}
		} break
		case Object: {
			newobj = {}
			for (var i in obj) {
				newobj[i] = typeof obj[i] === 'object' ? copy(obj[i], level+1) : obj[i]
			}
		} break
		default: {
			console.error('[src/utils/copy.js 36]\n', obj, obj.constructor)
		} break
	}

	return newobj
}
