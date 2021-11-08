<template>
	<el-breadcrumb separator-class="el-icon-arrow-right">
		<el-breadcrumb-item v-for="(item,index) in levelList" :key="index">{{item.meta.title}}
		</el-breadcrumb-item>
	</el-breadcrumb>
</template>

<script>
	export default {
		name: "Breadcrumb",
		created() {
			this.getBreadcrumb()
		},
		watch: {
			$route() {
				this.getBreadcrumb();
			}
		},
		data() {
			return {
				levelList: []
			}
		},
		methods: {
			isDashboard(route) {
				const name = route && route.name
				if (!name) {
					return false
				}
				return name.trim().toLocaleLowerCase() === 'dashboard'.toLocaleLowerCase()
			},
			getBreadcrumb() {
				let matched = this.$route.matched.filter(item => item.meta && item.meta.title)
				if (!this.isDashboard(matched[0])) {
					matched = [{
						path: "/dashboard",
						meta: {
							title: "控制台"
						}
					}].concat(matched)
				}
				this.levelList = matched
			}
		}
	}
</script>

<style>
</style>
