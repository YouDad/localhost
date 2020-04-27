<template lang="pug">
div
	h1.status 401
	div.content
		ul.suggestions
			li.suggest
				strong.suggest_title 不好，你好像没有权限访问这个网页
	el-button.back-button(@click="backToIndex" type="warning") 返回首页
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
	data() {
		return {
			location,
		}
	},
	methods: {
		backToIndex() {
			location.assign('/')
		}
	},
})
</script>

<style lang="scss">
* {
	box-sizing: border-box;
	-webkit-tap-highlight-color: transparent;
}

::selection {
	background-color: #fff;
	color: #3e9ce0;
}

@font-face {
	font-family: 'Nunito';
	font-style: normal;
	font-weight: 900;
	font-display: swap;
	src: local('Nunito Black'), local('Nunito-Black'), url(https://shadow.elemecdn.com/lib/404@1.0.0/NunitoBlack.woff2) format('woff2');
	unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191,
	U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

body,
html {
	position: relative;
	height: 100%;
	margin: 0;
	padding: 0;
	overflow: hidden;
	background-image: linear-gradient(135deg, #9cc3ed 0%, #4c9ee6 100%);
	background-size: cover;
	background-repeat: no-repeat;
	background-color: #73b0e9;
	color: #e9f1f9;
	font-family: -apple-system, 'Noto Sans', 'Helvetica Neue', Helvetica, 'Nimbus Sans L', Arial, 'Liberation Sans', 'PingFang SC',
	'Hiragino Sans GB', 'Noto Sans CJK SC', 'Source Han Sans SC', 'Source Han Sans CN', 'Microsoft YaHei', 'Wenquanyi Micro Hei',
	'WenQuanYi Zen Hei', 'ST Heiti', SimHei, 'WenQuanYi Zen Hei Sharp', sans-serif;
}

.man,
.status {
	position: absolute;
}

.status {
	font-size: 480px;
	right: 100px;
	bottom: 10px;
	margin: 0;
	padding: 0;
	line-height: 1;
	user-select: none;
	background-image: linear-gradient(to right, rgba(136, 186, 235, 0.678) 0%, rgb(94, 165, 228) 100%);
	background-clip: text;
	-webkit-background-clip: text;
	color: transparent;
	font-family: 'Nunito';
}

.logo {
	display: inline-block;
	height: 20px;
	vertical-align: middle;
	margin-right: 7px;
}

.man {
	z-index: 2;
	width: 25%;
	min-width: 250px;
	left: 5%;
	bottom: 70px;
}

.content {
	position: relative;
	z-index: 3;
	padding-left: 5%;
	margin-top: 80px;
}

.indicator {
	display: inline-block;
	font-size: 0;
	padding: 10px 25px;
	padding-left: 20px;
	margin-bottom: 40px;
	border-radius: 7px;
	color: #f2f8ff;
	background: rgba(255, 255, 255, 0.18);
	box-shadow: 7px 18px 31px -17px rgba(81, 161, 230, 0.5);
}

.indicator > span {
	font-size: 20px;
	vertical-align: middle;
}

.suggestions {
	margin: 0;
	padding-left: 0;
	list-style: none;
}

.suggest {
	margin-bottom: 50px;
}

.suggest_title {
	position: relative;
	display: block;
	margin-bottom: 20px;
	font-size: 20px;
	font-weight: normal;
}

.url {
	text-decoration: underline;
	font-size: 16px;
	color: #e9f1f9;
}

@media only screen and (max-width: 950px) {
	body {
		height: 100vh;
		width: 100vw;
	}

	.status {
		font-size: 400px;
		right: -15%;
	}

	.man {
		width: 50%;
		left: 10%;
		bottom: 5px;
	}
}

@media only screen and (max-width: 550px) {
	.content {
		margin-top: 5vh;
		padding-left: 8%;
	}

	.indicator {
		margin-bottom: 5vh;
	}

	.indicator > span {
		font-size: 13px;
	}

	.suggest_title {
		font-weight: bold;
		font-size: 13px;
		margin-bottom: 3vh;
	}

	.url {
		padding-right: 5vw;
		word-break: break-all;
	}

	.suggest_title::before {
		left: -13px;
		width: 6px;
		height: 6px;
		margin-top: -3px;
	}

	.status {
		font-size: 300px;
		bottom: 10px;
		right: -10%;
	}

	.man {
		width: 50%;
		left: 10%;
	}

	.url {
		font-weight: lighter;
	}
}

.back-button {
	margin-left: 5%;
}
</style>
