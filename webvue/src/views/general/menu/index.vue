<template>
	<div class="common-container">
		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="addMenu">添加</el-button>
			</div>

			<div class="table-container">
				<el-table :data="tableData" style="width: 100%" row-key="id" border lazy :load="load"
					:tree-props="{children: 'children', hasChildren: 'hasChildren'}"
					:header-cell-style="{background:'#f6f6f6',color:'#606266'}">
					<el-table-column label="ID" fixed="left" prop="id" width="80"></el-table-column>

					<el-table-column prop="address" label="排序" width="100">
						<template slot-scope="scope">
							<el-input v-model="scope.row.sort" size="mini" @change="updateSort(scope.row)"></el-input>
						</template>
					</el-table-column>

					<el-table-column prop="title" label="标题" width="180">
					</el-table-column>
					<el-table-column prop="group" label="分组" width="180">
						<template slot-scope="scope">
							<span v-if="scope.row.group == 'top'">顶部</span>
							<span v-if="scope.row.group == 'center'">中间</span>
							<span v-if="scope.row.group == 'footer'">底部</span>
							<span v-if="scope.row.group == 'other'">其他</span>
						</template>
					</el-table-column>
					<el-table-column prop="url" label="地址" width="180">
					</el-table-column>

					<el-table-column prop="amount" label="状态" align="center" show-overflow-tooltip>
						<template slot-scope="scope">
							<el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
								active-color="#5B7BFA" inactive-color="#dadde5" @click.native="updateStatus(scope.row)">
							</el-switch>
						</template>
					</el-table-column>


					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="editMenu(scope.row)">编辑</el-button>

							<el-popconfirm title="确认要删除？" confirmButtonText='确定' cancelButtonText='取消'
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
			</div>
		</div>

		<el-dialog :title="formTitle" :visible.sync="dialogFormVisible" destroy-on-close width="45%"
			custom-class="my-dialog" :close-on-click-modal="false">
			<el-scrollbar wrap-class="form-wrapper">
				<el-form ref="form" label-width="100px">

					<el-form-item label="标题" prop="title">
						<el-input autocomplete="off" v-model="form.title" placeholder="请输入标题"></el-input>
					</el-form-item>
					<el-form-item label="位置" prop="group">
						<el-select v-model="form.group" placeholder="请选择">
							<el-option label="顶部" value="top"></el-option>
							<el-option label="中间" value="center"></el-option>
							<el-option label="底部" value="footer"></el-option>
							<el-option label="其他" value="other"></el-option>
						</el-select>
					</el-form-item>
					<el-form-item label="上级" prop="pid">
						<el-select v-model="form.pid" placeholder="请选择">
							<el-option label="顶级菜单" value="0"></el-option>
							<el-option v-for="item in options" :key="item.id" :label="item.title" :value="item.id">
							</el-option>
						</el-select>

					</el-form-item>

					<el-form-item label="地址" prop="url">
						<el-input autocomplete="off" v-model="form.url" placeholder="请输入地址"></el-input>
					</el-form-item>

					<el-form-item label="打开方式">
						<el-radio-group v-model="form.target">
							<el-radio label="_self">当前页</el-radio>
							<el-radio label="_target">新页面</el-radio>
						</el-radio-group>
					</el-form-item>

					<el-form-item label="排序" prop="sort">
						<el-input autocomplete="off" v-model="form.sort" placeholder="请输入排序"></el-input>
					</el-form-item>

					<el-form-item label="状态" prop="status">
						<el-switch v-model="form.status" :active-value="1" :inactive-value="0" active-color="#5B7BFA"
							inactive-color="#dadde5">
						</el-switch>
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
		MenuIndex,
		MenuSave,
		MenuDelete,
		GetParentList,
		MenuUpdateInfo
	} from "@/api/general/menu.js"
	export default {
		name: "GeneralMenuIndex",
		data() {
			return {
				dialogFormVisible: false, //显示隐藏表单
				form: {
					id: 0,
					status: 1,
					pid: '0',
					url: "",
					title: "",
					sort: 1000,
					target: "_self",
					group: "top",
				},
				formTitle: "添加",
				options: [], //菜单
				tableData: [], //表格数据
			}
		},
		methods: {
			//懒加载下级
			load(tree, treeNode, resolve) {
				MenuIndex({
					id: tree.id
				}).then(res => {
					resolve(res.data.data)
				})
			},
			//保存
			save(refs) {

				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});


				if (this.form.pid == '') {
					this.form.pid = 0
				}

				this.form.sort = parseInt(this.form.sort)
				this.form.pid = parseInt(this.form.pid)
				this.form.status = parseInt(this.form.status)
				this.form.id = parseInt(this.form.id)
				MenuSave(this.form).then(res => {
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
						// window.location.reload()
					}, 1000);
					return true;
				})
			},

			//添加菜单
			addMenu() {
				this.form = {
					id: 0,
					status: 1,
					pid: 0,
					url: "",
					titlr: "",
					sort: 1000,
					target: "_self",
					group:"top"
				}
				if (this.form.pid == 0) {
					this.form.pid = ""
				}

				this.dialogFormVisible = true
				this.getTopMenu()
			},
			//修改菜单
			editMenu(row) {
				this.form = {
					id: row.id,
					status: row.status,
					pid: row.pid,
					url: row.url,
					title: row.title,
					sort: row.sort,
					target: row.target,
					group: row.group,
				}

				if (this.form.pid == 0) {
					this.form.pid = ""
				}

				this.getTopMenu()
				this.dialogFormVisible = true
			},

			//获取列表数据
			getData() {
				MenuIndex({}).then(res => {
					this.tableData = res.data.data
				})
			},
			//删除
			del(id) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});
				MenuDelete(id, {}).then(res => {
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
						// window.location.reload()
					}, 1000);
					return true;
				})
			},
			//修改状态
			updateStatus(row) {
				MenuUpdateInfo({
					id: row.id,
					status: row.status
				}).then(res => {
					this.getData()
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
				})
			},
			//修改排序
			updateSort(row) {
				MenuUpdateInfo({
					id: row.id,
					sort: row.sort
				}).then(res => {
					this.getData()
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
					// window.location.reload()
				})
			},
			getTopMenu() {
				GetParentList({}).then(res => {
					this.options = res.data.data
				})
			}
		},
		mounted() {
			this.getData()
		},
	}
</script>

<style>
</style>
