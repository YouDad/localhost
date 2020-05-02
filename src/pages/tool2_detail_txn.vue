<template lang="pug">
el-card.txn
	template(#header)
		div.txn__header
			i.txn__header__left.el-icon-sort.txn__icon
			span.txn__header__left {{ inputSum }} COIN
			span.txn__header__right Hash: 
				router-link(:to="getLink(txn.Hash)") {{ txn.Hash }}
	div.txn__body
		div.txn__body__left
			div.txn__item(v-for="item in txn.Vin")
				router-link.txn__item__left(:to="getLink(item.PubKeyHash)") {{ item.PubKeyHash }}
				span.txn__item__right - {{ item.Value }}
		div.txn__body__mid
			i.txn__body__mid__icon.el-icon-sort-up
		div.txn__body__right
			div.txn__item(v-for="item in txn.Vout")
				router-link.txn__item__left(:to="getLink(item.PubKeyHash)") {{ item.PubKeyHash }}
				span.txn__item__right + {{ item.Value }}
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
	name: 'MyTxnCard',
	props: { txn: Object, },
	computed: {
		inputSum() {
			let sum = 0
			for (let vin of this.txn.Vin) {
				sum += vin.VoutValue
			}
			return sum
		},
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
.txn {
	font-size: $font-size--default;

	&__icon {
		transform: rotate(90deg) rotateY(180deg);
		font-size: $font-size--big;
		margin-right: 1rem;
	}

	&__header {
		margin-bottom: 1rem;

		&__left {
			float: left;
		}

		&__right {
			float: right;
		}
	}

	&__body {
		margin-bottom: 2rem;

		&__left {
			float: left;
			width: 40%;
			color: red;
		}

		&__mid {
			float: left;
			width: 20%;
			text-align: center;

			&__icon {
				transform: rotate(90deg) rotateY(180deg);
				font-size: $font-size--big;
			}
		}

		&__right {
			float: right;
			width: 40%;
			color: green;
		}
	}

	&__item {
		display: inline-block;
		width: 100%;

		&__left {
			float: left;
		}

		&__right {
			float: right;
		}
	}
}
</style>
