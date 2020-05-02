import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const Page401                 = ()=>import('@/pages/401')
const Page404                 = ()=>import('@/pages/404')
const IndexPage               = ()=>import('@/pages/index')
const Tool1Page               = ()=>import('@/pages/tool1')
const Tool2Page               = ()=>import('@/pages/tool2')
const Tool2Index              = ()=>import('@/pages/tool2_index')
const Tool2Detail             = ()=>import('@/pages/tool2_detail')

export const baseRoutes = [
	{ path: '', redirect: '/index', },
	{ path: '/401', component: Page401, },
	{ path: '/404', component: Page404, },
	{ path: '/index', component: IndexPage, name: '首页', },
	{ path: '/tool1', component: Tool1Page, name: '日志查看工具', },
	{ path: '/tool2', component: Tool2Page, name: '区块链查看工具', 
		children: [
			{ path: ':db', component: Tool2Index, name: '区块链查看工具--首页' },
			{ path: ':db/:id', component: Tool2Detail, name: '区块链查看工具--详情'}
		],
	},
]

export default new Router({
	mode: 'history',
	routes: baseRoutes,
})
