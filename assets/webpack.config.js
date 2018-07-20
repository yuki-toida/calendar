const path = require('path');
const webpack = require('webpack');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader');
const Dotenv = require('dotenv-webpack');

module.exports = (_, argv) => ({
  entry: './src/index.js',
  output: {
    filename: 'app.js',
    path: argv.mode == 'development'
      ? path.resolve(__dirname, '../static/js')
      : path.resolve(__dirname, './dist/js')
  },
  // devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'babel-loader',
        options: {
          presets: [['env', { modules: false }]]
        }
      },
      {
        test: /\.css$/,
        use: [
          'style-loader',
          'css-loader',
        ],
      },
      {
        // enforce: 'pre'がついていないローダーより早く処理が実行される
        // babel-loaderで変換する前にコードを検証したいため、指定が必要
        enforce: 'pre',
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'eslint-loader',
        options: {
          fix: true,
          failOnError: true,
        },
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      },
    ]
  },
  resolve: {
    alias: {
      'vue$': 'vue/dist/vue.esm.js'
    }
  },
  plugins: [
    new VueLoaderPlugin(),
    new Dotenv({
      path: argv.mode == 'development'
        ? './env/local.env'
        : './env/dev.env',
    }),
    new CopyWebpackPlugin([
      {
        from: 'static',
        to: argv.mode == 'development'
          ? path.resolve(__dirname, '../static')
          : path.resolve(__dirname, './dist'),
      }
    ]),
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery'
    }),    
  ],
});