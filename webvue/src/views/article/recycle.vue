<template>
	<div class="common-container">
		<div class="container">
			<div class="search-form">
				<el-form :inline="true" class="demo-form-inline" size="mini">
					<el-form-item label="标题">
						<el-input v-model="form.title" placeholder="请输入标题"></el-input>
					</el-form-item>

					<el-form-item label="作者">
						<el-input v-model="form.username" placeholder="请输入作者"></el-input>
					</el-form-item>

					<div class="show-all" v-show="showAll">

					</div>
					<el-form-item>
						<el-button type="primary">查询</el-button>
					</el-form-item>
					<el-form-item>
						<el-button type="info" @click="resetForm()">重置</el-button>
					</el-form-item>
					<el-form-item>
						<el-button @click="shrink" type="text">
							{{word}}
							<i :class="showAll ? 'el-icon-arrow-up ': 'el-icon-arrow-down'"></i>
						</el-button>
					</el-form-item>
				</el-form>
			</div>

			<div class="table">
				<div class="toolbar">
					<el-button type="primary" size="mini" @click="add">添加</el-button>
				</div>
				<div class="table-container">
					<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%"
						:header-cell-style="{background:'#f6f6f6',color:'#606266'}">

						<el-table-column fixed="left" label="ID" prop="id" width="80"></el-table-column>
						<el-table-column prop="title" label="标题"></el-table-column>
						<el-table-column prop="author" label="作者">
							<template slot-scope="scope">
								{{scope.row.admin.username}}
							</template>
						</el-table-column>
						<!-- <el-table-column prop="attr" label="类型">
							<template slot-scope="scope">
								<span v-if="scope.row.attr == 0" style="color: green;">文章</span>
								<span v-else style="color: red;">单页</span>
							</template>
						</el-table-column> -->
						<el-table-column fixed="right" label="操作" width="100">
							<template #default="scope">
								<el-button @click="update({id:scope.row.id,is_del:0})" type="text" size="small">恢复</el-button>

								<el-popconfirm title="确认要删除？删除后不可恢复" confirmButtonText='确定' cancelButtonText='取消'
									@confirm="del(scope.row.id)">
									<template slot="reference">
										<el-button type="text" size="small">
											删除
										</el-button>
									</template>
								</el-popconfirm>

							</template>
						</el-table-column>
					</el-table>

					<el-pagination background layout="prev, pager, next" :total="form.count" :page-size="form.page_size"
						:current-page="form.page" @current-change="changePage">
					</el-pagination>
				</div>
			</div>
		</div>

	</div>
</template>

<script>
	import {
		ArticleList,
		ArticleUpdateInfo,
		ArticleRealDelete
	} from "@/api/article/index.js"
	export default {
		name: "ArticleRecycle",
		data() {
			return {
				title: "首页",
				showAll: false,
				showAll2: false,
				word: "展开搜索",
				tableData: [],
				multipleSelection: [],
				form: {
					count: 0,
					page: 1,
					page_size: 6,
					start: "",
					end: "",
					state: 0,
					title: "",
					cate_name: "",
					username: "",
					is_del:1,
				}
			}
		},
		methods: {
			shrink() {
				this.showAll = !this.showAll
				if (this.showAll) {
					this.word = "收起搜索"
				} else {
					this.word = "展开搜索"
				}
			},
			shrink2() {
				this.showAll2 = !this.showAll2
			},
			handleSelectionChange(val) {
				this.multipleSelection = val;
			},
			handleClick(row) {
				console.log(row);
			},
			//分页数据
			changePage(currentPage) {
				this.form.page = currentPage
				this.getData()
			},
			add() {
				let routeDate = this.$router.push({
					path: '/article/add',
					// query: {
					// 	id: id
					// }
				})
				window.open(routeDate.href, '_top');
			},
			getData() {
				//获取列表数据
				ArticleList(this.form).then(res => {
					this.tableData = res.data.data.list
					this.form.count = res.data.data.page.count
					this.form.page = res.data.data.page.page
					this.form.pageSize = res.data.data.page.page_size
				})
			},
			del(id) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});
				ArticleRealDelete(id, {}).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						setTimeout(() => {
							loading.close();
						}, 1000);
						return false;
					}

					this.getData()
					loading.close();
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
					return true;
				})
			},
			resetForm() {
				this.form.start = ""
				this.form.end = ""
				this.form.state = 10
				this.form.title = ""
				this.form.cate_name = ""
				this.form.username = ""
			},
			update(data) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});
			
				ArticleUpdateInfo(data).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						setTimeout(() => {
							loading.close();
						}, 1000);
						return false;
					}
			
					this.getData()
					loading.close();
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
					return true;
				})
			},
		},
		mounted() {
			this.getData()
		},
	}
</script>

<style lang="scss" scoped="scoped">
	.container {
		.search-form {
			border: 1px solid #e4e4e4;
			background-color: #f9f9f9;
			padding: 20px 20px 10px 20px;

			.show-all {
				display: inline-block;
			}
		}

		.tips {
			border: 1px solid #c9f2ff;
			background-color: #ecf8fc;
			padding: 10px 10px 10px 10px;
			border-radius: 5px;
			margin-top: 20px;

			ul {
				color: #748A8F;
				font-size: 12px;
				box-sizing: border-box;
			}
		}

		.table {
			padding: 20px 0 0 0;

			.table-container {
				padding-top: 20px;

				.el-pagination {
					padding-top: 20px;
					text-align: right;
				}
			}
		}
	}
</style>
