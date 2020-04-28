<template lang="pug">
div
	el-checkbox(:indeterminate="indeterminate" v-model="all" @change="onCheckAll") 全选
	el-checkbox-group(v-model="obj.checked" @change="onCheckChange")
		el-checkbox(v-for="item in obj.items" :label="item" :key="item") {{ item }}
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
	name: 'MyCheckboxs',
	props: {
		obj: Object,
	},
	data() {
		return {
			indeterminate: false,
			all: true,
		}
	},
	methods: {
		onCheckAll(val) {
			this.obj.checked.splice(0, this.obj.checked.length)
			if (val) {
				for (let item of this.obj.items) {
					this.obj.checked.push(item)
				}
			}
			this.indeterminate = false
		},
		onCheckChange(val) {
			let cnt = val.length
			this.all = cnt === this.obj.items.length
			this.indeterminate = 0 < cnt && cnt < this.obj.items.length
		},
	},
})
</script>

<style lang="scss" scoped>

</style>
