<template>
	<div class="common-container">
		<div style="padding: 0 0 10px 0;">
			<span style="color: red;">分类名称：{{title}}</span>
		</div>
		<div class="table">
			<div class="toolbar">
				<el-button type="primary" size="mini" @click="add">添加</el-button>

			</div>
			<div class="table-container">
				<el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%"
					:header-cell-style="{background:'#f6f6f6',color:'#606266'}" :border="true" :key="Math.random()">

					<el-table-column fixed="left" label="ID" prop="id" width="80"></el-table-column>

					<el-table-column prop="sort" label="排序" width="100">
						<template slot-scope="scope">
							<el-input v-model="scope.row.sort" size="mini" @change="update(scope.row)"></el-input>
						</template>
					</el-table-column>

					<el-table-column prop="thumb" label="缩略图">
						<template slot-scope="scope">
							<el-image style="width: 50px; height: 50px" :src=" scope.row.thumb"
								:preview-src-list="[ scope.row.thumb]">
							</el-image>
						</template>
					</el-table-column>
					<el-table-column prop="title" label="标题"></el-table-column>
					<el-table-column prop="remark" label="描述"></el-table-column>
					<el-table-column prop="createtime" label="创建时间"></el-table-column>
					<el-table-column prop="updatetime" label="更新时间"></el-table-column>
					<el-table-column prop="status" label="状态" align="center" show-overflow-tooltip>
						<template slot-scope="scope">
							<el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
								active-color="#5B7BFA" inactive-color="#dadde5" @click.native="update(scope.row)">
							</el-switch>
						</template>
					</el-table-column>

					<el-table-column fixed="right" label="操作" width="100">
						<template slot-scope="scope">
							<el-button type="text" size="small" @click="edit(scope.row)">编辑
							</el-button>

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

				<el-pagination background layout="prev, pager, next" :total="page.count" :page-size="page.pageSize"
					:current-page="page.page" @current-change="changePage">
				</el-pagination>
			</div>
		</div>

		<el-dialog :title="formTitle" :visible.sync="dialogFormVisible" destroy-on-close width="45%"
			custom-class="my-dialog" :close-on-click-modal="false">
			<el-scrollbar wrap-class="form-wrapper">
				<el-form ref="form" :model="form" label-width="100px">

					<el-form-item label="标题" prop="title">
						<el-input autocomplete="off" v-model="form.title" placeholder="请输入标题"></el-input>
					</el-form-item>

					<el-form-item label="缩略图" prop="thumb">
						<el-upload class="avatar-uploader" :action="upload" :show-file-list="false" :headers="headers"
							:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
							<img v-if="imageUrl" :src="imageUrl" class="avatar">
							<i v-else class="el-icon-plus avatar-uploader-icon"></i>
						</el-upload>
					</el-form-item>

					<el-form-item label="地址" prop="url">
						<el-input autocomplete="off" v-model="form.url" placeholder="请输入地址"></el-input>
					</el-form-item>
					
					<el-form-item label="状态" prop="status">
						<el-switch v-model="form.status" :active-value="1" :inactive-value="0" active-color="#5B7BFA"
							inactive-color="#dadde5">
						</el-switch>
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
		getToken
	} from "@/utils/auth.js"
	import {
		findOneSlideCategory,
		indexSlideBanner,
		saveSlideBanner,
		deleteSlideBanner
	} from '@/api/general/slide.js'
	export default {
		name: "GeneralSlideBanner",
		computed: {
			headers() {
				return {
					Authorization: getToken()
				}
			}
		},
		data() {
			return {
				title: "",
				cid: 0,
				showAll2: true,
				dialogFormVisible: false,
				word: "展开搜索",
				tips: "操作提示",
				tableData: [],
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: '',
				page: {
					count: 0,
					page: 1,
					pageSize: 8
				},
				formTitle: "添加",
				form: {
					id: 0,
					title: "",
					thumb: "",
					url: "",
					remark: "",
					status: 1,
					sort: 1000,
					cid: 0,
				},
			}
		},
		methods: {
			//获取数据
			getData() {
				// console.log(this.cid)
				indexSlideBanner(this.cid,{
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

				deleteSlideBanner({
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

				this.form.id = row.id
				this.form.title = row.title
				this.form.remark = row.remark
				this.form.thumb = row.thumb
				this.form.url = row.url
				this.form.status = row.status
				this.form.cid = row.cid
				this.form.sort = row.sort
				
				this.imageUrl = row.thumb
			},
			//更新信息
			update(row) {
				this.form.id = row.id
				this.form.title = row.title
				this.form.remark = row.remark
				this.form.thumb = row.thumb
				this.form.url = row.url
				this.form.status = row.status
				this.form.cid = row.cid
				this.form.sort = row.sort
				
				this.save()
			},
			add() {
				this.dialogFormVisible = true
				this.formTitle = "添加"
				this.imageUrl = ""
				
				this.form.id = 0
				this.form.title = ""
				this.form.remark = ""
				this.form.thumb = ""
				this.form.url = ""
				this.form.status = 1
				this.form.cid = this.cid
				this.form.sort = 1000
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
				this.form.status = parseInt(this.form.status)
				this.form.cid = parseInt(this.form.cid)
				this.form.sort = parseInt(this.form.sort)
				saveSlideBanner(this.form).then(res => {
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
			banner(id) {
				//跳转到编辑页面
				let routeDate = this.$router.push({
					path: '/general/slide/banner/' + id,
				})
				window.open(routeDate.href, '_top');
			},
			findOneSlideCategory() {
				let id = this.$route.params.cid
				findOneSlideCategory(id, {}).then(res => {
					this.title = res.data.data.title
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
			this.findOneSlideCategory()
			this.cid = this.$route.params.cid
			this.form.cid = this.$route.params.cid
			this.getData()
		},
	}
</script>

<style>
</style>
