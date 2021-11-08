const getters = {
	permission_routes: state => state.menu.routes,
	oss_config: state => state.oss.OSSConfig,
	visitedViews: state => state.tagsView.visitedViews,
	message: state => state.websocket.message,
	user_info: state=> state.user.userInfo
}

export default getters
