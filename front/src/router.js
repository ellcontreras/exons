import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Addexonsunity from './views/Addexonsunity.vue'
import Updateexonsunity from './views/Updateexonsunity.vue'
import Community from './views/Community.vue'
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
      path: '/add-Community',
      name: 'add Community',
      component: Addexonsunity
    },
    {
      path: '/Community/:id/update',
      name: 'update Community',
      component: Updateexonsunity
    },
    {
      path: '/Community/:id',
      name: 'Community',
      component: Community
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
