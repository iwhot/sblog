<template>
	<div class="login-container">
		<div class="login-item">
			<img src="../../assets/login.png" class="login-image">
			<div class="login-info">
				<h3>{{title}}</h3>
				<div class="login-form">
					<el-form ref="loginForm" :model="loginForm" :rules="rules" label-width="0">
						<el-form-item prop="username">
							<el-input v-model="loginForm.username" prefix-icon="el-icon-user" placeholder="请输入账号">
							</el-input>
						</el-form-item>
						<el-form-item prop="password">
							<el-input type="password" v-model="loginForm.password" prefix-icon="el-icon-key"
								placeholder="请输入密码">
							</el-input>
						</el-form-item>
						<el-form-item prop="captcha">
							<el-col :span="14">
								<el-input v-model="loginForm.captcha" prefix-icon="el-icon-present"
									placeholder="请输入验证码">
								</el-input>
							</el-col>
							<el-col :span="4">
								<img :src="captchaImg" class="capcha" @click="changeCaptcha">
							</el-col>
						</el-form-item>
						<el-button type="primary" class="submit-btn" @click="loginSubmit('loginForm')">登录</el-button>
					</el-form>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import {
		login,
		captcha
	} from '@/api/login/login.js'


	//创建随机数
	//var randomUuid = uuid()

	export default {
		name: "Login",
		data() {
			return {
				title: "后台内容管理系统",
				captchaImg: "",
				loginForm: {
					username: "",
					password: "",
					captcha: "",
					sess: ""
				},
				rules: {
					username: [{
						required: true,
						message: '请输入登陆账号',
						trigger: 'blur'
					}, ],
					password: [{
						required: true,
						message: '请输入密码',
						trigger: 'blur'
					}, ],
					captcha: [{
						required: true,
						message: '请输入验证码',
						trigger: 'blur'
					}, ]
				}
			}
		},
		created() {
			this.createCaptcha()
			this.keyupSubmit()
		},
		methods: {
			changeCaptcha() {
				this.createCaptcha()
			},
			//创建验证码
			createCaptcha() {
				captcha({}).then(res => {
					this.loginForm.sess = res.data.data.sess
					this.captchaImg = process.env.VUE_APP_BASE_API + res.data.data.url
				});
			},
			loginSubmit(formName) {
				this.$refs[formName].validate((valid) => {
					if (!valid) {
						return false;
					}

					const loading = this.$loading({
						lock: true,
						text: 'Loading',
						spinner: 'el-icon-loading',
						background: 'rgba(0, 0, 0, 0.7)'
					});
					//开始登陆
					login({
						username: this.loginForm.username,
						password: this.loginForm.password,
						captcha: this.loginForm.captcha,
						sess: this.loginForm.sess
					}).then(res => {
						if (res.data.code != 200) {
							this.$message.error(res.data.msg);
							window.localStorage.clear()
							window.sessionStorage.clear()
							setTimeout(() => {
								loading.close();
							}, 1000);
							return false;
						}

						this.$message({
							message: res.data.msg,
							type: 'success'
						});

						//存储token
						this.$store.dispatch('user/set_user_info', res.data.data.user)
						this.$store.dispatch('user/set_token', res.data.data.token)

						//缓存菜单

						setTimeout(() => {
							loading.close();
							//跳转页面
							window.location.reload()
						}, 1000);
						return true;
					})
				});
			},
			//监听回车键
			keyupSubmit(){
			  document.onkeydown=e=>{
			    let _key=window.event.keyCode;
			    if(_key===13){
			      this.loginSubmit('loginForm')
			    }
			  }
			}
		}
	}
</script>

<style scoped="scoped" lang="scss">
	.login-container {
		width: 100%;
		height: 100%;
		background-color: #f4f4f4;

		.login-item {
			width: 910px;
			height: 540px;
			background-color: #ffffff;
			border-radius: 10px;
			margin: 0 auto;
			box-shadow: 0 30px 50px 0 rgba(124, 172, 255, 0.12);
			display: block;
			position: absolute;
			left: 0;
			right: 0;
			top: 50%;
			transform: translateY(-50%);
			color: #1a263b;

			.login-image {
				width: 400px;
				display: inline-block;
			}

			.login-info {
				display: inline-block;
				padding: 50px 0 50px 50px;
				box-sizing: border-box;
				vertical-align: top;

				h3 {
					font-size: 30px;
					font-weight: 700;
				}

				.login-form {
					margin-top: 24px;

					.submit-btn {
						width: 410px;
						height: 50px;
						border-radius: 4px;
						font-size: 16px;
						color: #FFFFFF;
						margin-top: 5px;
					}

					.capcha {
						width: 156px;
						height: 40px;
						cursor: pointer;
						border-radius: 4px;
						margin-left: 10px;
					}
				}
			}
		}
	}
</style>
