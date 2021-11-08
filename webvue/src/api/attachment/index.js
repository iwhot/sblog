import request from '@/utils/request.js'

export async function index(mold, data) {
	return await request({
		url: "/attachment/index/" + mold,
		method: 'post',
		data
	})
}

export async function upload(data) {
	return await request({
		url: "/attachment/upload",
		method: 'post',
		data
	})
}

export async function del(data) {
	return await request({
		url: "/attachment/delete",
		method: 'post',
		data
	})
}

export async function download(data) {
	return await request({
		url: "/attachment/download",
		method: 'post',
		data
	})
}
