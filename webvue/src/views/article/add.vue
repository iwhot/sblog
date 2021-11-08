<template>
	<div class="tab-container">
		<vuescroll>
			<el-form ref="form" :label-position="labelPosition" label-width="150px">
				<el-tabs type="border-card">
					<el-tab-pane label="基础设置">
						<el-form-item label="文章标题">
							<el-input size="small" v-model="form.title" placeholder="请输入文章标题"></el-input>
						</el-form-item>

						<el-form-item label="所属分类">
							<el-cascader :options="options" :props="{ checkStrictly: true }" v-model="ap" clearable
								@change="GetCascaderValue"></el-cascader>
						</el-form-item>

						<el-form-item label="置顶">
							<el-radio-group v-model="form.is_top">
								<el-radio :label="0">否</el-radio>
								<el-radio :label="1">是</el-radio>
							</el-radio-group>
						</el-form-item>

						<el-form-item label="推荐">
							<el-radio-group v-model="form.is_nice">
								<el-radio :label="0">否</el-radio>
								<el-radio :label="1">是</el-radio>
							</el-radio-group>
						</el-form-item>

						<el-form-item label="首页">
							<el-radio-group v-model="form.is_index">
								<el-radio :label="0">否</el-radio>
								<el-radio :label="1">是</el-radio>
							</el-radio-group>
						</el-form-item>

						<el-form-item label="标签">
							<el-input size="small" v-model="form.tags" placeholder="请输入标签多个用英文逗号隔开"></el-input>
						</el-form-item>

						<el-form-item label="缩略图" prop="thumb">
							<el-upload class="avatar-uploader" :action="upload" :show-file-list="false"
								:headers="headers" :on-success="handleAvatarSuccess"
								:before-upload="beforeAvatarUpload">
								<img v-if="imageUrl" :src="imageUrl" class="avatar">
								<i v-else class="el-icon-plus avatar-uploader-icon"></i>
							</el-upload>
						</el-form-item>

						<el-form-item label="列表图" prop="thumb">
							<el-upload :action="upload" :headers="headers" list-type="picture-card"
								:on-preview="handlePictureCardPreview" :on-remove="handleRemove" :limit="4"
								:file-list="uploadList" :on-success="listSucess">
								<i class="el-icon-plus"></i>
							</el-upload>

							<el-dialog :visible.sync="dialogVisible">
								<img width="100%" :src="dialogImageUrl" alt="">
							</el-dialog>

						</el-form-item>

						<el-form-item label="摘要">
							<el-input size="small" v-model="form.desc" type="textarea" placeholder="请输入文章摘要">
							</el-input>
						</el-form-item>

						<el-form-item label="文章内容">
							<editor v-model="form.content" style="height: 375px; margin-bottom: 60px;">
							</editor>
						</el-form-item>

						<el-form-item label="可见性">
							<el-radio-group v-model="form.status">
								<el-radio :label="0">隐藏</el-radio>
								<el-radio :label="1">展示</el-radio>
							</el-radio-group>
						</el-form-item>

						<el-form-item label="发布">
							<el-radio-group v-model="form.state">
								<el-radio :label="1">草稿箱</el-radio>
								<el-radio :label="2">发布</el-radio>
							</el-radio-group>
						</el-form-item>

						<el-form-item label="排序">
							<el-input size="small" v-model="form.sort" placeholder="请输入排序"></el-input>
						</el-form-item>

					</el-tab-pane>

					<el-tab-pane label="SEO设置">
						<el-form-item label="SEO标题">
							<el-input size="small" v-model="form.seo_title" placeholder="请输入seo标题"></el-input>
						</el-form-item>

						<el-form-item label="SEO关键字">
							<el-input size="small" v-model="form.seo_key" type="textarea" placeholder="请输入seo关键字">
							</el-input>
						</el-form-item>

						<el-form-item label="SEO描述">
							<el-input size="small" v-model="form.seo_desc" type="textarea" rows="10"
								placeholder="请输入seo描述">
							</el-input>
						</el-form-item>
					</el-tab-pane>

					<el-form-item style="margin-bottom: 0;">
						<el-button type="primary" size="mini" @click="saveData()">保存</el-button>
					</el-form-item>
				</el-tabs>

			</el-form>
		</vuescroll>
	</div>
</template>

<script>
	import {
		CategoryCascader,
		ArticleSave
	} from "@/api/article/index.js"
	import {
		getToken
	} from "@/utils/auth.js"

	import editor from '@/components/editor/index.vue'
	import vuescroll from 'vuescroll';
	export default {
		name: "ArticleAdd",
		components: {
			editor,
			vuescroll
		},
		computed: {
			headers() {
				return {
					Authorization: getToken()
				}
			}
		},
		data() {
			return {
				labelPosition: 'center',
				contentPosition: 'center',
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: '',
				dialogImageUrl: '',
				dialogVisible: false,
				uploadList: [], //编辑的时候用得到，添加不用
				form: {
					cid: 0,
					title: "",
					desc: "",
					content: "",
					status: 1,
					is_top: 0,
					is_nice: 0,
					is_index: 0,
					sort: 1000,
					state: 2,
					thumb: "",
					thumb_ext: "",
					seo_title: "",
					seo_desc: "",
					seo_key: "",
					tags: "",
				},
				options: [],
				ap: []
			};
		},
		methods: {
			//级联分类
			Cascader() {
				CategoryCascader({}).then(res => {
					this.options = res.data.data
				})
			},
			//获取选中的值
			GetCascaderValue(e) {
				if (e == null) {
					this.form.cid = 0
				} else {
					this.form.cid = e.slice(-1)[0]
				}

			},
			//保存数据
			saveData() {
				const loading = this.$loading({
					lock: true,
					text: 'Loading',
					spinner: 'el-icon-loading',
					background: 'rgba(0, 0, 0, 0.7)',
				});

				let BOX = document.querySelector(".el-loading-mask");
				BOX.style.zIndex = "999999"; //固定一个最大值，是字符串

				ArticleSave(this.form).then(res => {
					if (res.data.code != 200) {
						loading.close();
						this.$message.error(res.data.msg);
						return false;
					}

					this.$message({
						message: res.data.msg,
						type: 'success'
					});

					setTimeout(() => {
						loading.close();
						//关闭当前页面
						var visitedViews = this.$store.state.tagsView.visitedViews
						this.$store.dispatch('tagsView/del_one_visited_view', this.$route).then(() => {
							const latestView = visitedViews.slice(-2)[0]
							if (latestView) {
								this.$router.push(latestView.fullPath) //跳转页面
							}
						})

					}, 1000);
					return true;
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
			handleRemove(file, fileList) {
				if (fileList.length <= 0) {
					this.form.thumb_ext = ""
				} else {
					var lst = []
					fileList.forEach(function(item,index){
						lst.push(item.url)
					})
					this.form.thumb_ext = lst.join(",")
				}
			},
			handlePictureCardPreview(file) {
				this.dialogImageUrl = file.url;
				this.dialogVisible = true;
			},
			listSucess(res, file) {
				this.form.thumb_ext += "," + res.data
				this.form.thumb_ext = this.form.thumb_ext.replace(/^,+/,"").replace(/,+$/,"")//去除两端逗号
				// console.log(this.form.thumb_ext)
			}
		},
		mounted() {
			//载入级联分类
			this.Cascader()
		}
	}
</script>

<style lang="scss" scoped="scoped">
	.tab-container {
		width: 100%;
		height: 100%;
		background-color: #f5f7f9;

		.gray {
			background-color: #f9f9f9;
			padding: 10px 10px;
		}
	}
</style>
