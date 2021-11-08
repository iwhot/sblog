<template>
	<div class="info-container">
		<div class="user-info">
			<div class="line"></div>
			<div class="title">
				个人资料
			</div>
			<div class="head">
				<div class="head-img">
					<img :src="headImg" width="100" height="100">
				</div>
				<div class="username">
					<div class="username-name">{{form.username}}</div>
					<div class="username-mobile">{{form.email}}</div>
				</div>
			</div>

			<div class="form-content">
				<el-form ref="form" :model="form" label-width="100px">

					<el-form-item label="用户名" prop="username">
						<el-input autocomplete="off" v-model="form.username" :disabled="true">
						</el-input>
					</el-form-item>

					<el-form-item label="昵称" prop="nickname">
						<el-input autocomplete="off" v-model="form.nickname" placeholder="请输入昵称"></el-input>
					</el-form-item>

					<el-form-item label="密码" prop="password">
						<el-input autocomplete="off" :type=" flag ? 'text' : 'password' " v-model="form.password"
							placeholder="请输入密码">
							<template slot="suffix">
								<i :class="flag ? 'el-input__icon el-icon-more' : 'el-input__icon el-icon-view' "
									style="cursor: pointer;" @click="flag = !flag"></i>
							</template>
						</el-input>
					</el-form-item>

					<el-form-item label="头像" prop="avatar">
						<el-upload class="avatar-uploader" :action="upload" :show-file-list="false" :headers="headers"
							:on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
							<img v-if="imageUrl" :src="imageUrl" class="avatar">
							<i v-else class="el-icon-plus avatar-uploader-icon"></i>
						</el-upload>
					</el-form-item>

					<el-form-item label="邮箱" prop="email">
						<el-input autocomplete="off" v-model="form.email" placeholder="请输入邮箱"></el-input>
					</el-form-item>

				</el-form>
				
				<div class="footer">
					<el-button type="success" @click="save('form')" size="mini">保存</el-button>
				</div>
				
			</div>
		</div>
	</div>
</template>

<script>
	import {
		getToken
	} from "@/utils/auth.js"
	import {userInfo} from '@/api/common/index.js'
	
	import {
		saveAdmin2
	} from '@/api/auth/index.js'
	
	export default {
		name: "GeneralConfigInfo",
		computed: {
			headers() {
				return {
					Authorization: getToken()
				}
			}
		},
		mounted() {
			this.getUserData()
		},
		data() {
			return {
				headImg: "",
				flag: false,
				upload: process.env.VUE_APP_BASE_UPLOAD,
				imageUrl: "",
				form: {
					id: this.$store.state.user.userInfo.id,
					username: this.$store.state.user.userInfo.username,
					nickname: this.$store.state.user.userInfo.nickname,
					password: "",
					avatar: this.$store.state.user.userInfo.avatar,
					email: this.$store.state.user.userInfo.email
				},
			}
		},
		methods: {
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
			save(){
				saveAdmin2(this.form).then(res => {
					if (res.data.code != 200) {
						this.$message.error(res.data.msg);
						return false;
					}
					this.getUserData()
					//更新store
					this.$store.dispatch("user/refresh_user")
					
					this.$message({
						message: res.data.msg,
						type: 'success'
					});
					
				})
			},
			getUserData(){
				userInfo({}).then( res => {
					this.form = {
						id:res.data.data.id,
						username:res.data.data.username,
						nickname:res.data.data.nickname,
						password:"",
						avatar:res.data.data.avatar,
						email:res.data.data.email,
					}
					
					this.imageUrl = res.data.data.avatar
					this.headImg = res.data.data.avatar
					
					// console.log(this.headImg)
				})
			}
		}
	}
</script>

<style scoped="scoped" lang="scss">
	.info-container {
		background-color: #f5f7f9;
		width: 100%;
		height: 100%;

		.user-info {
			// width: 530px;
			min-height: 570px;
			border: 1px solid #dee5e7;
			background-color: #ffffff;

			.line {
				width: 100%;
				height: 5px;
				background-color: #42b983;
			}

			.title {
				width: 100%;
				height: 30px;
				line-height: 30px;
				text-indent: 15px;
				font-size: 14px;
			}

			.head {
				display: flex;
				flex-direction: column;
				text-align: center;

				.head-img {
					width: 100px;
					height: 100px;
					border-radius: 50px;
					background-color: #000000;
					overflow: hidden;
					margin: 0 auto;
				}

				.username {
					padding: 3px;

					.username-name {
						font-size: 24px;
					}
				}
			}

			.form-content {
				padding: 10px 20px 20px 0;
				
				.footer{
					padding-left: 50px;
				}
			}
		}
	}
</style>
