import axios from 'axios'
import {
	encrypt,
	decrypt
} from './aes.js'
import {
	randomString
} from './random.js'
import md5 from 'js-md5'
import router from '@/router'
import {
	Message
} from 'element-ui'

// const Base64 = require('js-base64').Base64

let config = {
	baseURL: process.env.VUE_APP_BASE_API,
	timeout: 50000,
	withCredentials: true
};

const service = axios.create(config)

//请求拦截
service.interceptors.request.use(
	function(config) {
		const token = window.localStorage.getItem('token')
		token && (config.headers.Authorization = token)

		// var tm = Math.round(new Date().getTime() / 1000).toString()
		// config.headers.Timestamp = tm
		// config.headers.Sign = md5(tm)
		// //只有post才参与加密
		// if (config.method == "post" || config.method == "POST") {
		// 	//console.log(config.headers)
		// 	config.headers["content-type"] = "text/plain"
		// 	var iv = md5(tm).substr(0, 16)
		// 	// config.data.timestamp = tm
		// 	config.data = encrypt(JSON.stringify(config.data), process.env.VUE_APP_KEY, iv)

		// }
		return config
	},
	function(error) {
		return Promise.reject(error);
	}
);

//响应拦截
service.interceptors.response.use(
	function(response) {
		//返回参数解密
		if (response.data && response.data.code && (response.data.code == 401)) {
			Message.error(response.data.msg);
			window.localStorage.clear()
			window.sessionStorage.clear()
			window.location.reload()
		}
		return response
	},
	function(error) {
		return Promise.reject(error);
	}
);

export default service
