import {
	getToken,
	getUserInfo
} from "@/utils/auth.js"

const state = {
	websock: null,
	url: '',
	lockReconnect: false, //是否真正建立连接
	timeout: 25 * 1000, //25秒一次心跳
	timeoutObj: null, //心跳心跳倒计时
	serverTimeoutObj: null, //心跳倒计时
	timeoutnum: null, //断开 重连倒计时
	message: {},
}

const mutations = {
	WEBSOCKET_INIT(state, url) {
		var that = this;
		state.websock = new WebSocket(url);
		state.url = url;
		state.websock.onopen = function() {
			console.log("ws连接成功") //发送用户JWT令牌 后端解析后自动绑定用户
			//开始发送登陆
			var userInfo = getUserInfo()
			var msgJson = {
				type: 'login',
				user_id: userInfo.id,
				state: 1
			}
			state.websock.send(JSON.stringify(msgJson))
			// state.websock.send('OpenBarScanner')
			//发送心跳包
			that.commit('websocket/start')
		}
		state.websock.onmessage = function(callBack) {
			
			//重置心跳
			//console.log(callBack.data)
			that.commit('websocket/reset')
			
			if (callBack.data == '{"type":"ping"}') {
				console.log('心跳检测ing...')
				return;
			}
			state.message = callBack.data
		}
		state.websock.οnerrοr = function() { //e错误
			// console.log(e)
			that.commit('websocket/reconnect')
		}
		state.websock.onclose = function(e) { //e关闭
			that.commit('websocket/reconnect')
		}
	},
	WEBSOCKET_SEND(state, message) {
		state.websock.send(message);
	},
	reconnect(state) { //重新连接
		// console.log("重新连接")
		var that = this;
		if (state.lockReconnect) {
			return;
		}
		state.lockReconnect = true;
		//没连接上会一直重连，设置延迟避免请求过多
		state.timeoutnum && clearTimeout(state.timeoutnum);
		state.timeoutnum = setTimeout(function() {
			//新连接
			that.commit('websocket/WEBSOCKET_INIT', state.url)
			state.lockReconnect = false;
		}, 5000);
	},
	reset(state) { //重置心跳
		//清除时间
		clearTimeout(state.timeoutObj);
		clearTimeout(state.serverTimeoutObj);
		//心跳
		this.commit('websocket/start')
	},
	start(state) { //开启心跳
		// console.log("开启心跳")
		var that = this;
		state.timeoutObj && clearTimeout(state.timeoutObj);
		state.serverTimeoutObj && clearTimeout(state.serverTimeoutObj);
		state.timeoutObj = setTimeout(function() {
			//这里发送一个心跳，后端收到后，返回一个心跳消息，
			// console.log(state.websock)
			if (state.websock.readyState == 1) { //如果连接正常
				state.websock.send('{"type":"ping"}');
				return true;
			} else { //否则重连
				that.commit('websocket/reconnect')
			}
			
			state.serverTimeoutObj = setTimeout(function() {
				//超时关闭
				state.websock.close();
			}, state.timeout);

		}, state.timeout)
	}

}


const actions = {
	WEBSOCKET_INIT({
		commit
	}, url) {
		commit('WEBSOCKET_INIT', url)
	},
	WEBSOCKET_SEND({
		commit
	}, message) {
		commit('WEBSOCKET_SEND', message)
	},
}

export default {
	namespaced: true,
	state,
	mutations,
	actions
}
