<template>
	<div class="common-container">

		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="addtags">添加</el-button>

			</div>
			<div class="table-container">
				<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%"
					:header-cell-style="{background:'#f6f6f6',color:'#606266'}" :border="true" :key="Math.random()">

					<el-table-column label="ID" prop="id" width="80"></el-table-column>
					<el-table-column prop="tag_name" label="名称"></el-table-column>
					<el-table-column prop="createtime" label="创建时间"></el-table-column>
					<el-table-column prop="updatetime" label="更新时间"></el-table-column>
					
					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="edittags(scope.row)">编辑
							</el-button>

							<el-popconfirm title="确认要删除？删除后不可恢复" confirmButtonText='确定'
								cancelButtonText='取消' @confirm="del(scope.row.id)">
								<template slot="reference">
									<el-button type="text" size="small">
										删除
									</el-button>
								</template>
							</el-popconfirm>

						</template>
					</el-table-column>
				</el-table>

				<el-pagination background layout="prev, pager, next" :total="page.count" :page-size="page.pageSize"
					:current-page="page.page" @current-change="changePage">
				</el-pagination>
			</div>
		</div>

		<el-dialog :title="formTitle" :visible.sync="dialogFormVisible" destroy-on-close width="45%"
			custom-class="my-dialog" :close-on-click-modal="false">
			<el-scrollbar wrap-class="form-wrapper">
				<el-form ref="form" :model="form" label-width="100px">
					<el-form-item label="名称" prop="tag_name">
						<el-input autocomplete="off" v-model="form.tag_name" placeholder="请输入标签名称"></el-input>
					</el-form-item>
				</el-form>
			</el-scrollbar>
			<template slot="footer">
				<span class="dialog-footer">
					<el-button @click="dialogFormVisible = false" size="mini">取 消</el-button>
					<el-button type="primary" @click="save('form')" size="mini">确 定</el-button>
				</span>
			</template>
		</el-dialog>

	</div>
</template>

<script>
	import {
		indexTags,
		saveTags,
		deleteTags,
	} from '@/api/article/index.js'

	export default {
		name: "tagsIndex",
		data() {
			return {
				title: "首页",
				showAll2: true,
				dialogFormVisible: false,
				word: "展开搜索",
				tips: "操作提示",
				tableData: [],
				page: {
					count: 0,
					page: 1,
					pageSize: 8
				},
				formTitle: "添加",
				form: {
					id: 0,
					tag_name:""
				},
			}
		},
		methods: {
			//获取数据
			getData() {
				indexTags({
					page: this.page.page,
					page_size: this.page.pageSize
				}).then(res => {
					this.tableData = res.data.data.list
					this.page.count = res.data.data.page.count
					this.page.page = res.data.data.page.page
					this.page.pageSize = res.data.data.page.page_size
				})
			},
			del(id) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				deleteTags({
					id: id
				}).then(res => {

					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						setTimeout(() => {
							loading.close();
						}, 1000);
						return false;
					}

					this.getData()

					this.$message({
						message: res.data.msg,
						type: 'success'
					});

					setTimeout(() => {
						loading.close();
						//跳转页面
						// window.location.reload()
					}, 1000);
					return true;
				})
			},
			//编辑
			edittags(row) {
				// console.log(row)
				this.dialogFormVisible = true
				this.formTitle = "编辑"
				this.form.id = row.id
				this.form.tag_name = row.tag_name
			},
			//更新信息
			update(row) {
				this.form.id = row.id
				this.form.tag_name = row.tag_name
				this.save()
			},
			addtags() {
				this.dialogFormVisible = true
				this.formTitle = "添加"
				this.form.tag_name = ""
				this.form.id = 0

			},
			save() {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				saveTags(this.form).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						setTimeout(() => {
							loading.close();
						}, 1000);
						return false;
					}

					this.getData()

					this.$message({
						message: res.data.msg,
						type: 'success'
					});

					this.dialogFormVisible = false

					setTimeout(() => {
						loading.close();
						//跳转页面
						//window.location.reload()
					}, 1000);
					return true;
				})
			},
			changePage(currentPage) {
				this.page.page = currentPage
				this.getData()
			},
		},
		mounted() {
			this.getData()
		},
	}
</script>

<style>
</style>
