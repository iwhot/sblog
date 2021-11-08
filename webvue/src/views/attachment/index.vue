<template>
	<div class="file-container">
		<el-row :gutter="20">
			<el-col :span="6">
				<div class="ibox-content">

					<el-upload class="upload-bar" :action="upload" :show-file-list="false" :headers="headers"
						:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
						<el-button size="small" class="btn"><i class="el-icon-upload el-icon--left"></i>上传附件
						</el-button>
					</el-upload>



					<div class="hr-line-dashed"></div>
					<h5>文件夹</h5>
					<ul class="folder-list">
						<li @click="getData('all')"><i class="el-icon-folder"></i><span>全部</span></li>
						<li @click="getData('picture')"><i class="el-icon-folder"></i><span>图片</span></li>
						<li @click="getData('video')"><i class="el-icon-folder"></i><span>视频</span></li>
						<li @click="getData('document')"><i class="el-icon-folder"></i><span>文档</span></li>
						<li @click="getData('excel')"><i class="el-icon-folder"></i><span>表格</span></li>
						<li @click="getData('music')"><i class="el-icon-folder"></i><span>音乐</span></li>
						<li @click="getData('other')"><i class="el-icon-folder"></i><span>其他</span></li>
					</ul>
				</div>
			</el-col>
			<el-col :span="18">
				
				<div class="file-container">
					<div v-for="(item,index) in fileList" :key="index" class="file-box">
						<el-popconfirm title="确认要删除？删除后不可恢复" confirmButtonText='确定' cancelButtonText='取消'
							@confirm="del(item.id)">
							<template slot="reference">
								
								<img :src=" item.url" class="img">
								
							</template>
						</el-popconfirm>
					</div>
				</div>
				
				<div class="file-page">
					<el-pagination background layout="prev, pager, next" :total="page.count" :page-size="page.pageSize"
						:current-page="page.page" @current-change="changePage">
					</el-pagination>
				</div>

			</el-col>
		</el-row>

		<!-- <el-dialog :title="formTitle" :visible.sync="dialogFormVisible" destroy-on-close width="45%"
			custom-class="my-dialog" :close-on-click-modal="false">
			<el-form ref="form" :model="form" label-width="100px">
				
			</el-form>
		</el-dialog> -->
	</div>
</template>

<script>
	import {
		index,
		upload,
		del,
		download
	} from "@/api/attachment/index.js"
	import {
		getToken
	} from "@/utils/auth.js"

	export default {
		name: "AttachmentIndex",
		data() {
			return {
				dialogFormVisible: false,
				formTitle: "编辑",
				fileList: [],
				upload: process.env.VUE_APP_BASE_UPLOAD,
				page: {
					count: 0,
					page: 1,
					pageSize: 15
				},
				// form:{

				// }
			}
		},
		methods: {
			getData(mold) {
				index(mold, {
					page: this.page.page,
					page_size: this.page.pageSize
				}).then(res => {
					// console.log(res.data.data);
					this.fileList = res.data.data.list
					this.page.count = res.data.data.page.count
					this.page.page = res.data.data.page.page
					this.page.pageSize = res.data.data.page.page_size
				})
			},
			edit(row) {
				// console.log(row)
				// this.dialogFormVisible = true
			},
			del(id) {
				del({
					id: id
				}).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						return false;
					}
					this.getData("all")
					this.$message({
						message: res.data.msg,
						type: 'success'
					});

				})
			},
			handleAvatarSuccess(res, file) {
				this.getData("all")
			},
			beforeAvatarUpload(file) {
				//暂时先只能上传图片，然后慢慢完善
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
			//分页数据
			changePage(currentPage) {
				this.page.page = currentPage
				this.getData()
			},
		},
		mounted() {
			this.getData("all")
		},
		computed: {
			headers() {
				return {
					Authorization: getToken()
				}
			}
		},

	}
</script>

<style lang="scss" scoped="scoped">
	.btn {
		width: 100%;
	}

	.ibox-content {
		border: 1px solid #e7eaec;
		background-color: #FFFFFF;
		padding: 20px;
	}

	h5 {
		margin: 0;
		padding: 0;
	}

	.folder-list {
		padding: 0;

		li {
			list-style: none;
			border-bottom: 1px dashed #e7eaec;
			display: block;
			line-height: 30px;
			font-size: 14px;
			color: #666666;
			cursor: pointer;

			span {
				padding-left: 10px;
			}
		}
	}

	.hr-line-dashed {
		border-top: 1px dashed #e7eaec;
		color: #ffffff;
		background-color: #ffffff;
		height: 1px;
		margin: 20px 0;
	}


	.file-box {
		width: 198px;
		height: 158px;
		border: 1px solid #e7eaec;
		margin: 0 20px 20px 0;
		background-color: #FFFFFF;
		float: left;
		overflow: hidden;
		// position: absolute;
		border-radius: 5px;
		cursor: pointer;

		.img {
			width: 100%;
			object-fit: cover;
			position: relative;
			left: 50%;
			top: 50%;
			transform: translate3d(-50%, -50%, 0);
			-webkit-transform: translate3d(-50%, -50%, 0);
		}

	}
	
	.upload-bar /deep/ .el-upload{
		width: 100% !important;
	}
	
	.file-container{
		width: 100%;
		overflow: hidden;
		clear: both;
	}
	
	.file-page{
		width: 100%;
	}
	
	

</style>
