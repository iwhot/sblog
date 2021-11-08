<template>

		<el-aside :width="isCollapse === false ? '200px' : '64px'" class="aside-right">
			<div :class="isCollapse === false ? 'left-menu' : 'left-menu-mini'">
				<!--菜单开始-->
				<vuescroll :ops="ops">
					<el-menu :default-active="activeMenu" class="el-menu-vertical" background-color="#415b94"
						text-color="#bfcbd9" active-text-color="#ffffff" :collapse="isCollapse"
						:collapse-transition="collapseTransition" :unique-opened="true" mode="vertical" :router="true">

						<sidebar v-for="route in permission_routes" :key="route.path" :item="route"
							:base-path="route.path">
						</sidebar>

					</el-menu>
				</vuescroll>
				<!--菜单结束-->
			</div>
		</el-aside>

</template>

<script>
	import {
		mapGetters
	} from 'vuex'
	import Sidebar from './sidebar.vue'
	import vuescroll from 'vuescroll';

	export default {
		name: "LeftMenu",
		components: {
			Sidebar,
			vuescroll
		},
		computed: {
			...mapGetters([
				'permission_routes'
			]),
			activeMenu() {
				const route = this.$route
				// console.log(route)
				const {
					meta,
					path
				} = route

				if (meta.activeMenu) {
					return meta.activeMenu
				}
				return path
			}
		},
		data() {
			return {
				collapseTransition: false,
				ops: {
					bar: {
						hoverStyle: true,
						onlyShowBarOnScroll: false, //是否只有滚动的时候才显示滚动条
						background: "#F5F5F5", //滚动条颜色
						opacity: 0.5, //滚动条透明度
						"overflow-x": "hidden"
					}
				}
			}
		},
		props: {
			isCollapse: {
				type: Boolean,
				default: false
			}
		}
	}
</script>

<style lang="scss" scoped="scoped">
	.aside-left {
		background-color: #415b94;
		height: 100%;
		background-color: #415b94;
	}

	.aside-right {
		/* background-color: #fafafa; */
		background-color: #415b94;
		box-shadow: 10px 0 10px -10px #c7c7c7;
		z-index: 9999;
		// 侧边栏折叠动画速度
		transition: width 0.3s;
		-webkit-transition: width 0.3s;
		-moz-transition: width 0.3s;
		-webkit-transition: width 0.3s;
		-o-transition: width 0.3s;
		height: 100%;

		.left-menu {
			background-color: #415b94;
			-webkit-transition: width .28s;
			bottom: 0;
			font-size: 0;
			height: 100%;
			left: 0;
			overflow: hidden;
			position: fixed;
			top: 50px;
			transition: width .28s;
			width: 200px !important;
			z-index: 1001;

			.scrollbar-wrapper {
				overflow-x: hidden;
			}

			.el-menu-vertical {
				border: none;
			}
		}

		.left-menu-mini {
			background-color: #415b94;
			-webkit-transition: width .28s;
			bottom: 0;
			font-size: 0;
			height: 100%;
			left: 0;
			overflow: hidden;
			position: fixed;
			top: 50px;
			transition: width .28s;
			width: 64px !important;
			z-index: 1001;

			.scrollbar-wrapper {
				overflow-x: hidden;
			}

			.el-menu-vertical {
				border: none;
			}
		}
	}
</style>
