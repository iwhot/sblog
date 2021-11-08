const state = {
	visitedViews: []
}

const mutations = {
	//增加
	ADD_VISITED_VIEW(state, view) {
		//如果里面有了则不加入
		if (state.visitedViews.some(v => v.path === view.path)) return
		state.visitedViews.push(view)
	},
	//删除一个
	DEL_ONE_VISITED_VIEW(state, view) {
		state.visitedViews = state.visitedViews.filter( v => {
			return v.path !== view.path
		})
	},
	//删除其他
	DEL_OTHER_VISITED_VIEW(state, view) {
		//删除其他 == 只找到当前
		state.visitedViews = state.visitedViews.filter(v => {
			return v.meta.affix || v.path === view.path
		})
	},
	//删除所有
	DEL_ALL_VISITED_VIEW(state) {
		//只保留一个默认的
		const affixTags = state.visitedViews.filter(tag => tag.meta.affix)
		state.visitedViews = affixTags
	}
}

const actions = {
	add_visited_view({
		commit
	}, view) {
		commit('ADD_VISITED_VIEW', view)
	},
	del_one_visited_view({
		commit
	}, view) {
		commit('DEL_ONE_VISITED_VIEW', view)
	},
	del_other_visited_view({
		commit
	}, view) {
		commit('DEL_OTHER_VISITED_VIEW', view)
	},
	del_all_visited_view({
		commit
	}, view) {
		commit('DEL_ALL_VISITED_VIEW', view)
	}
}

export default {
	namespaced: true,
	state,
	mutations,
	actions
}
