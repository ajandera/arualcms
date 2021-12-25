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
import SettingsComponent from '@/components/SettingsComponent'
import PostsComponent from "@/components/PostsComponent"
import TextsComponent from "@/components/TextsComponent"
import UsersComponent from "@/components/UsersComponent"
import FilesComponent from "@/components/FilesComponent"
import SignComponent from "@/components/SignComponent"
import RecoveryComponent from "@/components/RecoveryComponent"
import LanguagesComponent from "@/components/LanguagesComponent"

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
  { path: '/admin', name: 'home', component: App },
  { path: '/admin/posts', name: 'posts', component: PostsComponent },
  { path: '/admin/files', name: 'files', component: FilesComponent },
  { path: '/admin/settings', name: 'settings', component: SettingsComponent },
  { path: '/admin/texts', name: 'texts', component: TextsComponent },
  { path: '/admin/users', name: 'users', component: UsersComponent },
  { path: '/admin/signin', name: 'sign', component: SignComponent },
  { path: '/admin/languages', name: 'languages', component: LanguagesComponent },
  { path: '/admin/recovery/:username', name: 'recovery', component: RecoveryComponent }
]
const router = new VueRouter({
  mode: 'history',
  routes
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
