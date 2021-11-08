import request from '@/utils/request.js'

export async function MenuIndex(data) {
	return await request({
		url: "/menu/index",
		method: 'post',
		data
	})
}

export async function MenuSave(data) {
	return await request({
		url: "/menu/save",
		method: 'post',
		data
	})
}

export async function MenuDelete(mold, data) {
	return await request({
		url: "/menu/delete/" + mold,
		method: 'post',
		data
	})
}

export async function GetParentList(data) {
	return await request({
		url: "/Menu/get-parent-list",
		method: 'post',
		data
	})
}

export async function MenuUpdateInfo(data) {
	return await request({
		url: "/menu/update-info",
		method: 'post',
		data
	})
}

export async function MenuCascader(data) {
	return await request({
		url: "/menu/cascade",
		method: 'post',
		data
	})
}

export async function GetAllMenuId(id,data) {
	return await request({
		url: "/menu/pid/" + id,
		method: 'post',
		data
	})
}