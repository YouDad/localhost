const utils = require('./utils')
const VueLoaderPlugin = require('vue-loader/lib/plugin')
const ExtractPlugin = require('extract-text-webpack-plugin')

let baseConfig = {
	mode: process.env.NODE_ENV,
	entry: {
		app: utils.resolve('src', 'index.js'),
	},
	output: {
		path: utils.resolve('dist'),
		filename: '[name].js',
		publicPath: '/',
	},
	resolve: {
		extensions: ['.js', '.vue', '.json'],
		alias: {
			'@': utils.resolve('src'),
			'@c': utils.resolve('src', 'components'),
		},
	},
	module: {
		rules: [
			/*
			* TODO: ESLint
			* */
			{
				test: /\.vue$/,
				loader: 'vue-loader',
				/*
				* TODO: VueLoaderOptions
				* */
			},
			{
				test: /\.pug$/,
				loader: 'pug-plain-loader',
				/*
				* TODO: PugLoaderOptions
				* */
			},
			{
				test: /\.css$/,
				use: ['style-loader', 'css-loader'],
			},
			// FIXME: consider about resource file path
			{
				test: /\.(png|jpg|gif|svg)$/,
				loader: 'file-loader',
				options: {
					name: 'src/image/[name].[ext]?[hash]',
				},
			},
			{
				test: /\.(eot|svg|ttf|woff|woff2)(\?\S*)?$/,
				loader: 'file-loader',
			},
			{
				test: /\.scss$/,
				use: ExtractPlugin.extract({
					fallback: 'style-loader',
					use: [
						'style-loader',
						'css-loader',
						'sass-loader',
						{
							loader: 'sass-resources-loader',
							options: {
								resources: utils.resolve('src', 'css', 'main.scss'),
							},
						},
					],
				})
			},
		],
	},
	plugins: [
		new VueLoaderPlugin(),
	],
}

module.exports = baseConfig
