<template>
	<div class="right">
		<el-tooltip effect="light" :content="tip1" placement="bottom">
			<div class="full" @click="full">
				<!-- <i class="el-icon-full-screen"></i> -->
				<svg-icon :icon-class="screen" class-name="svg-font"></svg-icon>
			</div>
		</el-tooltip>

		<el-tooltip effect="light" content="刷新" placement="bottom">
			<div class="refresh" @click="refresh">
				<!-- <i class="el-icon-refresh"></i> -->
				<svg-icon icon-class="refresh" class-name="svg-font"></svg-icon>
			</div>
		</el-tooltip>

		<el-dropdown trigger="click">
			<span class="el-dropdown-link">
				<img v-if="!logo" src="../../../assets/default.gif">
				<img v-else :src="logo">
				<i class="el-icon-caret-bottom el-icon--right"></i>
			</span>
			<template #dropdown>
				<el-dropdown-menu>
					<el-dropdown-item @click.native="info">个人资料</el-dropdown-item>
					<el-dropdown-item @click.native="logout">退出登陆</el-dropdown-item>
				</el-dropdown-menu>
			</template>
		</el-dropdown>
	</div>
</template>

<script>
	import screenfull from 'screenfull'

	export default {
		name: "TopRight",
		data() {
			return {
				logo: this.$store.state.user.userInfo.avatar ? process.env.VUE_APP_BASE_URI + this.$store.state.user
					.userInfo.avatar : "",
				screen: "fullscreen",
				tip1: "全屏",
				dot: false
			}
		},
		methods: {
			full() {
				if (this.screen == "fullscreen") {
					this.screen = "exit-fullscreen";
					this.tip1 = "取消全屏"
				} else {
					this.screen = "fullscreen";
					this.tip1 = "全屏"
				}
				screenfull.toggle()
			},
			refresh() {
				location.reload()
			},
			//清除缓存
			clear() {
				clear({}).then(res => {
					this.$message.success({
						message: res.data.msg,
						type: 'success'
					});
				})
			},
			//登出
			logout() {
				this.$message.success({
					message: "successfully",
					type: 'success'
				});
				this.$store.dispatch("user/logout")
			},
			//打开个人资料
			info() {
				this.$router.push({
					path: '/general/config/info',
					replace: true
				});
			}
		},
		watch: {
			listenUserInfo(info) {
				console.log(info)
			}
		},
		computed: {
			listenUserInfo() {
				return this.$store.state.user.userInfo
			}
		}
	}
</script>

<style scoped="scoped" lang="scss">
	.right {
		display: flex;
		align-items: center;
		padding-right: 15px;

		.el-dropdown {
			.el-dropdown-link {
				cursor: pointer;
				padding-left: 8px;

				img {
					width: 40px;
					height: 40px;
					border-radius: 10px;
				}
			}
		}

		.full,
		.refresh,
		.wechat {
			padding: 0 8px;
			cursor: pointer;
			outline: 0 !important;
		}

		.svg-font {
			width: 18px;
			height: 18px;
			color: #555;
			font-weight: 0;
		}
	}
</style>
