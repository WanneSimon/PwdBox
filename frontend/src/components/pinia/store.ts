import { defineStore, storeToRefs } from "pinia"
import cookieApi from '@/api/cookie_api'

// const store = globalStore()
const cookieName = import.meta.env.VITE_COOKIE_AUTH

// https://pinia.vuejs.org/introduction.html#a-more-realistic-example

type StateInfo = {
  count: number,
  token: string | null |undefined,
  userinfo: any | null |undefined,
  permissions: Array<any> | null |undefined
}

const globalStore = defineStore('store', {
  state() {
      const data : StateInfo = {
        count: 0, // for test
        token: null,
        userinfo: null,
        permissions: null,  //  [ {role: '', permissions: [''] } ]
      }
      return data
  },

  getters: {
    getCounter: (state) => { // for test
      console.log('this.counter', state)
      return state.count
    },

    getUserInfo: (state) => {
      return state.userinfo
    },
    getPermissions: (state) => {
      return state.permissions
    }
    //=====
  },

  actions: {
    addCounter (delta:number) { // for test
      this.count += delta
    },
    setToken(newToken:string, isSaveCookie=true) {
      this.token = newToken
      if(isSaveCookie && cookieName) {
        cookieApi.setRootCookie(cookieName, newToken, 7*24*60*60)
      }
    },
    clearToken() {
      this.token = null
      if(cookieName) {
        // cookieApi.removeCookie(cookieName)
        cookieApi.removeRootCookie(cookieName)
      }
      this.permissions = null
    },
    setUserinfo(userInfo:any) {
      this.userinfo = userInfo
    },
    setPermissions(list:Array<any>) {
      this.permissions = list
    },
    hasPermission(permission:any) {
      if(!this.permissions) {
        return false
      }

      // [ {role: '', permissions: [''] } ]
      for(let i=0; i<this.permissions.length; i++) {
        let item = this.permissions[i]
        if(item?.permissions) {
          let list = item?.permissions
          if(list.indexOf(permission) >= 0) {
            return true
          }
        }
      }
      return false
    },
    
    //=====
  }

})


// const store = null // 不要提前调用此函数创建对象
// console.log('store', store)
// storeToRefs(store) // 在需要使用的地方转换内部数据为响应式数据
export default globalStore

// 下面的方式可以直接返回一个 store 和 storeRef 对象, 只需要在 setup 函数中调用一次
// 返回一个匿名函数，不立即调用，需要在setup中使用
// export default () => { 
//   const store = globalStore()
//   return {
//     store, 
//     storeRef: storeToRefs(store)
//   }
// }



