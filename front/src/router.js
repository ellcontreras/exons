import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import AddCommunity from './views/AddCommunity.vue'
import UpdateCommunity from './views/UpdateCommunity.vue'
import Community from './views/Community.vue'
import UserSignUp from './views/UserSignUp.vue'
import UserLogin from './views/UserLogin.vue'
import Profile from './views/Profile.vue'

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
      path: '/add-community',
      name: 'add Community',
      component: AddCommunity
    },
    {
      path: '/community/:id/update',
      name: 'update Community',
      component: UpdateCommunity
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
    },
    {
      path: '/user/:id',
      name: 'user profile',
      component: Profile
    }
  ]
})
