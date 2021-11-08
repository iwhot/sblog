import Vue from 'vue'
import VueRouter from 'vue-router'


const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
	return originalPush.call(this, location).catch(err => err)
}


Vue.use(VueRouter)

export const routes = [{
	path: '/login',
	name: 'Login',
	component: () => import("@/views/login/login.vue"),
	hidden: false,
	meta: {
		title: "登录页"
	}
}]


export const asyncRoutes = [{
		path: "/",
		name: "Index",
		component: 'Layout',
		redirect: "/dashboard",
		hidden: false,
		children: [{
			path: '/dashboard',
			name: "Dashboard",
			component: "/dashboard.vue",
			hidden: false,
			icon: "dashboard",
			title: "控制台",
			affix: true
		}]
	}

]

/*
, {
	path: '*',
	name: "Page404",
	component: "/404.vue",
	hidden: true,
}
*/

const createRouter = () => new VueRouter({
	// mode: 'history', // require service support
	scrollBehavior: () => ({
		y: 0
	}),
	routes
})

const router = createRouter()

export function resetRouter() {
	const newRouter = createRouter()
	router.matcher = newRouter.matcher // reset router
}

// router.addRoutes([{
// 	path: '/aaa',
// 	name: 'AAA',
// 	component: () => import("@/views/dashboard.vue"),
// 	hidden: false,
// 	meta: {
// 		title: "啊啊啊"
// 	}
// }]);

export default router
