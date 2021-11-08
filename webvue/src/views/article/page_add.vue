<template>
	<div class="tab-container">
		<vuescroll>
			<el-form ref="form" :label-position="labelPosition" label-width="150px">
				<el-tabs type="border-card">
					<el-tab-pane label="基础设置">
						<el-form-item label="标题">
							<el-input size="small" v-model="form.title" placeholder="请输入标题"></el-input>
						</el-form-item>
						
						<el-form-item label="内容">
							<editor v-model="form.content" style="height: 375px; margin-bottom: 60px;">
							</editor>
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
	import editor from '@/components/editor/index.vue'
	import {
		CategoryCascader,
		ArticleSave
	} from "@/api/article/index.js"
	import {
		getToken
	} from "@/utils/auth.js"

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
					title: "",
					content: "",
					status: 1,
					sort: 1000,
					state: 2,
					attr: 1,
					seo_title: "",
					seo_desc: "",
					seo_key: "",
					path:"",
				},
				options: [],
				ap: [],
			};
		},
		methods: {
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
				this.form.attr = 1
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
