<template>
	<el-main>
		<transition name="fade-transform" mode="out-in">
			<div class="main-container" :key="key">
				<router-view :key="key"></router-view>
			</div>
		</transition>
	</el-main>
</template>

<script>
	export default {
		name: "AppMain",
		computed: {
			key() {
				// console.log(this.$route.path)
				return this.$route.path
			}
		}
	}
</script>

<style lang="scss" scoped="scoped">
	/* fade-transform */
	.fade-transform-leave-active,
	.fade-transform-enter-active {
		transition: all .5s;
	}

	.fade-transform-enter-from {
		opacity: 0;
		transform: translateX(-30px);
	}

	.fade-transform-leave-to {
		opacity: 0;
		transform: translateX(30px);
	}

	.el-main {
		background-color: #f5f7f9;
		--el-main-padding: 20px;
		display: block;
		flex: 1;
		flex-basis: auto;
		overflow: auto;
		padding: var(--el-main-padding);
		width: 100%;
		height: 100%;

		.main-container {
			// background-color: #FFFFFF;
			// border-radius: 8px;
			// padding: 20px;
			display: block;
			// height: 100%;
			min-height: 800px;
		}
	}
</style>
