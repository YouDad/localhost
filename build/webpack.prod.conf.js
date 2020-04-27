let baseConfig = require('./webpack.base.conf')
const merge = require('webpack-merge')
const HTMLWebpackPlugin  = require('html-webpack-plugin')
const webpack = require('webpack')
const ExtractPlugin = require('extract-text-webpack-plugin')
const utils = require('./utils')

prodConfig = merge(baseConfig, {
	externals: {
		'axios': 'axios',
		'element-ui': 'ELEMENT',
		'js-cookie': 'Cookies',
		'nprogress': 'NProgress',
		'vue': 'Vue',
		'vue-router': 'VueRouter',
		'vuex': 'Vuex',
	},
	output: {
		filename: '[name].[chunkhash:8].js',
	},
	plugins: [
		new HTMLWebpackPlugin({
			filename: 'index.html',
			template: 'index.cdn.html',
			inject: true,
			minify: {
				removeComments: true,
				collapseWhitespace: true,
				removeAttributeQuotes: true
				// more options:
				// https://github.com/kangax/html-minifier#options-quick-reference
			},
		}),
		new ExtractPlugin('styles.[md5:contentHash:hex:8].css'),
		new webpack.optimize.RuntimeChunkPlugin({
			name: 'vendor',
		}),
		new webpack.optimize.RuntimeChunkPlugin({
			name: 'runtime',
		}),
		new webpack.DefinePlugin({
			'process.env': {
				NODE_ENV: '"production"',
			},
		}),
	]
})

module.exports = prodConfig
