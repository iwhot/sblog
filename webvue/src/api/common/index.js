import request from '@/utils/request.js'

//一些公共请求
export async function getLeftMenu(data) {
	return await request({
		url: "/auth-rule/get-menu",
		method: 'post',
		data
	})
	
}

//获取用户信息
export async function userInfo(data) {
	return await request({
		url: "/user-info",
		method: 'post',
		data
	})
}