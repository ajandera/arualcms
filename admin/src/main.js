import Vue from 'vue'
import VueMoment from 'vue-moment'
import VueRouter from 'vue-router'
import VueCookie from 'vue-cookies'
import VueSession from 'vue-session'
import VModal from 'vue-js-modal'
import App from './App.vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import moment from 'moment'
import Settings from '@/components/Settings'
import Posts from "@/components/Posts"
import Texts from "@/components/Texts"
import Users from "@/components/Users"
import Files from "@/components/Files"
import Sign from "@/components/Sign"
import Recovery from "@/components/Recovery"
import Languages from "@/components/Languages"

library.add(fas)
Vue.use(VueMoment)
Vue.use(VueRouter)
Vue.use(VueCookie)
Vue.use(VueSession)
Vue.use(VModal)
Vue.config.productionTip = true
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.prototype.$hostname = process.env.VUE_APP_API
Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('DD/MM/YYYY hh:mm')
  }
})

const routes = [
  { path: '/', component: App },
  { path: '/posts', component: Posts },
  { path: '/files', component: Files },
  { path: '/settings', component: Settings },
  { path: '/texts', component: Texts },
  { path: '/users', component: Users },
  { path: '/signin', component: Sign },
  { path: '/languages', component: Languages },
  { path: '/recovery/:username', component: Recovery }
]
const router = new VueRouter({
  mode: 'history',
  routes
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
