
import {createRouter, createWebHistory, createWebHashHistory } from 'vue-router'

import routes from './routerOptions'

const basePath = import.meta.env.VITE_FRONT_BASE

const router = createRouter({
  //history: createWebHashHistory(),
  //（不能设置根路径，否则 wails 会找不到路径）
  history: createWebHistory(basePath),
  routes,
})

router.beforeEach( async (to, from, next) => {
  //await checkRemember(to)
  // console.log("route-change")
  // TODO 跳转到需要权限的页面，预先检查权限
  next()
})
router.afterEach(() => {
})

export default router