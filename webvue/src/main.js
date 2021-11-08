import Vue from 'vue'

//引入element ui
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'

import VueAxios from 'vue-axios'
import dialogDrag from '@/utils/drag.js'
import "@/assets/common.css"
import request from './utils/request.js'
import "@/icons";
import './permission'
import Fragment from 'vue-fragment'


// require styles
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'


Vue.config.productionTip = false
//引入组件
Vue.use(ElementUI);
Vue.use(dialogDrag)
Vue.use(VueAxios, request)
Vue.use(Fragment.Plugin)

new Vue({
	router,
	store,
	render: h => h(App)
}).$mount('#app')
