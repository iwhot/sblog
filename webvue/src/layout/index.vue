<template>
	<div class="layout-container">
		<header class="header">
			<!--左边logo-->
			<div class="left">

				<div :class="!isCollapse ? 'logo' : 'logo2' ">
					<svg-icon :icon-class="!isCollapse ? 'logo1' : 'logo2'" class-name="svg-logo" />
				</div>

				<div class="icon-bar" @click="shrink">
					<i :class="isCollapse === false ? 'el-icon-s-fold' : 'el-icon-s-unfold'"></i>
				</div>

				<!--面包屑-->
				<div class="breadcrumb">
					<breadcrumb></breadcrumb>
				</div>

			</div>

			<!--顶部右边-->
			<right></right>
		</header>

		<div class="bottom">
			<AssideMenu :isCollapse="isCollapse"></AssideMenu>
			<div class="main">
				<ScrollPane></ScrollPane>
				<AppMain></AppMain>
			</div>
		</div>
	</div>
</template>

<script>
	import right from './components/top/right.vue'
	import AssideMenu from './components/left/menu.vue'
	import AppMain from './components/main/main.vue'
	import ScrollPane from './components/main/scrollPanel.vue'
	import Breadcrumb from './components/top/breadcrumb.vue'
	export default {
		name: "Layout",
		components: {
			right,
			AssideMenu,
			AppMain,
			ScrollPane,
			Breadcrumb
		},
		data() {
			return {
				isCollapse: false,
			}
		},
		methods: {
			shrink() {
				this.isCollapse = !this.isCollapse
			}
		},
	}
</script>

<style lang="scss" scoped="scoped">
	.layout-container {
		width: 100%;
		height: 100%;

		//头部
		.header {
			width: 100%;
			height: 50px;
			background-color: #FFFFFF;
			display: flex;
			justify-content: space-between; //两端对齐
			align-items: center;

			.left {
				display: flex;
				align-items: center;

				.icon-bar {
					width: 64px;
					height: 50px;
					cursor: pointer;
					display: flex;
					align-items: center;
					justify-content: center;
					color: #333333;
					font-size: 30px;
				}

				.icon-bar:hover {
					background-color: #f9f9f9;
				}

				.logo {
					// width: 180px;
					transition: all 0.3s;
					background-color: #415b94;
					padding: 8px 58px;

					.svg-logo {
						width: 84px;
						height: 35px;
					}
				}

				.logo2 {
					padding: 7px 0px;
					background-color: #415b94;
					// width: 0px;
					transition: all 0.3s;
					.svg-logo {
						width: 64px;
						height: 35px;
					}
				}

				.breadcrumb {
					// padding-left: 20px;
				}

			}

		}

		//下面部分
		.bottom {
			width: 100%;
			height: calc(100% - 50px);
			display: flex;

			.main {
				width: 100%;
				height: 100%;
				overflow: hidden;
			}
		}
	}
</style>
