import axios from 'axios'
import globalStore from '/src/components/pinia/store.js'

const SUCCESS_CODE = "00"
// const storeProvider = () => {
//   // return storeToRefs(globalStore())
//   return globalStore()
// }
// const store = () => {
//   return globalStore()
// }

// create an axios instance
const service = axios.create({
  baseURL: import.meta.env.VITE_API_HOST,
  timeout: 30000 // request timeout
})

service.interceptors.request.use(
  config => {
    // do something before request is sent
    // 后端使用了 cookie，没有使用这种方式，该功能仅作保留
    const HEADER_AUTH_KEY = import.meta.env.VITE_HEADER_AUTH
    const store = globalStore()
    // console.log("request", store.token, store)
    // console.log("request", HEADER_AUTH_KEY, store.token)
    if (HEADER_AUTH_KEY && store.token) {
     // let each request carry token
     // ['X-Token'] is a custom headers key
     // please modify it according to the actual situation
      config.headers[HEADER_AUTH_KEY] = store.token
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * 错误error请求拦截，否则放行
   */
  response => {
    const resData = response.data
    // 如果resData代表错误信息
    if (resData.code != SUCCESS_CODE) {
      let errorInfo = resData.info || 'Error'
      return Promise.reject(new Error(errorInfo)) // resData.info
    } else {
      return resData
    }
  },
  error => {
    //console.error(error.response) // for debug
    let errorMsg = error.response?.data?.message || error.statusText || error.message 

    let httpCode = error.response?.status
    // if(httpCode == 403) {
    //   errorMsg = 
    // }
    switch(httpCode) {
      case 400:  errorMsg = '参数错误'; break;
      case 401:  errorMsg = '未授权'; break;
      case 403:  errorMsg = '暂无权限'; break;
      case 404:  errorMsg = '请求资源不存在'; break;
      case 405:  errorMsg = '请求类型错误'; break;
      case 500:  errorMsg = '服务器内部错误'; break;
    }
    return Promise.reject(new Error(errorMsg))
  }
)

export default service
