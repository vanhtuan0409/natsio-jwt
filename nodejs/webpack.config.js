const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: path.resolve(__dirname, "src", "main.js"),
  plugins: [
    new HtmlWebpackPlugin({
      title: "NATS",
      template: path.resolve(__dirname, "index.html"),
    }),
  ],
  devServer: {
    static: "./dist",
  },
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "index.js",
  },
};
