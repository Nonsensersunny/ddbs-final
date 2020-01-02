import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import VueCookie from 'vue-cookie'

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.use(VueCookie)

router.beforeEach((to, from, next) => {
  if (!store.state.isLogin && to.path !== '/') {
    ElementUI.Message.error("登录后操作");
    next('/');
    return
  }
  next();
})

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
