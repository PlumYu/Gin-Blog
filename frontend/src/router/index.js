import Vue from 'vue';
import store from '@/store';
import VueRouter from 'vue-router';
import HomeView from '../views/HomeView.vue';
import userRoutes from './module/user';


Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
  },
  ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 判断是否需要登录
    // 判断用户是否登录
    if (store.state.userModule.token) {
      // 判断 token 的有效性 比如有没有过期 需要后台发放 token 的时候 带上 token 的有效期
      //  如果 token 无效 需要 请求 token
      next()
    } else {
      // 跳转登录
      router.push({ name: 'login' });
    }
  } else {
    next()
  }
});

export default router;
