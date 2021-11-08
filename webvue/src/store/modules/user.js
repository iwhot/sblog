import {
	getToken,
	getUserInfo
} from "@/utils/auth.js"
import {
	userInfo
} from '@/api/common/index.js'

//刷新用户
function refreshUser() {
	userInfo({}).then(res => {
		window.localStorage.removeItem("user")
		window.localStorage.setItem("user", JSON.stringify(res.data.data))
	})
}

const state = {
	userInfo: getUserInfo(),
	token: getToken()
}

const mutations = {
	SET_USER_INFO(state, user) {
		var userInfo2 = JSON.stringify(user)
		window.localStorage.removeItem("user")
		state.userInfo = user
		window.localStorage.setItem("user", userInfo2)
	},
	SET_TOKEN(state, token) {
		window.localStorage.removeItem("token")
		state.token = token
		window.localStorage.setItem("token", token)
	},
	LOGOUT(state) {
		window.localStorage.clear()
		window.sessionStorage.clear()
		state.userInfo = []
		state.token = ""
		window.location.reload()
	},
	REFRESH_USER(state) {
		refreshUser()
	}
}


const actions = {
	set_user_info({
		commit
	}, user) {
		commit('SET_USER_INFO', user)
	},
	set_token({
		commit
	}, token) {
		commit('SET_TOKEN', token)
	},
	logout({
		commit
	}) {
		commit('LOGOUT')
	},
	//用户刷新
	refresh_user({
		commit
	}) {
		commit('REFRESH_USER')
		var user = JSON.parse(window.localStorage.getItem("user"))
		commit('SET_USER_INFO',user)
	}
}

export default {
	namespaced: true,
	state,
	mutations,
	actions
}
