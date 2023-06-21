// const componentModules = import.meta.globEager('/src/views/**/*.vue') // 立即引入
const componentModules = import.meta.glob('/src/views/**/*.vue') // 懒加载

import { RouteRecordRaw } from 'vue-router'
// import Hello2 from '/src/views/manage/hello.vue'
import EmptyView from '/src/views/components/empty-view.vue'

/**
 * 路由菜单的说明：
 * co_menu: { label: '后台管理', disabled: false, icon: '', alwaysShow:true, permission: 'permission:*' },
 * 路由中带有 co_menu 属性的被认定为菜单，
 * label: 必选，名字 
 * disabled: 可选，是否禁用，展示出来无法的点击。不填或默认为 false
 * icon：可选，渲染的图标，图标来源图标库`vicons` - @vicons/ionicons5 (可以选择 vicons 中的其他)
 * alwaysShow：可选，当级联菜单只有一个菜单时，默认会使用子菜单代替父菜单，设置为 false 可以保持级联结构。默认 false
 * permission：可选，菜单需要的权限。不填写则不需要权限。
 */

// 后台管理根路径
const MANAGE_ROOT = import.meta.env.VITE_MANAGE_ROOT
// console.log("MANAGE_ROOT", MANAGE_ROOT)

function _importView(_path:string){
  return componentModules[`/src/views/${_path}`]
}
function _importFrontView(_path:string){
  return componentModules[`/src/views/front/${_path}`]
}
function _import(_path:string){
  return componentModules[`/src/${_path}`]
}

const routes :Array<RouteRecordRaw> = [
    { path: '/', redirect: '/home3' },
    { path: '/index', redirect: '/home3' },
    { path: '/home', redirect: '/home3' },
    { path: '/home3', component: _importFrontView("home/home-v3.vue") },
    { path: '/tools', component: _importFrontView("tools/tools.vue") },
    { path: '/pwdbox', component: _importFrontView("pwdbox/pwdbox.vue") },
    { 
      path: '/tool',  component: EmptyView,
      children: [
        // flv拉流 (flv.js)
        { name: 'live_', path: 'live',
          components: {
            default: _importFrontView('tools/live/live.vue')
          }
        }
      ]
    },
    // ========
    // { path: "/404", name: "notFound", component: _importView('home.vue') },
    { path: "/404", name: "notFound", redirect: '/home' },
    {
      path: "/:pathMatch(.*)", // 此处需特别注意置于最底部
      redirect: "/404",
    }
  ]

export default routes