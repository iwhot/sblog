<template>
	<div class="tabs">
		<el-scrollbar class="scroll-container">
			<div class="flex-content">

				<router-link ref="tag" v-for="(item,index) in visitedViews" :key="index" :to="item.path"
					:class="isActive(item) ? 'tags-view-item active' : 'tags-view-item'"
					@contextmenu.prevent.native="openMenu(item,$event)">
					&nbsp;{{item.meta.title}}&nbsp; <span class="el-icon-close"
						@click.prevent.stop="closeOneTags(item)"></span>
				</router-link>

			</div>
		</el-scrollbar>
		<ul class="contextmenu" v-show="visible" :style="{left:left+'px',top:top+'px'}">
			<li @click="closeOneTags(selectedTag)">关闭</li>
			<li @click="closeOtherTags(selectedTag)">关闭其他</li>
			<li @click="closeAllTags(selectedTag)">关闭所有</li>
		</ul>
	</div>
</template>

<script>
	export default {
		name: "ScrollPane",
		data() {
			return {
				visible: false,
				top: 0,
				left: 0,
				selectedTag: {},
				affixTags: []
			}
		},
		methods: {
			//打开菜单
			openMenu(item, e) {
				const menuMinWidth = 105
				const offsetLeft = this.$el.getBoundingClientRect().left // container margin left
				const offsetWidth = this.$el.offsetWidth // container width
				const maxLeft = offsetWidth - menuMinWidth // left boundary
				const left = e.clientX - offsetLeft + 15 // 15: margin right

				if (left > maxLeft) {
					this.left = maxLeft
				} else {
					this.left = left
				}

				this.top = e.clientY - 34
				this.visible = true
				this.selectedTag = item
			},
			//关闭菜单
			closeMenu() {
				this.visible = false
			},
			//初始化标签
			initTags() {
				//默认加入控制台进去
				var tag = {
					path: '/dashboard',
					fullPath: "/dashboard",
					meta: {
						title: '控制台',
						affix: true
					}
				}
				this.$store.dispatch('tagsView/add_visited_view', tag)
			},
			//添加标签
			addTags() {
				if (this.$route.meta.title) {
					this.$store.dispatch('tagsView/add_visited_view', this.$route)
				}
				return false;
			},
			//是否选中
			isActive(route) {
				return route.path === this.$route.path
			},
			//关闭一个标签
			closeOneTags(item) {
				if (item.path == '/dashboard') return false;
				this.$store.dispatch('tagsView/del_one_visited_view', item).then(() => {
					this.toLastView(item)
				})
			},
			//关闭其他标签
			closeOtherTags(item) {
				this.$store.dispatch('tagsView/del_other_visited_view', item)
				//跳转到当前
				this.$router.push(item)

			},
			//关闭所有标签
			closeAllTags(item) {
				this.$store.dispatch('tagsView/del_all_visited_view', item)
				this.$router.push({
					path: '/dashboard'
				})
			},
			toLastView(item) {
				const latestView = this.visitedViews.slice(-1)[0]
				if (latestView) {
					this.$router.push(latestView.fullPath)
				}
			}
		},
		watch: {
			//监控路由
			$route() {
				this.addTags()
			},
			visible(value) {
				if (value) {
					document.body.addEventListener('click', this.closeMenu)
				} else {
					document.body.removeEventListener('click', this.closeMenu)
				}
			}
		},
		mounted() {
			this.initTags()
			this.addTags()
		},
		computed: {
			visitedViews() {
				return this.$store.state.tagsView.visitedViews
			}
		}
	}
</script>

<style scoped="scoped" lang="scss">
	.scroll-container {
		overflow: hidden;
		position: relative;
		white-space: nowrap;
		width: 100%;
	}

	.tabs {
		height: 34px;
		width: 100%;
		background-color: #fff;
		border-bottom: 1px solid #d8dce5;
		border-top: 1px solid #eef0f5;
		box-shadow: 0 1px 4px 0 #ccc;
		position: relative;

		.flex-content {
			padding-left: 10px;
		}

		.contextmenu {
			-webkit-box-shadow: 2px 2px 3px 0 rgba($color: #000000, $alpha: 0.3);
			background: #fff;
			border-radius: 4px;
			box-shadow: 2px 2px 3px 0 rgba($color: #000000, $alpha: 0.3);
			color: #333;
			font-size: 12px;
			font-weight: 400;
			list-style-type: none;
			margin: 0;
			padding: 5px 0;
			position: absolute;
			z-index: 100;

			li {
				cursor: pointer;
				margin: 0;
				padding: 7px 16px;
			}

			li:hover {
				background-color: #EEEEEE;
			}
		}

		.tags-view-item {
			background: #fff;
			border: 1px solid #d8dce5;
			color: #495060;
			display: inline-block;
			font-size: 12px;
			height: 26px;
			line-height: 26px;
			margin-left: 5px;
			margin-top: 4px;
			padding: 0 8px;
			position: relative;
			text-decoration: none;

			.el-icon-close {
				-webkit-transform-origin: 100% 50%;
				-webkit-transition: all .3s cubic-bezier(.645, .045, .355, 1);
				border-radius: 50%;
				height: 16px;
				text-align: center;
				transform-origin: 100% 50%;
				transition: all .3s cubic-bezier(.645, .045, .355, 1);
				vertical-align: 2px;
				width: 16px;
			}

			.el-icon-close:before {
				-webkit-transform: scale(.6);
				display: inline-block;
				transform: scale(.6);
				vertical-align: -3px;
			}

			.el-icon-close:hover {
				background-color: #b4bccc;
				color: #fff;
			}
		}

		.tags-view-item.active {
			background-color: #8a93a9;
			border-color: #8a93a9;
			color: #fff;
			border-radius: 2px;
		}

		.tags-view-item.active:before {
			background: #fff;
			border-radius: 50%;
			content: "";
			display: inline-block;
			height: 8px;
			margin-right: 2px;
			position: relative;
			width: 8px;
		}
	}
</style>
