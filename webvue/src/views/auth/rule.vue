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
					<el-table-column prop="url" label="路由地址" width="180">
					</el-table-column>
					<el-table-column prop="component" label="组件路径" width="180">
					</el-table-column>

					<el-table-column prop="redirect" label="跳转路径" width="180">
					</el-table-column>

					<el-table-column prop="icon" label="图标">
						<template slot-scope="scope">
							<svg-icon :icon-class="scope.row.icon && scope.row.icon" class-name="svg-font"></svg-icon>
						</template>
					</el-table-column>

					<el-table-column prop="amount" label="状态" align="center" show-overflow-tooltip>
						<template slot-scope="scope">
							<el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
								active-color="#5B7BFA" inactive-color="#dadde5" @click.native="updateStatus(scope.row)">
							</el-switch>
						</template>
					</el-table-column>

					<el-table-column prop="remark" label="描述" width="180">
					</el-table-column>

					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="editMenu(scope.row)">编辑</el-button>

							<el-popconfirm v-if="scope.row.is_not_del == 0" title="父菜单删除子菜单也会跟着删除,确认要删除？"
								confirmButtonText='确定' cancelButtonText='取消' @confirm="del(scope.row.id)">
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
					<!-- <el-form-item label="所属模块" prop="mod_id">
					<el-select v-model="form.mod_id" placeholder="请选择上级菜单">
						<el-option v-for="item in moduleList" :label="item.mod_name" :value="item.id" :key="item.id">
						</el-option>
					</el-select>
				</el-form-item> -->
					<el-form-item label="是否菜单">
						<el-radio-group v-model="form.is_menu">
							<el-radio :label="1">是</el-radio>
							<el-radio :label="0">否</el-radio>
						</el-radio-group>
					</el-form-item>

					<el-form-item label="菜单名称" prop="title">
						<el-input autocomplete="off" v-model="form.title" placeholder="请输入菜单名称"></el-input>
					</el-form-item>

					<el-form-item label="菜单图标" prop="icon">
						<SvgChoice v-model="form.icon"></SvgChoice>
					</el-form-item>

					<el-form-item label="上级菜单" prop="pid">
						<el-cascader :options="selectValues" :show-all-levels="true" v-model="ap"
							:props="{checkStrictly: true }" clearable @change="getPid"></el-cascader>
					</el-form-item>

					<el-form-item label="路由地址" prop="url">
						<el-input autocomplete="off" v-model="form.url" placeholder="请输入路由地址"
							@keyup.native="changeComponent()">
						</el-input>
					</el-form-item>

					<el-form-item label="组件路径" prop="component">
						<el-input autocomplete="off" v-model="form.component" placeholder="请输入载入路由组件的路径"></el-input>
					</el-form-item>

					<el-form-item label="跳转路径" prop="redirect">
						<el-input autocomplete="off" v-model="form.redirect" placeholder="请输入载入路由跳转的路径"></el-input>
					</el-form-item>

					<el-form-item label="菜单排序" prop="sort">
						<el-input autocomplete="off" v-model="form.sort" placeholder="请输入排序"></el-input>
					</el-form-item>

					<el-form-item label="菜单描述" prop="icon">
						<el-input type="textarea" v-model="form.remark"></el-input>
					</el-form-item>

					<el-form-item label="菜单状态" prop="status">
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
		saveMenu,
		getMenuIndex,
		cascadeMenu,
		delMenu,
		getMenuPid,
		updateMenuStatus,
		updateMenuSort
	} from "@/api/auth/index.js"


	import SvgChoice from '@/components/SvgChoice/index.vue'

	export default {
		name: "AuthRule",
		components: {
			SvgChoice
		},
		data() {
			return {
				dialogFormVisible: false, //显示隐藏表单
				ap: [], //级联菜单默认值
				form: {
					id: 0,
					status: 1,
					pid: 0,
					url: "",
					title: "",
					icon: "",
					component: "",
					mod_id: "",
					remark: "",
					redirect: "",
					sort: 1000,
					is_menu: 1
				},
				formTitle: "添加",
				moduleList: [], //模块列表
				selectValues: [], //级联菜单
				tableData: [], //表格数据
			}
		},
		methods: {
			//懒加载下级
			load(tree, treeNode, resolve) {
				getMenuIndex({
					id: tree.id
				}).then(res => {
					resolve(res.data.data)
				})
			},
			//保存菜单
			save(refs) {

				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				saveMenu(this.form).then(res => {
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
						window.location.reload()
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
					title: "",
					icon: "",
					component: "",
					mod_id: "",
					remark: "",
					redirect: "",
					sort: 1000,
					is_menu: 1
				}
				this.ap = []
				this.cascadeMenu() //获取级联菜单
				this.dialogFormVisible = true
			},
			//修改菜单
			editMenu(row) {
				this.form = {
					id: row.id,
					status: row.status,
					pid: row.pid,
					url: row.url,
					title: row.title,
					icon: row.icon,
					component: row.component,
					mod_id: row.mod_id,
					remark: row.remark,
					redirect: row.redirect,
					sort: row.sort,
					is_menu: row.is_menu
				}

				if (row.mod_id == 0) {
					this.form.mod_id = ""
				}
				
				getMenuPid({
					id: row.pid
				}).then(res => {
					this.ap = res.data.data
				})

				this.cascadeMenu() //获取级联菜单
				this.dialogFormVisible = true
			},
			//修改组件地址
			changeComponent() {
				var com = this.form.url
				if (com.lastIndexOf("/") + 1 === com.length) {
					com = com.substring(0, com.length - 1)
				}
				if (com !== "") {
					this.form.component = com + ".vue"
				} else {
					this.form.component = ""
				}

			},
			//获取列表数据
			getData() {
				getMenuIndex({}).then(res => {
					this.tableData = res.data.data
				})
			},
			//获取级联菜单
			cascadeMenu() {
				cascadeMenu({}).then(res => {
					this.selectValues = res.data.data
				})
			},
			//获取选中的值
			getPid(e) {
				if (e == null) {
					this.form.pid = 0
				} else {
					this.form.pid = e.slice(-1)[0]
				}

			},
			//删除菜单
			del(id) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});
				delMenu({
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
						window.location.reload()
					}, 1000);
					return true;
				})
			},
			//修改状态
			updateStatus(row) {
				updateMenuStatus({
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
				updateMenuSort({
					id: row.id,
					sort: row.sort
				}).then(res => {
					this.getData()
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
					window.location.reload()
				})
			},
			//更新
			update(row) {
				this.form = {
					id: row.id,
					status: row.status,
					pid: row.pid,
					url: row.url,
					title: row.title,
					icon: row.icon,
					component: row.component,
					mod_id: row.mod_id,
					remark: row.remark,
					redirect: row.redirect,
					sort: row.sort
				}
				
				this.save()
			},
			//获取名称
			getName(data) {
				this.form.icon = data
			},
			//监听回车键
			keyupSubmit() {
				document.onkeydown = e => {
					let _key = window.event.keyCode;
					if (_key === 13) {
						this.save('form')
					}
				}
			}
		},
		mounted() {
			this.getData()
		},
		created() {
			this.keyupSubmit()
		}
	}
</script>

<style lang="scss" scoped="scoped">
	.svg-font {
		font-size: 20px;
	}
</style>
