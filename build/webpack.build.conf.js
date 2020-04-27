const baseConfig = require('./webpack.base.conf')
const merge = require('webpack-merge')
const HTMLWebpackPlugin  = require('html-webpack-plugin')
const webpack = require('webpack')
const ExtractPlugin = require('extract-text-webpack-plugin')

prodConfig = merge(baseConfig, {
	devtool: 'cheap-module-eval-source-map',
	output: {
		filename: '[name].js',
	},
	plugins: [
		new HTMLWebpackPlugin({
			filename: 'index.html',
			template: 'index.html',
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
				// NODE_ENV: '"production"',
				NODE_ENV: '"development"',
			},
		}),
	]
})

module.exports = prodConfig
