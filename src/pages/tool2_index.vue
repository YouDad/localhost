<template lang="pug">
el-table(:data="blocks" show-summary)
	el-table-column(label="高度")
		template(#default="{row}")
			router-link(:to="`${$route.path}/${row.Height}`") {{ row.Height }}
	el-table-column(label="哈希")
		template(#default="{row}")
			router-link.ellipsis(:to="`${$route.path}/${row.Hash}`") {{ row.Hash }}
	el-table-column(
		v-for="meta in blockMeta"
		:prop="meta.prop"
		:label="meta.label"
	)
</template>

<script>
import Vue from 'vue'

import {GetBlocks} from '@/api/db'
import {to} from '@/utils'

export default Vue.extend({
	watch: {
		'$route.params': {
			immediate: true,
			deep: true,
			handler() {
				this.load()
			},
		},
	},
	data() {
		return {
			blocks: [],

			blockMeta: [
				{prop: 'Time', label: '时间戳',},
				{prop: 'TxnNumber', label: '交易数量',},
				{prop: 'Target', label: '难度',},
				{prop: 'GroupBase', label: 'GroupBase',},
				{prop: 'BatchSize', label: 'BatchSize',},
			],
		}
	},
	methods: {
		async load() {
			let [err, ret] = await to(GetBlocks(this.$route.params.db))
			if (err) {
				console.error(err)
				return
			}

			let blocks = []
			let keys = Object.keys(ret.data).reverse()
			for (let key of keys) {
				blocks.push(ret.data[key])
			}
			this.blocks = blocks
		},
	},
})
</script>

<style lang="scss" scoped>
.ellipsis {
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
</style>
