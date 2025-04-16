const path = require('path');

let entryPoints = {
    landing: "./src/pages/landing/index.jsx"
}

module.exports = {
  entry: entryPoints,
    mode: "development",
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
      clean: true,
  },
module: {
    rules: [
      {
        // Pass application JS/TS files through babel-loader,
        // transpiling to targets defined in browserslist
        test: /\.([jt]sx?|[cm]js)$/,
        // Only compile application files and specific dependencies
        // (other npm and vendored dependencies must be in ES5 already)
        exclude: [
          /node_modules\/(?!(react-dnd|chart\.js|@uppy|pdfjs-dist|react-resizable-panels)\/)/
        ],
        use: [
          {
            loader: 'babel-loader',
            options: {
              // Configure babel-loader to cache compiled output so that
              // subsequent compile runs are much faster
              cacheDirectory: true,
        //      configFile: path.join(__dirname, './babel.config.json'),
              plugins: [
                process.env.REACT_REFRESH_ENABLED === 'true' &&
                  'react-refresh/babel',
              ].filter(Boolean),
            },
          },
        ],
        type: 'javascript/auto',
      }
    ]
},
 resolve: {
    extensions: ['.js', '.jsx']
  },


};
