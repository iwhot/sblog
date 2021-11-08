import router,{resetRouter} from './router/index.js'
import store from './store/index.js'
import NProgress from 'nprogress' // 进度条
import 'nprogress/nprogress.css' // 进度条样式
import {
	getToken
} from '@/utils/auth.js'

NProgress.configure({
	showSpinner: false
})

var token = getToken()

const whiteList = ['/login']

router.beforeEach(async (to, from, next) => {
	// console.log(to)
	NProgress.start()
	if (token) {
		if (store.state.menu.routes.length == 0) {
			await store.dispatch('menu/generateRoutes').then(res => {
				// console.log(res)
				router.addRoutes(res)
				next({
					...to,
					replace: true
				})
				
			})
		}
		
		if (to.path === '/login') {
			next({
				path: '/'
			})
			
			NProgress.done()
		} else {
			next()
		}
		
	} else {
		if (whiteList.indexOf(to.path) !== -1) {
			next()
		} else {
			next('/login')
			NProgress.done()
		}
	}
	
	// console.log(router.matcher.getRoutes())

	NProgress.done()
})

router.afterEach((to, from, failure) => {
	if (to.meta.title) {
		document.title = to.meta.title;
	}
	NProgress.done()
})
