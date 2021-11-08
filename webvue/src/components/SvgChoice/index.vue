<template>
	<el-popover placement="bottom" trigger="click" :width="564" v-model:visible="visible">
		<template slot="reference">
			<el-input v-model="svgName" :value="value" placeholder="请输入图标" @input="transmission">
			</el-input>
		</template>
		<ul class="icon-content">
			<li v-for="(item,index) in content" :key=index @click="setSvg(item)">
				<svg-icon :icon-class="item" class-name="svg-font"></svg-icon>
			</li>
		</ul>
	</el-popover>
</template>

<script>
	export default {
		name: "SvgChoice",
		props: {
			value: ""
		},
		model: {
			prop: 'value',
			event: 'getName'
		},
		created(){
			this.getSvg();
		},
		methods: {
			//获取svg
			getSvg() {
				//载入svg
				var req = require.context('@/icons/svg', false, /\.svg$/).keys()
				var iconsName = []
				req.forEach(item => {
					item = item.replace(".svg", "")
					item = item.replace("./", "")
					iconsName.push(item)
				})
				this.content = iconsName
			},
			//设置svg
			setSvg(name) {
				this.svgName = name
				this.visible = false
				this.$emit("getName", name)
			},
			//向父组件传值
			transmission() {
				// console.log("子组件",this.svgName)
				this.$emit("getName", this.svgName)
			}
		},
		data() {
			return {
				visible: false,
				content: [],
				svgName: this.value
			}
		}
	}
</script>

<style scoped="scoped" lang="scss">
	.icon-content {
		padding: 0;
		overflow: hidden;

		li {
			list-style: none;
			float: left;
			padding: 10px;
			border: 1px solid #dddddd;
			border-radius: 5px;
			margin: 0 5px 5px 0;
			cursor: pointer;
		}

		li:hover {
			background-color: #EEEEEE;
		}

		.svg-font {
			font-size: 20px;
		}
	}
</style>
