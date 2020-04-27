const baseConfig = require('./webpack.base.conf')
const merge = require('webpack-merge')
const HTMLWebpackPlugin  = require('html-webpack-plugin')
const webpack = require('webpack')

devConfig = merge(baseConfig, {
	devtool: 'cheap-module-eval-source-map',
	devServer: {
		port:8000,
		host:'0.0.0.0',
		publicPath: '/',
		// proxy: {},
		overlay: {
			errors: true,
		},
		hot: true,
		historyApiFallback: true,
	},
	plugins: [
		new HTMLWebpackPlugin({
			filename: 'index.html',
			inject: true,
		}),
		new webpack.HotModuleReplacementPlugin(),
		new webpack.NoEmitOnErrorsPlugin(),
		new webpack.DefinePlugin({
			'process.env': {
				NODE_ENV: '"development"',
			},
		}),
	],
})

module.exports = devConfig
