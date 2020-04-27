import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

import router from '@/router'
import store from '@/store'
import {
	to,
	setTitle,
} from '@/utils'


router.beforeEach(async (toRoute, fromRoute, next) => {
	NProgress.start()
	setTitle(toRoute.name)
	next()
})

router.afterEach(() => NProgress.done())
