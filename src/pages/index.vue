<template lang="pug">
div.index
	el-select.index__select(v-model="file")
		el-option(
			v-for="fileInfo in fileInfos"
			:key="fileInfo.Name"
			:value="fileInfo.Name"
		)
			span.index__select__option-left {{ fileInfo.Name }}
			span.index__select__option-right {{ fileInfo.Size }}
</template>

<script>
import Vue from 'vue'

import {List} from '@/api/file'
import {to} from '@/utils'

export default Vue.extend({
	async created() {
		let [err, ret] = await to(List())
		if (err) {
			console.error(err)
			return
		}

		this.fileInfos = ret.data
	},
	data() {
		return {
			fileInfos: [],
			file: {},
		}
	},
})
</script>

<style lang="scss" scoped>
.index {

	&__select {

		&__option {

			&-left {
				float: left;
				font-size: $font-size--default;
			}

			&-right {
				float: right;
				color: $color-grey-light-3;
				font-size: $font-size--small;
			}
		}
	}
}
</style>
