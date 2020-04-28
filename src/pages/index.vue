<template lang="pug">
el-container.index
	el-aside
		el-select.index__select(v-model="file")
			el-option(
				v-for="fileInfo in fileInfos"
				:key="fileInfo.Name"
				:value="fileInfo.Name"
			)
				span.index__select__option-left {{ fileInfo.Name }}
				span.index__select__option-right {{ fileInfo.Size }}

		el-input.index__info(v-for="meta in metas" :value="row ? row[meta.key] : ''")
			template(#prepend) {{ meta.text }}

		el-button(@click="load") 加载

		my-checkboxs(:obj="port")

		el-input(v-model="routine")
			template(#prepend) 筛选线程

		my-checkboxs(:obj="level")

		el-input(v-model="regexp")
			template(#prepend) 筛选内容

	el-main
		el-tabs(v-model="currentTab")
			el-tab-pane(label="未匹配的日志" name="unmatch")
				el-table(:data="unmatchs" border height="80rem")
					el-table-column(prop="line" label="行数" width="70")
					el-table-column(prop="content" label="内容")

			el-tab-pane(label="按行数排序日志" name="lineOrder")
				el-table(
					:data="fileLines" height="80rem"
					:row-class-name="lineLevel"
					border highlight-current-row
					@current-change="onTableChange"
				)
					el-table-column(prop="line" label="行数" width="70")
					el-table-column(prop="port" label="端口" width="70")
					el-table-column(prop="routine" label="线程" width="70")
					el-table-column(prop="content" label="内容")
</template>

<script>
import Vue from 'vue'

import {
	List,
	Line,
} from '@/api/file'
import {to} from '@/utils'

import MyCheckboxs from '@c/checkboxs'

const RE = /(.*\/.*\/.*) (.*:.*:.*\..*)\[(.*)\]\[(.*)\]\[(.*)\]: { (.* .*) } (.*)/

export default Vue.extend({
	components: {
		MyCheckboxs,
	},
	async created() {
		let [err, ret] = await to(List())
		if (err) {
			console.error(err)
			return
		}

		this.fileInfos = ret.data
		this.file = this.fileInfos[0].Name
	},
	watch: {
		file() {
			this.load()
		},
	},
	computed: {
		fileLines() {
			let routine = this.routine
			let ports = this.port.checked
			let levels = this.level.checked
			let re = this.regexp ? RegExp(this.regexp) : undefined
			return this.fileLinesOrderByLine.filter(line=>{
				if (routine) {
					if (routine !== line.routine) {
						return false
					}
				}

				if (re && !re.test(line.content)) {
					return false
				}

				return ports.includes(line.port) && levels.includes(line.level)
			})
		},
	},
	data() {
		return {
			fileInfos: [],
			file: '',
			fileLinesOrderByLine: [],
			unmatchs: [],
			row: {},
			metas: [
				{ text: "行号", key: "line", },
				{ text: "日期", key: "date", },
				{ text: "时间", key: "time", },
				{ text: "端口", key: "port", },
				{ text: "线程", key: "routine", },
				{ text: "来源", key: "source", },
				{ text: "等级", key: "level", },
			],

			currentTab: 'unmatch',

			port: {
				checked: [],
				items: [],
			},

			routine: '',
			regexp: '',

			level: {
				checked: ['DEBUG', 'INFO', 'WARN', 'ERROR', 'TRACE'],
				items: ['DEBUG', 'INFO', 'WARN', 'ERROR', 'TRACE'],
			},
		}
	},
	methods: {
		async load() {
			let length = this.fileLinesOrderByLine.length + this.unmatchs.length
			let [err, ret] = await to(Line(this.file, length))
			if (err) {
				console.error(err)
				return
			}

			for (let line in ret.data) {
				let res = RE.exec(ret.data[line])
				if (!res) {
					this.unmatchs.push({line, content: ret.data[line]})
				} else {
					this.fileLinesOrderByLine.push({
						line,
						date: res[1],
						time: res[2],
						port: res[3],
						routine: res[4],
						level: res[5],
						source: res[6],
						content: res[7],
					})
					if (!this.port.items.includes(res[3])) {
						this.port.items.push(res[3])
						this.port.checked.push(res[3])
					}
				}
			}
		},
		lineLevel({row}) {
			switch(row.level) {
				case 'DEBUG':
					return ['el-alert--info', 'is-light']
				case 'INFO':
					return ''
				case 'WARN':
					return ['el-alert--warning', 'is-light']
				case 'ERROR':
					return ['el-alert--error', 'is-light']
				case 'TRACE':
					return ['el-alert--success', 'is-light']
			}
			return ''
		},
		onTableChange(row) {
			this.row = row
		},
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

	&__info {
		font-size: $font-size--small;
	}
}
</style>
