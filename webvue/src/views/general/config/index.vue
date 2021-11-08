<template>
	<div class="config-container">
		<el-form ref="form" :label-position="labelPosition" label-width="150px">
			<el-tabs type="border-card">
				<el-tab-pane label="基础配置">

					<el-form-item label="站点名称">
						<el-input size="small" v-model="form.basic.site_name" placeholder="请输入站点名称"></el-input>
					</el-form-item>

					<el-form-item label="站点地址">
						<el-input size="small" v-model="form.basic.site_url" placeholder="请输入站点地址"></el-input>
					</el-form-item>

					<el-form-item label="　备案号">
						<el-input size="small" v-model="form.basic.bei" placeholder="请输入备案号"></el-input>
					</el-form-item>

					<el-form-item label="微信">
						<el-upload class="avatar-uploader" :action="upload" :show-file-list="false" :headers="headers"
							:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
							<img v-if="imageUrl" :src="imageUrl" class="avatar">
							<i v-else class="el-icon-plus avatar-uploader-icon"></i>
						</el-upload>
					</el-form-item>

					<el-form-item label="QQ">
						<el-input size="small" v-model="form.basic.qq" placeholder="请输入QQ"></el-input>
					</el-form-item>

					<el-form-item label="邮箱">
						<el-input size="small" v-model="form.basic.email" placeholder="请输入邮箱"></el-input>
					</el-form-item>

					<el-form-item label="邮我">
						<el-input size="small" v-model="form.basic.email_me" type="textarea" 
							placeholder="请输入邮我"></el-input>
					</el-form-item>

					<el-form-item label="网名">
						<el-input size="small" v-model="form.basic.username" placeholder="请输入网名"></el-input>
					</el-form-item>

					<el-form-item label="职业">
						<el-input size="small" v-model="form.basic.occupation" placeholder="请输入职业"></el-input>
					</el-form-item>

					<el-form-item label="现居">
						<el-input size="small" v-model="form.basic.address" placeholder="请输入现居"></el-input>
					</el-form-item>

					<el-form-item label="页面底部信息">
						<el-input size="small" v-model="form.basic.footer" type="textarea" rows="5"></el-input>
					</el-form-item>
					
					<el-form-item label="转载">
						<el-input size="small" v-model="form.basic.reprint" type="textarea" rows="5"></el-input>
					</el-form-item>

				</el-tab-pane>
				<el-tab-pane label="SEO配置">
					<el-form-item label="SEO标题">
						<el-input size="small" v-model="form.seo.seo_title" placeholder="请输入SEO标题"></el-input>
					</el-form-item>

					<el-form-item label="SEO关键字">
						<el-input size="small" v-model="form.seo.seo_key" type="textarea" rows="5"
							placeholder="请输入SEO关键字"></el-input>
					</el-form-item>

					<el-form-item label="SEO描述">
						<el-input size="small" v-model="form.seo.seo_desc" type="textarea" rows="5"
							placeholder="请输入SEO描述"></el-input>
					</el-form-item>
				</el-tab-pane>
				<el-form-item style="margin-bottom: 0;">
					<el-button type="primary" size="mini" @click="saveData()">保存</el-button>
				</el-form-item>
			</el-tabs>


		</el-form>
	</div>
</template>

<script>

	import {
		indexConfig,
		saveConfig
	} from '@/api/general/index.js'

	import {
		getToken
	} from "@/utils/auth.js"

	// import OSS from 'ali-oss'
	import {
		randomString
	} from '@/utils/random.js'

	export default {
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
				percentage2: 0,
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: '',
				form: {
					basic: {
						site_name: '',
						bei: '',
						footer: '',
						wechat: '',
						qq: '',
						email: '',
						username: "",
						address: "",
						occupation: "",
						site_url: "",
						email_me: "",
						reprint:""
					},
					seo: {
						seo_title: "",
						seo_key: "",
						seo_desc: ""
					}
				}
			};
		},
		mounted() {
			this.getData()
		},
		methods: {
			getData() {
				indexConfig({}).then(res => {
					this.form = res.data.data
					// console.log(res.data.data.basic.wechat)
					if(res.data.data.basic.wechat != ""){
						this.imageUrl = res.data.data.basic.wechat
					}
				})
			},
			saveData() {
				saveConfig(this.form).then(res => {
					this.$message({
						message: "保存成功",
						type: 'success'
					});
				})
			},
			handleAvatarSuccess(res, file) {
				this.form.basic.wechat = res.data
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
			// async UploadOSS(option) {
			// 	try {
			// 		//连接oss
			// 		const client = new OSS({
			// 			region: this.$store.state.oss.OSSConfig.region,
			// 			accessKeyId: this.$store.state.oss.OSSConfig.accessKeyId,
			// 			accessKeySecret: this.$store.state.oss.OSSConfig.accessKeySecret,
			// 			bucket: this.$store.state.oss.OSSConfig.bucket,
			// 		});

			// 		const progress = (p, _checkpoint) => {
			// 			var process = p * 100
			// 			this.percentage2 = process.toFixed(2)
			// 		};

			// 		var file = option.file; // 拿到 file
			// 		var point = file.name.lastIndexOf('.')
			// 		var suffix = file.name.substr(point)
			// 		let date = new Date()
			// 		var dir = '/zhuixi_kecheng/' + date.getFullYear() + "" + (date.getMonth() + 1) + "/" + date
			// 			.getTime() + "_" + randomString(
			// 				20) + suffix

			// 		// object-name可以定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至Bucket根目录或Bucket下的指定目录。
			// 		const result = await client.multipartUpload(dir, option.file, {
			// 			progress
			// 		});

			// 		this.$message({
			// 			message: "上传成功",
			// 			type: 'success'
			// 		});
			// 	} catch (error) {
			// 		console.log(error)
			// 		this.$message.error("上传失败");
			// 	}


			// }

		}
	};
</script>


<style lang="scss" scoped="scoped">
	.config-container {
		width: 100%;
		height: 100%;
		background-color: #f5f7f9;

		.gray {
			background-color: #f9f9f9;
			padding: 10px 10px;
		}
	}
</style>
