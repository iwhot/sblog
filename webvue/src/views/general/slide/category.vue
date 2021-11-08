<template>
	<div class="common-container">

		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="add">添加</el-button>

			</div>
			<div class="table-container">
				<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%"
					:header-cell-style="{background:'#f6f6f6',color:'#606266'}" :border="true" :key="Math.random()">

					<el-table-column fixed="left" label="ID" prop="id" width="80"></el-table-column>

					<el-table-column prop="title" label="名称"></el-table-column>
					<el-table-column prop="attr" label="属性">
						<template slot-scope="scope">
							<span v-if="scope.row.attr == 0">幻灯</span>
							<span v-if="scope.row.attr == 1">单图</span>
							<span v-if="scope.row.attr == 2">文本</span>
						</template>
					</el-table-column>
					<el-table-column prop="width" label="宽度"></el-table-column>
					<el-table-column prop="height" label="高度"></el-table-column>
					<el-table-column prop="remark" label="描述"></el-table-column>
					<el-table-column prop="createtime" label="创建时间"></el-table-column>
					<el-table-column prop="updatetime" label="更新时间"></el-table-column>
					<el-table-column prop="banner" label="幻灯管理">
						<template slot-scope="scope">
							 <el-button plain @click="banner(scope.row.id)">幻灯管理</el-button>
						</template>
					</el-table-column>

					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="edit(scope.row)">编辑
							</el-button>

							<!-- <el-popconfirm title="确认要删除？删除后不可恢复" confirmButtonText='确定' cancelButtonText='取消'
								@confirm="del(scope.row.id)">
								<template slot="reference">
									<el-button type="text" size="small">
										删除
									</el-button>
								</template>
							</el-popconfirm> -->

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
					<el-form-item label="名称" prop="title">
						<el-input autocomplete="off" v-model="form.title" placeholder="请输入名称"></el-input>
					</el-form-item>

					<el-form-item label="属性">
						<el-radio-group v-model="form.attr">
							<el-radio :label="0">幻灯</el-radio>
							<el-radio :label="1">单图</el-radio>
							<el-radio :label="2">文本</el-radio>
						</el-radio-group>
					</el-form-item>

					<el-form-item label="宽度" prop="width">
						<el-input autocomplete="off" type="number" v-model="form.width" placeholder="请输入宽度"></el-input>
					</el-form-item>

					<el-form-item label="高度" prop="height">
						<el-input autocomplete="off" type="number" v-model="form.height" placeholder="请输入高度"></el-input>
					</el-form-item>

					<el-form-item label="描述" prop="remark">
						<el-input autocomplete="off" type="textarea" v-model="form.remark" rows="5"></el-input>
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
		indexSlideCategory,
		saveSlideCategory,
		deleteSlideCategory,
	} from '@/api/general/slide.js'

	export default {
		name: "GeneralSlideCategory",
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
					title: "",
					width: 0,
					height: 0,
					remark: "",
					attr: 0,
				},
			}
		},
		methods: {
			//获取数据
			getData() {
				indexSlideCategory({
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

				deleteSlideCategory({
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
			edit(row) {
				// console.log(row)
				this.dialogFormVisible = true
				this.formTitle = "编辑"

				this.form.title = row.title
				this.form.width = row.width
				this.form.id = row.id
				this.form.height = row.height
				this.form.remark = row.remark
				this.form.attr = row.attr
			},
			//更新信息
			update(row) {
				this.form.title = row.title
				this.form.width = row.width
				this.form.id = row.id
				this.form.height = row.height
				this.form.remark = row.remark
				this.form.attr = row.attr
				this.save()
			},
			add() {
				this.dialogFormVisible = true
				this.formTitle = "添加"

				this.form.title = ""
				this.form.width = 0
				this.form.height = 0
				this.form.id = 0
				this.form.remark = ""
				this.form.attr = 0
			},
			save() {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				//需要转码一下要不然后端认为是字符串会报错
				this.form.id = parseInt(this.form.id)
				this.form.width = parseInt(this.form.width)
				this.form.height = parseInt(this.form.height)
				this.form.attr = parseInt(this.form.attr)
				saveSlideCategory(this.form).then(res => {
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
			banner(id){
				//跳转到编辑页面
				let routeDate = this.$router.push({
					path: '/general/slide/banner/' + id,
				})
				window.open(routeDate.href, '_top');
			}
		},
		mounted() {
			this.getData()
		},
	}
</script>

<style>
</style>
