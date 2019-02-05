// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import chance from 'chance'
// import jquery from 'jquery'
Vue.use(chance)
// Vue.use(jquery)
// import VueRouter from 'vue-router'
// import router from './router/router.js'
// Vue.use(VueRouter)
Vue.config.productionTip = false

/* eslint-disable no-new */
let vm = new Vue({
  el: '#app',
  components: { App },
  template: '<App/>'
})
