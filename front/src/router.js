import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Addexonsunity from './views/Addexonsunity.vue'
import Updateexonsunity from './views/Updateexonsunity.vue'
import exonsunity from './views/exonsunity.vue'
import UserSignUp from './views/UserSignUp.vue'
import UserLogin from './views/UserLogin.vue'

Vue.use(Router)
    

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/add-exonsunity',
      name: 'add exonsunity',
      component: Addexonsunity
    },
    {
      path: '/exonsunity/:id/update',
      name: 'update exonsunity',
      component: Updateexonsunity
    },
    {
      path: '/exonsunity/:id',
      name: 'exonsunity',
      component: exonsunity
    },
    {
      path: '/sign-up',
      name: 'user sign up',
      component: UserSignUp
    },
    {
      path: '/login',
      name: 'login',
      component: UserLogin
    }
  ]
})
