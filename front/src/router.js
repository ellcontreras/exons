import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import AddCommunity from './views/AddCommunity.vue'
import Community from './views/Community.vue'

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
      name: 'add community',
      component: AddCommunity
    },
    {
      path: '/community/:id',
      name: 'community',
      component: Community
    }
  ]
})
