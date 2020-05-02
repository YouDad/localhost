<template lang="pug">
el-container
	el-aside
		el-select(v-model="db")
			el-option(v-for="db in dbs" :key="db" :label="db" :value="db")
		el-button.button(@click="onClick") 加载

	el-main
		router-view
</template>

<script>
import Vue from 'vue'

import {List} from '@/api/db'
import {to} from '@/utils'

export default Vue.extend({
	async created() {
		let [err, ret] = await to(List())
		if (err) {
			console.error(err)
			return
		}

		this.dbs = ret.data
		this.db = this.$route.params.db
		if (!this.db) {
			this.db = this.dbs[0]
		}
	},
	data() {
		return {
			dbs: [],
			db: '',
		}
	},
	methods: {
		onClick() {
			this.$router.push(`/tool2/${this.db}`)
		},
	},
})
</script>

<style lang="scss" scoped>
.button {
	width: 100%;
}
</style>
