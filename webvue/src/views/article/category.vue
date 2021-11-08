<template>
	<div class="common-container">
		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="addCategory">添加</el-button>
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

					<el-table-column prop="name" label="标题" width="180">
					</el-table-column>

					<el-table-column prop="name" label="缩略图" width="180">
						<template slot-scope="scope">
							<el-image style="width: 50px; height: 50px" :src=" scope.row.thumb"
								:preview-src-list="[ scope.row.thumb]">
							</el-image>
						</template>
					</el-table-column>


					<el-table-column prop="key" label="地址" width="180">
					</el-table-column>

					<el-table-column prop="desc" label="描述"></el-table-column>

					<el-table-column prop="amount" label="状态" align="center" show-overflow-tooltip>
						<template slot-scope="scope">
							<el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
								active-color="#5B7BFA" inactive-color="#dadde5" @click.native="updateStatus(scope.row)">
							</el-switch>
						</template>
					</el-table-column>


					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="editCategory(scope.row)">编辑</el-button>

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

					<el-form-item label="分类名称" prop="name">
						<el-input autocomplete="off" v-model="form.name" placeholder="请输入分类名称"></el-input>
					</el-form-item>

					<el-form-item label="路径" prop="url">
						<el-input autocomplete="off" v-model="form.key" placeholder="请输入分类路径"></el-input>
					</el-form-item>

					<el-form-item label="上级分类" prop="pid">

						<el-select v-model="form.pid" placeholder="请选择">
							<el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id">
							</el-option>
						</el-select>

					</el-form-item>

					<el-form-item label="分类图片" prop="avatar">
						<el-upload class="avatar-uploader" :action="upload" :show-file-list="false" :headers="headers"
							:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
							<img v-if="imageUrl" :src="imageUrl" class="avatar">
							<i v-else class="el-icon-plus avatar-uploader-icon"></i>
						</el-upload>
					</el-form-item>

					<el-form-item label="排序" prop="sort">
						<el-input autocomplete="off" v-model="form.sort" placeholder="请输入排序"></el-input>
					</el-form-item>

					<el-form-item label="状态" prop="status">
						<el-switch v-model="form.status" :active-value="1" :inactive-value="0" active-color="#5B7BFA"
							inactive-color="#dadde5">
						</el-switch>
					</el-form-item>

					<el-form-item label="分类描述" prop="desc">
						<el-input autocomplete="off" v-model="form.desc" type="textarea" rows="5" placeholder="请输入分类描述">
						</el-input>
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
		getToken
	} from "@/utils/auth.js"
	import {
		categoryIndex,
		categorySave,
		categoryDelete,
		GetParentList,
		CategoryUpdateInfo
	} from "@/api/article/index.js"
	export default {
		name: "ArticleCategory",
		data() {
			return {
				dialogFormVisible: false, //显示隐藏表单
				form: {
					id: 0,
					status: 1,
					pid: 0,
					key: "",
					name: "",
					sort: 1000,
					desc: "",
					thumb: "",
				},
				formTitle: "添加",
				options: [], //菜单
				tableData: [], //表格数据
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: '',
			}
		},
		methods: {
			//懒加载下级
			load(tree, treeNode, resolve) {
				categoryIndex({
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

				categorySave(this.form).then(res => {
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
			addCategory() {
				this.form = {
					id: 0,
					status: 1,
					pid: 0,
					key: "",
					name: "",
					sort: 1000,
					desc: "",
					thumb: ""
				}

				this.imageUrl = ""

				if (this.form.pid == 0) {
					this.form.pid = ""
				}

				this.dialogFormVisible = true
				this.getTopCategory()
			},
			//修改菜单
			editCategory(row) {
				this.form = {
					id: row.id,
					status: row.status,
					pid: row.pid,
					key: row.key,
					name: row.name,
					sort: row.sort,
					desc: row.desc,
					thumb: row.thumb
				}

				if (row.thumb != "") {
					this.imageUrl = row.thumb
				}

				if (this.form.pid == 0) {
					this.form.pid = ""
				}

				this.getTopCategory()
				this.dialogFormVisible = true
			},

			//获取列表数据
			getData() {
				categoryIndex({}).then(res => {
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
				categoryDelete(id, {}).then(res => {
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
				CategoryUpdateInfo({
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
				CategoryUpdateInfo({
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
			getTopCategory() {
				GetParentList({}).then(res => {
					this.options = res.data.data
				})
			},
			handleAvatarSuccess(res, file) {
				this.form.thumb = res.data
				this.imageUrl = URL.createObjectURL(file.raw);
			},
			beforeAvatarUpload(file) {
				const isJPG = file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/gif';
				const isLt2M = file.size / 1024 / 1024 < 2;

				if (!isJPG) {
					this.$message.error('上传头像图片只能是 JPG、PNG、GIF 格式!');
				}

				if (!isLt2M) {
					this.$message.error('上传头像图片大小不能超过 2MB!');
				}
				return isJPG && isLt2M;
			},
		},
		mounted() {
			this.getData()
		},
		computed: {
			headers() {
				return {
					Authorization: getToken()
				}
			}
		}
	}
</script>

<style>
</style>
