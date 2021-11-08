<template>
	<fragment v-if="!item.hidden">

		<template v-if="hasOneShowingChild(item)">
			
			<el-menu-item :index="item.path">
				<svg-icon v-if="item.meta && item.meta.icon" :icon-class="item.meta.icon" class-name="svg-font"></svg-icon>
				<template slot="title"> <span class="title-font">{{item.title}}</span> </template>
			</el-menu-item>
		</template>

		<template v-else>
			<el-submenu :index="item.path">
				<template slot="title">
					<svg-icon v-if="item.meta && item.meta.icon" :icon-class="item.meta.icon" class-name="svg-font"></svg-icon>
					<span class="title-font">{{item.title}}</span>
				</template>

				<sidebar v-for="child in item.children" :key="child.path" :item="child" :base-path="child.path"></sidebar>
			</el-submenu>
		</template>

	</fragment>
</template>

<script>
	export default {
		name: "Sidebar",
		props: {
			item: {
				type: Object,
				required: true
			},
			basePath: {
				type: String,
				default: ''
			}
		},
		data() {
			return {}
		},
		methods: {
			hasOneShowingChild(item) {
				var child = item.children
				if (!child || child.length === 0) {
					return true
				}


				const showingChildren = child.filter(item => {
					if (item.hidden) {
						return false
					} else {
						return true
					}
				})

				if (showingChildren.length >= 1) {
					return false
				}

				return true

			}
		}
	}
</script>

<style lang="scss" scoped="scoped">
	.svg-font{
		font-size: 16px;
	}
	.title-font{
		padding-left: 10px;
	}
</style>
