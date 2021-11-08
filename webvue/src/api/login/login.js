import request from '@/utils/request.js'

export async function login(data) {
	return await request({
		url: "/login",
		method: 'post',
		data
	})
}

export async function captcha(data) {
	return await request({
		url: "/get-captcha",
		method: 'post',
		data
	})
}
