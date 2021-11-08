import request from '@/utils/request.js'
//保存
export async function saveConfig(data) {
    return await request({
        url: "/config/save",
        method: 'post',
        data
    })
}


//列表
export async function indexConfig(data) {
    return await request({
        url: "/config/index",
        method: 'post',
        data
    })
}

//链接---------------------------------------------
export async function indexLink(data) {
    return await request({
        url: "/link/index",
        method: 'post',
        data
    })
}

export async function saveLink(data) {
    return await request({
        url: "/link/save",
        method: 'post',
        data
    })
}

export async function deleteLink(data) {
    return await request({
        url: "/link/delete",
        method: 'post',
        data
    })
}