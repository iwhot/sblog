import request from '@/utils/request.js'

//管理员管理=================================
export async function indexAdmin(data) {
	return await request({
		url: "/admin/index",
		method: 'post',
		data
	})
}

export async function saveAdmin(data) {
	return await request({
		url: "/admin/save",
		method: 'post',
		data
	})
}

export async function delAdmin(data) {
	return await request({
		url: "/admin/delete",
		method: 'post',
		data
	})
}

export async function delAllAdmin(data) {
	return await request({
		url: "/admin/delete-all",
		method: 'post',
		data
	})
}

export async function saveAdmin2(data) {
	return await request({
		url: "/admin/save2",
		method: 'post',
		data
	})
}

export async function GetMyGroupID(data) {
	return await request({
		url: "/admin/get-group-id",
		method: 'post',
		data
	})
}

//用户组相关=================================
export async function getAllGroup(data) {
	return await request({
		url: "/auth-group/get-all",
		method: 'post',
		data
	})
}

export async function getGroupIndex(data) {
	return await request({
		url: "/auth-group/index",
		method: 'post',
		data
	})
}

export async function saveGroup(data) {
	return await request({
		url: "/auth-group/save",
		method: 'post',
		data
	})
}

export async function delGroup(data) {
	return await request({
		url: "/auth-group/delete",
		method: 'post',
		data
	})
}

export async function delAllGroup(data) {
	return await request({
		url: "/auth-group/delete-all",
		method: 'post',
		data
	})
}

//菜单相关==================================
export async function cascadeMenu(data) {
	return await request({
		url: "/auth-rule/cascade-menu",
		method: 'post',
		data
	})
}

export async function saveMenu(data) {
	return await request({
		url: "/auth-rule/save",
		method: 'post',
		data
	})
}

export async function getMenuIndex(data) {
	return await request({
		url: "/auth-rule/index",
		method: 'post',
		data
	})
}

export async function delMenu(data) {
	return await request({
		url: "/auth-rule/delete",
		method: 'post',
		data
	})
}

export async function getMenuPid(data) {
	return await request({
		url: "/auth-rule/get-menu-pid",
		method: 'post',
		data
	})
}

export async function updateMenuStatus(data) {
	return await request({
		url: "/auth-rule/update-menu-status",
		method: 'post',
		data
	})
}

export async function updateMenuSort(data) {
	return await request({
		url: "/auth-rule/update-menu-sort",
		method: 'post',
		data
	})
}
