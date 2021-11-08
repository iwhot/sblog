<template>
	<div class="control-content">
		<div class="panel-group">
			<el-row :gutter="20">
				<el-col :span="6">
					<div class="card-panel">
						<div class="card-panel-icon green">
							<svg-icon icon-class="peoples" class-name="svg-font"></svg-icon>
						</div>
						<div class="card-panel-cont">
							<div class="num">666</div>
							<div class="desc">会员总数</div>
						</div>
					</div>
				</el-col>

				<el-col :span="6">
					<div class="card-panel">
						<div class="card-panel-icon purple">
							<svg-icon icon-class="chart" class-name="svg-font"></svg-icon>
						</div>
						<div class="card-panel-cont">
							<div class="num">666</div>
							<div class="desc">今日新增会员</div>
						</div>
					</div>
				</el-col>

				<el-col :span="6">
					<div class="card-panel">
						<div class="card-panel-icon red">
							<svg-icon icon-class="money" class-name="svg-font"></svg-icon>
						</div>
						<div class="card-panel-cont">
							<div class="num">666</div>
							<div class="desc">销售总额</div>
						</div>
					</div>
				</el-col>

				<el-col :span="6">
					<div class="card-panel">
						<div class="card-panel-icon blue">
							<svg-icon icon-class="message" class-name="svg-font"></svg-icon>
						</div>
						<div class="card-panel-cont">
							<div class="num">666</div>
							<div class="desc">评论总数</div>
						</div>
					</div>
				</el-col>
			</el-row>

			<el-row :gutter="20">
				<!--测试富文本-->
				<!-- <quill-editor ref="myTextEditor" v-model="content" :options="quillOption" style="height: 500px;"
					@ready="ready($event)">
				</quill-editor> -->
				<editor style="height: 500px;"></editor>
			</el-row>

		</div>
	</div>
</template>

<script>
	import editor from '@/components/editor/index.vue'
	export default {
		name: "Dashboard",
		components:{
			editor
		},
		data() {
			return {
				content: "",
			}
		},
		methods: {
			
		}
	}
</script>

<style lang="scss" scoped="scoped">
	.control-content {
		width: 100%;
		height: 100%;
		background-color: #f5f7f9;
		display: flex;
		align-content: stretch;
		min-width: 1300px;

		.panel-group {
			width: 100%;

			.card-panel {
				padding: 20px;
				background-color: #FFFFFF;
				cursor: pointer;

				display: flex;
				justify-content: flex-start;

				.card-panel-icon {
					width: 80px;
					height: 80px;
					display: flex;
					justify-content: center;
					align-items: center;
					border-radius: 10px;
				}

				//绿色
				.green {
					.svg-font {
						font-size: 48px;
						color: #27C24C;
					}
				}

				.green:hover {
					background-color: #27C24C;

					.svg-font {
						color: #FFFFFF;
					}
				}

				//紫色
				.purple {
					.svg-font {
						font-size: 48px;
						color: #7266ba;
					}
				}

				.purple:hover {
					background-color: #7266ba;

					.svg-font {
						color: #FFFFFF;
					}
				}

				//蓝色
				.blue {
					.svg-font {
						font-size: 48px;
						color: #23b7e5;
					}
				}

				.blue:hover {
					background-color: #23b7e5;

					.svg-font {
						color: #FFFFFF;
					}
				}

				//红色
				.red {
					.svg-font {
						font-size: 48px;
						color: #F05050;
					}
				}

				.red:hover {
					background-color: #F05050;

					.svg-font {
						color: #FFFFFF;
					}
				}


				.card-panel-cont {
					margin-left: 10px;

					.num {
						line-height: 50px;
						font-size: 24px;
						font-weight: bold;
					}

					.desc {
						font-size: 12px;
					}
				}
			}

		}


	}
</style>
