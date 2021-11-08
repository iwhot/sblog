const path = require("path");

function resolve(dir) {
	return path.join(__dirname, dir);
}

const name = '后台管理系统'
const port = process.env.port || process.env.npm_config_port || 9527
const uri = process.env.VUE_APP_BASE_URI

module.exports = {
	publicPath: './',
	outputDir: 'dist',
	assetsDir: 'static',
	lintOnSave: false, //process.env.NODE_ENV === 'development',
	transpileDependencies: ["vuetify"],
	productionSourceMap: false,
	devServer: {
		port: port,
		open: true,
		https: false,
		hotOnly: false,
		overlay: {
			warnings: false,
			errors: true
		},
		proxy: {
			'/api/v1': {
				target: uri, //API服务器的地址
				changeOrigin: true,
				pathRewrite: {
					'^/api/v1': '/api/v1'
				}
			},
			'/storage/uploads':{
				target: uri, //API服务器的地址
				changeOrigin: true,
				pathRewrite: {
					'^/storage/uploads': '/storage/uploads'
				}
			}
		}
	},
	configureWebpack: {
		// provide the app's title in webpack's name field, so that
		// it can be accessed in index.html to inject the correct title.
		name: name,
		resolve: {
			alias: {
				'@': resolve('src')
			}
		},
		//关闭编译警告
		performance: {
			hints: false
		}
	},
	chainWebpack(config) {
		// set svg-sprite-loader
		// 第一步：让其他svg loader不要对src/assets/imgs/svgs进行操作
		config.module
			.rule("svg")
			.exclude.add(resolve("src/icons/svg"))
			.end();
		// 第二步：使用svg-sprite-loader 对 src/assets/imgs/svgs下的svg进行操作
		config.module
			.rule("icons")
			.test(/\.svg$/)
			.include.add(resolve("src/icons/svg"))
			.end()
			.use("svg-sprite-loader")
			.loader("svg-sprite-loader")
			//定义规则 使用时 <svg class="icon"> <use xlink:href="#icon-svg文件名"></use>  </svg>
			.options({
				symbolId: "icon-[name]"
			})
			.end();
	}
};
