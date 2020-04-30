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

		el-button.index__button(@click="load") 加载

		my-checkboxs(:obj="port")

		el-input(v-model="routine")
			template(#prepend) 筛选线程

		my-checkboxs(:obj="level")

		el-input(v-model="regexp")
			template(#prepend) 筛选内容

		el-button.index__button(@click="onFilterClick") 过滤

	el-main
		el-tabs(v-model="currentTab")
			el-tab-pane(label="未匹配的日志" name="unmatch")
				el-table(:data="unmatchs" border height="80rem")
					el-table-column(prop="line" label="行数" width="70")
					el-table-column(prop="content" label="内容")

			el-tab-pane(label="按行数排序日志" name="lineOrder")
				div.index__log-table
					table.index__log-table__head
						tr
							th.index__log-table__head__param 行数
							th.index__log-table__head__param 端口
							th.index__log-table__head__param 线程
							th.index__log-table__head__content 内容
					div.index__log-table__body__wrapper
						table.index__log-table__body
							tr(
								v-for="(line, index) in fileLines"
								:class="lineLevel(line)"
								@click="onTableChange(index)"
							)
								td.index__log-table__body__param {{ line.line }}
								td.index__log-table__body__param {{ line.port }}
								td.index__log-table__body__param {{ line.routine }}
								td.index__log-table__body__content {{ line.content }}
</template>

<script>
import Vue from 'vue'

import {
	List,
	Line,
} from '@/api/file'
import {to} from '@/utils'

import MyCheckboxs from '@c/checkboxs'

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
	data() {
		return {
			fileInfos: [],
			file: '',

			fileLinesOrderByLine: [],
			fileLines: [],
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
				checked: ['1101', '1102', '1103', '1111', '2201', '2202', '2203', '2222'],
				items: ['1101', '1102', '1103', '1111', '2201', '2202', '2203', '2222'],
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

			this.fileLinesOrderByLine = this.fileLinesOrderByLine.concat(ret.data.match)
			this.unmatchs = this.unmatchs.concat(ret.data.unmatch)
		},
		lineLevel(row) {
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
		onTableChange(index) {
			this.row = this.fileLines[index]
		},
		onFilterClick() {
			let routine = this.routine ? RegExp(this.routine) : undefined
			let ports = this.port.checked
			let levels = this.level.checked
			let re = this.regexp ? RegExp(this.regexp) : undefined
			this.fileLines = this.fileLinesOrderByLine.filter(line=>{
				if (routine && !routine.test(String(line.routine))) {
					return false
				}

				if (re && !re.test(line.content)) {
					return false
				}

				return ports.includes(line.port) && levels.includes(line.level)
			})
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

	&__button {
		width: 100%;
	}

	&__log-table {
		font-size: $font-size--default;
		word-break: break-all;

		&__head {
			width: 100%;

			&__param {
				width: 7rem;

				border-top: 1px solid $color-grey-light-3;
				border-left: 1px solid $color-grey-light-3;
				border-bottom: 1px solid $color-grey-light-3;
			}

			&__content {
				width: calc(100% - 21rem);

				border-top: 1px solid $color-grey-light-3;
				border-left: 1px solid $color-grey-light-3;
				border-bottom: 1px solid $color-grey-light-3;
				border-right: 1px solid $color-grey-light-3;
			}
		}

		&__body {

			&__wrapper {
				height: 77rem;
				overflow-y: auto;
			}

			&__param {
				text-align: center;
				width: 7rem;
				padding: 1rem;

				border-left: 1px solid $color-grey-light-3;
				border-bottom: 1px solid $color-grey-light-3;
			}

			&__content {
				padding: 1rem;

				border-left: 1px solid $color-grey-light-3;
				border-right: 1px solid $color-grey-light-3;
				border-bottom: 1px solid $color-grey-light-3;
			}
		}
	}
}
</style>
