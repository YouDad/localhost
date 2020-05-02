<template lang="pug">
div
	div.block(v-if="block.Hash === $route.params.id")
		el-card
			template(#header)
				div.block__header
					div.block__header--left 区块简介
					div.block__header--right 哈希: 
						router-link(:to="getLink(block.Hash)") {{block.Hash}}
			el-row.block__attr--line
				el-col(v-for="meta in blockMetas.slice(0, 4)" :span="6")
					div.block__attr--key {{meta.label}}
					div.block__attr--value {{block[meta.key]}}
			el-row.block__attr--line
				el-col(v-for="meta in blockMetas.slice(4, 8)" :span="6")
					div.block__attr--key {{meta.label}}
					div.block__attr--value {{block[meta.key]}}
		el-card.block__hash
			template(#header)
				div.block__hash__header
					div.block__hash__header--left 哈希
			div.block__hash__line 上一个区块的哈希:
				router-link(:to="getLink(block.PrevHash)") {{block.PrevHash}}
			div.block__hash__line Merkle树根的哈希:
				span {{block.MerkleRoot}}
			div.block__hash__line 批的Merkle树路径:
				div.block__hash--indent(v-for="path in block.BatchMerklePath")
					| [{{ path.Left?'Left':'Right' }}]{{ path.HashValue }}

		h1.block__txn 一共{{block.TxnNumber}}个交易
		my-txn-card.block__txn(v-for="t in block.Txns" :txn="t")
	div(v-if="txn.Hash === $route.params.id")
		my-txn-card(:txn="txn")
</template>

<script>
import Vue from 'vue'

import {GetObject} from '@/api/db'
import {to} from '@/utils'

import MyTxnCard from './tool2_detail_txn'

export default Vue.extend({
	props: { db: String, },
	components: { MyTxnCard, },
	watch: {
		'$route.params': {
			deep: true,
			immediate: true,
			async handler() {
				let p = this.$route.params
				let [err, ret] = await to(GetObject(p.db, p.id))
				if (err) {
					console.error(err)
					return
				}

				if (ret.data.MerkleRoot) {
					this.block = ret.data
				} else {
					this.txn = ret.data
				}
			},
		}
	},
	data() {
		return {
			txn: {},
			block: {},

			blockMetas: [
				{label: '组', key: 'Group',},
				{label: '高度', key: 'Height',},
				{label: '时间', key: 'Time',},
				{label: '难度', key: 'Target',},
				{label: '组基', key: 'GroupBase',},
				{label: '连锁规模', key: 'BatchSize',},
				{label: '随机数', key: 'Nonce',},
			],
		}
	},
	methods: {
		getLink(id) {
			let path = this.$route.path
			return path.slice(0, path.lastIndexOf('/') + 1) + id
		},
	},
})
</script>

<style lang="scss" scoped>
.block {
	font-size: $font-size--default;

	&__header {
		margin-bottom: 2rem;

		&--left {
			float: left;
		}

		&--right {
			float: right;
		}
	}

	&__hash {
		margin-top: 2rem;

		&__header {
			margin-bottom: 2rem;

			&__left {
				float: left;
			}
		}

		&__line {
			margin-top: 1rem;
			margin-bottom: 1rem;
		}

		&--indent {
			margin-left: 2rem;
		}
	}

	&__attr {

		&--line {
			margin-bottom: 2rem;
		}

		&--key {
			color: $color-grey-light-3;
			font-size: $font-size--small;
		}

		&--value {
		}
	}

	&__txn {
		margin-top: 2rem;
	}
}
</style>
