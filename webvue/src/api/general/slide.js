import request from '@/utils/request.js'

//分类===================================================
export async function indexSlideCategory(data) {
    return await request({
        url: "/slide-category/index",
        method: 'post',
        data
    })
}

export async function saveSlideCategory(data) {
    return await request({
        url: "/slide-category/save",
        method: 'post',
        data
    })
}

export async function deleteSlideCategory(data) {
    return await request({
        url: "/slide-category/delete",
        method: 'post',
        data
    })
}

export async function findOneSlideCategory(id,data) {
    return await request({
        url: "/slide-category/find-one/" + id,
        method: 'post',
        data
    })
}


//banner============================================================

export async function indexSlideBanner(cid,data) {
    return await request({
        url: "/slide-banner/index/" + cid,
        method: 'post',
        data
    })
}

export async function saveSlideBanner(data) {
    return await request({
        url: "/slide-banner/save",
        method: 'post',
        data
    })
}

export async function deleteSlideBanner(data) {
    return await request({
        url: "/slide-banner/delete",
        method: 'post',
        data
    })
}