const rootDir = `${__dirname}/dist/`;

module.exports = [
  {
    // メインとなるJavaScriptファイル（エントリーポイント）
    entry: `./src/index.tsx`,

    // ファイルの出力設定
    output: {
      //  出力ファイルのディレクトリ名
      path: rootDir + `/js/`,
      // 出力ファイル名
      filename: 'main.js',
    },
    module: {
      rules: [
        {
          // 拡張子 .js .jsx の場合
          test: /\.(js|jsx)$/,
          use: [
            {
              // Babel を利用する
              loader: 'babel-loader',
              // Babel のオプションを指定する
              options: {
                presets: [
                  ['@babel/react'],
                  [
                    '@babel/env',
                    {
                      targets: {
                        ie: 11,
                      },
                      useBuiltIns: 'usage',
                      corejs: 3,
                    },
                  ],
                  '@babel/preset-typescript',
                ],
                plugins: ['@babel/plugin-proposal-class-properties'],
              },
            },
          ],
          // node_modules は除外する
          exclude: /node_modules/,
        },
        {
          // 拡張子 .ts もしくは .tsx の場合
          test: /\.(ts|tsx)$/,
          // TypeScript をコンパイルする
          use: 'ts-loader',
          exclude: '/node_modules/',
        },
        {
          test: /\.(scss)$/,
          use: [
            {
              loader: 'style-loader', // inject CSS to page
            },
            {
              loader: 'css-loader', // translates CSS into CommonJS modules
            },
            {
              loader: 'postcss-loader', // Run post css actions
              options: {
                plugins: () => {
                  // post css plugins, can be exported to postcss.config.js
                  return [require('precss'), require('autoprefixer')];
                },
              },
            },
            {
              loader: 'sass-loader', // compiles Sass to CSS
            },
          ],
        },
      ],
    },
    // import 文で .ts や .tsx ファイルを解決するため
    resolve: {
      extensions: ['.ts', '.tsx', '.js', 'jsx', '.json'],
    },
    //ソースマップはbuild時削除
    devtool: 'inline-source-map',
    // performance: { hints: false },
    devServer: {
      contentBase: rootDir,
      publicPath: '/js/',
      disableHostCheck: true,
      open: true,
      // watchContentBase: true,
      port: 8080,
      host: '0.0.0.0',
    },
  },
];
