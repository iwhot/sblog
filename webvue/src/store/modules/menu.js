import {
	routes,
	asyncRoutes
} from '@/router'
import Layout from '@/layout'

import {
	getLeftMenu
} from '@/api/common/index.js'
import EmptyLayout from '@/layout/empty.vue'


export function filterAsyncRoutes(routes) {
	const res = []

	// console.log(routes)

	routes.forEach(route => {
		const tmp = {
			...route
		}

		//console.log(tmp)

		var cmt = tmp.component
		if (cmt) {
			// console.log(tmp.component)
			// tmp.component = require(`../../views${tmp.component}`).default

			if (cmt === 'Layout') {
				tmp.component = Layout
			} else {
				tmp.component = resolve => require([`../../views${cmt}`], resolve)
				// tmp.component = () => import(`@/views${tmp.component}`)
			}

		} else {
			tmp.component = EmptyLayout
		}

		if (tmp.isShow == 0) {
			tmp.hidden = true
		} else {
			tmp.hidden = false
		}

		//重置meta
		tmp.meta = {}
		if (tmp.title) {
			tmp.meta.title = tmp.title
		}

		if (tmp.icon) {
			tmp.meta.icon = tmp.icon
		}

		if (tmp.affix == true) {
			tmp.meta.affix = tmp.affix
		} else {
			tmp.meta.affix = false
		}

		if (tmp.children && tmp.children.length) {
			tmp.children = filterAsyncRoutes(tmp.children)
		}
		// console.log(tmp)
		res.push(tmp)

	})

	return res
}

const state = {
	routes: [],
	allRoutes: [],
	addRoutes: []
}

const mutations = {
	SET_ROUTES: (state, route2) => {
		state.addRoutes = route2
		state.routes = route2
		state.allRoutes = routes.concat(route2)

		// console.log(state.routes)
	},
	SET_MENU: (state) => {
		var menu = window.localStorage.getItem("menu")
		if (menu) {
			return true
		}
		getLeftMenu({}).then(res => {
			if (res.data.code == 200) {
				newMenu = newMenu.concat(res.data.data)
			}
			window.localStorage.setItem("menu", JSON.stringify(newMenu))
		})
	}
}

const actions = {
	generateRoutes({
		commit
	}) {
		//获取菜单
		return new Promise((resolve, reject) => {
			getLeftMenu({}).then(res => {
				window.localStorage.removeItem('menu')
				var newMenu = asyncRoutes[0].children
				if (res.data.code == 200) {
					newMenu = newMenu.concat(res.data.data)
				}
				asyncRoutes[0].children = newMenu
				window.localStorage.setItem("menu", JSON.stringify(newMenu))
				//加入路由
				var accessedRoutes = filterAsyncRoutes(asyncRoutes)
				commit('SET_ROUTES', accessedRoutes[0].children)
				resolve(accessedRoutes)
			}).catch(err => {
				reject(err)
			})
		})
	}
}


export default {
	namespaced: true,
	state,
	mutations,
	actions
}
