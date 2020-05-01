import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const Page401                 = ()=>import('@/pages/401')
const Page404                 = ()=>import('@/pages/404')
const IndexPage               = ()=>import('@/pages/index')
const Tool1Page               = ()=>import('@/pages/tool1')

export const baseRoutes = [
	{ path: '', redirect: '/index', },
	{ path: '/401', component: Page401, },
	{ path: '/404', component: Page404, },
	{ path: '/index', component: IndexPage, name: '首页', },
	{ path: '/tool1', component: Tool1Page, name: '日志查看工具', },
]

export default new Router({
	mode: 'history',
	routes: baseRoutes,
})
