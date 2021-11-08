<template>
	<div class="common-container">
		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="add">添加</el-button>

				<el-popconfirm title="确认要删除？删除后不可恢复" confirmButtonText='确定' cancelButtonText='取消' @confirm="delAll">
					<template slot="reference">
						<el-button type="danger" size="mini" style="margin-left: 10px;">删除</el-button>
					</template>
				</el-popconfirm>


			</div>
			<div class="table-container">
				<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%"
					:header-cell-style="{background:'#f6f6f6',color:'#606266'}" :border="true" :key="Math.random()"
					@selection-change="selectionLineChangeHandle">
					<el-table-column fixed="left" type="selection" width="55" :selectable="selectable">
					</el-table-column>

					<el-table-column label="ID" prop="id" width="80"></el-table-column>

					<el-table-column prop="address" label="排序" width="100">
						<template slot-scope="scope">
							<el-input v-model="scope.row.sort" size="mini" @change="update(scope.row)"></el-input>
						</template>
					</el-table-column>

					<el-table-column prop="avatar" label="头像">
						<template slot-scope="scope">
							<img :src=" scope.row.avatar" width="60">
						</template>
					</el-table-column>

					<el-table-column prop="username" label="名称"></el-table-column>
					<el-table-column prop="nickname" label="昵称"></el-table-column>
					<el-table-column prop="email" label="邮箱"></el-table-column>

					<el-table-column prop="logintime" label="登陆时间"></el-table-column>
					<el-table-column prop="loginip" label="登陆IP"></el-table-column>

					<el-table-column prop="amount" label="状态" align="center" show-overflow-tooltip>
						<template slot-scope="scope">
							<el-switch v-if="scope.row.id > 1" v-model="scope.row.status" :active-value="1"
								:inactive-value="0" active-color="#5B7BFA" inactive-color="#dadde5"
								@click.native="update(scope.row)">
							</el-switch>
						</template>
					</el-table-column>

					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>

							<el-popconfirm v-if="scope.row.id > 1" title="确认要删除？删除后不可恢复" confirmButtonText='确定'
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
					<el-form-item label="所属组别" prop="group">
						<el-select v-model="form.group" multiple placeholder="请选择">
							<el-option v-for="item in options" :key="item.value" :label="item.label"
								:value="item.value">
							</el-option>
						</el-select>
					</el-form-item>

					<el-form-item label="　用户名" prop="username">
						<el-input autocomplete="off" v-model="form.username" :disabled="disableUsername"
							placeholder="请输入用户名">
						</el-input>
					</el-form-item>

					<el-form-item label="　　昵称" prop="nickname">
						<el-input autocomplete="off" v-model="form.nickname" placeholder="请输入昵称"></el-input>
					</el-form-item>

					<el-form-item label="　　密码" prop="password">
						<el-input autocomplete="off" :type=" flag ? 'text' : 'password' " v-model="form.password"
							placeholder="请输入密码">
							<template slot="suffix">
								<i :class="flag ? 'el-input__icon el-icon-more' : 'el-input__icon el-icon-view' "
									style="cursor: pointer;" @click="flag = !flag"></i>
							</template>
						</el-input>
					</el-form-item>

					<el-form-item label="　　头像" prop="avatar">
						<el-upload class="avatar-uploader" :action="upload" :show-file-list="false" :headers="headers"
							:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
							<img v-if="imageUrl" :src="imageUrl" class="avatar">
							<i v-else class="el-icon-plus avatar-uploader-icon"></i>
						</el-upload>
					</el-form-item>

					<el-form-item label="　　邮箱" prop="email">
						<el-input autocomplete="off" v-model="form.email" placeholder="请输入邮箱"></el-input>
					</el-form-item>


					<el-form-item label="　　状态" prop="status">
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
		getToken
	} from "@/utils/auth.js"

	import {
		getAllGroup,
		saveAdmin,
		indexAdmin,
		delAdmin,
		delAllAdmin,
		GetMyGroupID
	} from '@/api/auth/index.js'



	export default {
		name: "AuthAdmin",
		data() {
			return {
				title: "首页",
				dialogFormVisible: false,
				disableUsername: false,
				flag: false,
				tableData: [],
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: '',
				options: [],
				page: {
					count: 0,
					page: 1,
					pageSize: 8
				},
				formTitle: "添加",
				form: {
					id: 0,
					username: "",
					nickname: "",
					password: "",
					avatar: "",
					email: "",
					status: 1,
					group: []
				},
				selects: "",
			}
		},
		methods: {
			//添加
			add() {
				this.form = {
					id: 0,
					username: "",
					nickname: "",
					password: "",
					avatar: "",
					email: "",
					status: 1,
					group: []
				}
				this.disableUsername = false
				this.getAllGroup()
				this.dialogFormVisible = true
				this.imageUrl = ''
			},
			//编辑
			edit(row) {
				this.form = {
					id: row.id,
					username: row.username,
					nickname: row.nickname,
					password: "",
					avatar: row.avatar,
					email: row.email,
					status: row.status,
					group: []
				}
				this.disableUsername = true
				if(row.avatar){
					this.imageUrl = row.avatar
				}
				
				GetMyGroupID({}).then( res => {
					this.form.group = res.data.data
				})
				
				// console.log(row)
				this.getAllGroup()
				this.dialogFormVisible = true
			},
			//保存
			save() {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				saveAdmin(this.form).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						setTimeout(() => {
							loading.close();
						}, 1000);
						return false;
					}

					this.getData()

					//更新store
					this.$store.dispatch("user/refresh_user")

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
			//选中
			selectable(row, index) {
				// console.log(row,index)
				if (row.id == 1) {
					return false
				} else {
					return true
				}
			},
			//选中的值
			selectionLineChangeHandle(val) {
				var selects = []
				val.forEach((item, index) => {
					selects.push(item.id)
				})
				this.selects = selects.join(",")
			},
			//修改密码框类型
			changeType() {
				this.flag = !this.flag
			},
			handleAvatarSuccess(res, file) {
				this.form.avatar = res.data
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
			//获取所有组
			getAllGroup() {
				getAllGroup({}).then(res => {
					this.options = res.data.data
				})
			},
			//分页数据
			changePage(currentPage) {
				this.page.page = currentPage
				this.getData()
			},
			//获取内容
			getData() {
				indexAdmin({
					page: this.page.page,
					page_size: this.page.pageSize
				}).then(res => {
					// console.log(res.data)
					this.tableData = res.data.data.list
					this.page.count = res.data.data.page.count
					this.page.page = res.data.data.page.page
					this.page.pageSize = res.data.data.page.page_size
				})
			},
			//更新信息
			update(row) {
				this.form = {
					id: row.id,
					username: row.username,
					nickname: row.nickname,
					password: "",
					avatar: row.avatar,
					email: row.email,
					status: row.status,
					group: row.group
				}
				this.save()
			},
			//删除
			del(id) {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)'
				});

				delAdmin({
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
			//删除所有
			delAll() {
				// console.log(this.selects)
				delAllAdmin({
					id: this.selects
				}).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						return false;
					}

					this.getData()

					this.$message({
						message: res.data.msg,
						type: 'success'
					});
				})
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
		computed: {
			headers() {
				return {
					Authorization: getToken()
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

<style>

</style>
