<template>
	<div class="quill-editor">
		<slot name="toolbar"></slot>
		<div ref="editor"></div>
	</div>
</template>

<script>
	// require sources
	import _Quill from 'quill'
	import defaultOptions from './options'
	import {
		getToken
	} from "@/utils/auth";

	const Quill = window.Quill || _Quill

	// pollfill
	if (typeof Object.assign != 'function') {
		Object.defineProperty(Object, 'assign', {
			value(target, varArgs) {
				if (target == null) {
					throw new TypeError('Cannot convert undefined or null to object')
				}
				const to = Object(target)
				for (let index = 1; index < arguments.length; index++) {
					const nextSource = arguments[index]
					if (nextSource != null) {
						for (const nextKey in nextSource) {
							if (Object.prototype.hasOwnProperty.call(nextSource, nextKey)) {
								to[nextKey] = nextSource[nextKey]
							}
						}
					}
				}
				return to
			},
			writable: true,
			configurable: true
		})
	}

	// export
	export default {
		name: 'Editor',
		data() {
			return {
				_options: {},
				_content: '',
				defaultOptions,
				container: null,
			}
		},
		props: {
			content: String,
			value: String,
			disabled: {
				type: Boolean,
				default: false
			},
			options: {
				type: Object,
				required: false,
				default: () => ({})
			},
			globalOptions: {
				type: Object,
				required: false,
				default: () => ({})
			}
		},
		mounted() {
			this.initialize()
		},
		beforeDestroy() {
			this.quill = null
			delete this.quill
		},
		methods: {
			// Init Quill instance
			initialize() {
				if (this.$el) {

					const Parchment = Quill.import("parchment");
					class lineHeightAttributor extends Parchment.Attributor.Style {}
					const lineHeightStyle = new lineHeightAttributor("lineheight", "line-height", {
						scope: Parchment.Scope.INLINE,
						whitelist: ["1", "1.5", "1.75", "2", "3", "4", "5"],
					});
					Quill.register({
						"formats/lineheight": lineHeightStyle
					}, true);

					let fontSizeStyle = Quill.import('attributors/style/size')
					fontSizeStyle.whitelist = ['10px', '12px', '14px', '16px' ,'18px', '20px', '22px', '24px', '26px', '32px', '48px',false]
					Quill.register(fontSizeStyle, true) //重写字号选择列表

					// Options
					this._options = Object.assign({}, this.defaultOptions, this.globalOptions, this.options)

					// Instance
					this.quill = new Quill(this.$refs.editor, this._options)
					this.quill.enable(false)

					//添加handle
					var toolbar = this.quill.getModule('toolbar');
					this.container = this.$refs.editor.previousElementSibling
					toolbar.addHandler('image', this.image);
					toolbar.addHandler('lineheight', this.lineheight);

					// Set editor content
					if (this.value || this.content) {
						this.quill.pasteHTML(this.value || this.content)
					}

					// Disabled editor
					if (!this.disabled) {
						this.quill.enable(true)
					}

					// Mark model as touched if editor lost focus
					this.quill.on('selection-change', range => {
						if (!range) {
							this.$emit('blur', this.quill)
						} else {
							this.$emit('focus', this.quill)
						}
					})

					// Update model if text changes
					this.quill.on('text-change', (delta, oldDelta, source) => {
						let html = this.$refs.editor.children[0].innerHTML
						const quill = this.quill
						const text = this.quill.getText()
						if (html === '<p><br></p>') html = ''
						this._content = html
						this.$emit('input', this._content)
						this.$emit('change', {
							html,
							text,
							quill
						})
					})

					// Emit ready event
					this.$emit('ready', this.quill)
				}
			},
			//图片重写
			image() {
				console.log(this)
				/*富文本编辑图片上传配置*/
				const uploadConfig = {
					action: process.env.VUE_APP_BASE_UPLOAD + "-editor", // 必填参数 图片上传地址
					methods: 'POST', // 必填参数 图片上传方式
					token: '', // 可选参数 如果需要token验证，假设你的token有存放在sessionStorage
					name: 'file', // 必填参数 文件的参数名
					size: 500, // 可选参数   图片大小，单位为Kb, 1M = 1024Kb
					accept: 'image/png, image/gif, image/jpeg, image/bmp, image/x-icon' // 可选 可上传的图片格式
				};

				var self = this;
				var Authorization = getToken();
				var fileInput = this.container.querySelector('input.ql-image[type=file]');
				if (fileInput === null) {
					fileInput = document.createElement('input');
					fileInput.setAttribute('type', 'file');
					// 设置图片参数名
					if (uploadConfig.name) {
						fileInput.setAttribute('name', uploadConfig.name);
					}
					// 可设置上传图片的格式
					fileInput.setAttribute('accept', uploadConfig.accept);
					fileInput.classList.add('ql-image');
					// 监听选择文件
					fileInput.addEventListener('change', function() {
						// 创建formData
						var formData = new FormData();
						formData.append(uploadConfig.name, fileInput.files[0]);
						formData.append('object', 'product');
						// 如果需要token且存在token
						if (uploadConfig.token) {
							formData.append('token', uploadConfig.token)
						}
						// 图片上传
						var xhr = new XMLHttpRequest();
						xhr.open(uploadConfig.methods, uploadConfig.action, true);
						xhr.setRequestHeader("Authorization", Authorization)
						// 上传数据成功，会触发
						xhr.onload = function(e) {
							if (xhr.status === 200) {
								var res = JSON.parse(xhr.responseText);
								let length = self.quill.getSelection(true).index;
								//这里很重要，你图片上传成功后，img的src需要在这里添加，res.path就是你服务器返回的图片链接。            
								self.quill.insertEmbed(length, 'image', res.path);
								self.quill.setSelection(length + 1)
							}
							fileInput.value = ''
						};
						// 开始上传数据
						xhr.upload.onloadstart = function(e) {
							fileInput.value = ''
						};
						// 当发生网络异常的时候会触发，如果上传数据的过程还未结束
						xhr.upload.onerror = function(e) {};
						// 上传数据完成（成功或者失败）时会触发
						xhr.upload.onloadend = function(e) {
							// console.log('上传结束')
						};
						xhr.send(formData)
					});
					this.container.appendChild(fileInput);
				}
				fileInput.click();
			},
			lineheight(value) {
				if (value) {
					this.quill.format("lineheight", value);
				}
			}
		},
		watch: {
			// Watch content change
			content(newVal, oldVal) {
				if (this.quill) {
					if (newVal && newVal !== this._content) {
						this._content = newVal
						this.quill.pasteHTML(newVal)
					} else if (!newVal) {
						this.quill.setText('')
					}
				}
			},
			// Watch content change
			value(newVal, oldVal) {
				if (this.quill) {
					if (newVal && newVal !== this._content) {
						this._content = newVal
						this.quill.pasteHTML(newVal)
					} else if (!newVal) {
						this.quill.setText('')
					}
				}
			},
			// Watch disabled change
			disabled(newVal, oldVal) {
				if (this.quill) {
					this.quill.enable(!newVal)
				}
			}
		}
	}
</script>
