import request from '@/utils/request.js'

//分类======================================
export async function categoryIndex(data) {
	return await request({
		url: "/category/index",
		method: 'post',
		data
	})
}

export async function categorySave(data) {
	return await request({
		url: "/category/save",
		method: 'post',
		data
	})
}

export async function categoryDelete(mold, data) {
	return await request({
		url: "/category/delete/" + mold,
		method: 'post',
		data
	})
}

export async function GetParentList(data) {
	return await request({
		url: "/category/get-parent-list",
		method: 'post',
		data
	})
}

export async function CategoryUpdateInfo(data) {
	return await request({
		url: "/category/update-info",
		method: 'post',
		data
	})
}

export async function CategoryCascader(data) {
	return await request({
		url: "/category/cascade",
		method: 'post',
		data
	})
}

export async function GetAllCategoryId(id,data) {
	return await request({
		url: "/category/pid/" + id,
		method: 'post',
		data
	})
}


//文章===============================================
//添加
export async function ArticleSave(data) {
	return await request({
		url: "/article/save",
		method: 'post',
		data
	})
}

//获取信息
export async function ArticleInfo(id,data) {
	return await request({
		url: "/article/info/" + id,
		method: 'post',
		data
	})
}

//列表
export async function ArticleList(data) {
	return await request({
		url: "/article/index",
		method: 'post',
		data
	})
}

//删除
export async function ArticleDelete(id,data) {
	return await request({
		url: "/article/delete/" + id,
		method: 'post',
		data
	})
}

//真实删除
export async function ArticleRealDelete(id,data) {
	return await request({
		url: "/article/real-delete/" + id,
		method: 'post',
		data
	})
}

export async function ArticleUpdateInfo(data) {
	return await request({
		url: "/article/update-info",
		method: 'post',
		data
	})
}

//标签================================================
export async function indexTags(data) {
    return await request({
        url: "/tags/index",
        method: 'post',
        data
    })
}

export async function saveTags(data) {
    return await request({
        url: "/tags/save",
        method: 'post',
        data
    })
}

export async function deleteTags(data) {
    return await request({
        url: "/tags/delete",
        method: 'post',
        data
    })
}
