import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/loginView.vue'
import RegisterView from '@/views/registerView.vue'
import Dashboardview from '@/views/DasboardView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboardview ,
      meta: { requiresAuth: true }
    }
  ]
})


router.beforeEach((to, from, next) => {
  const loggedIn = !!localStorage.getItem('token');
  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    next('/');
  } else {
    next();
  }
})

export default router
