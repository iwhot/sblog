<template>
	<div class="page-content">
		<svg-icon icon-class="404" class-name="svg-font" />
		<br>
		<el-button @click="index()">返回首页</el-button>
	</div>
</template>

<script>
	export default {
		name: "Page404",
		methods: {
			index() {
				this.$router.push({
					path: '/',
					replace: true
				});
			}
		},
	}
</script>

<style scoped="scoped" lang="scss">
	.page-content {
		width: 100%;
		text-align: center;
		padding: 50px 0 0 0;
	}

	.svg-font {
		font-size: 300px;
	}
</style>
