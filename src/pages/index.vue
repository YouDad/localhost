<template lang="pug">
div.index
	router-link.index__tool(to="/tool1") 日志查看工具
	router-link.index__tool(to="/tool2") 区块链查看工具
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
})
</script>

<style lang="scss" scoped>
.index {
	padding: 2rem;

	&__tool {
		display: block;
		font-size: $font-size--default;
	}
}
</style>
