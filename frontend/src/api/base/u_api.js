// 保留一个使用样例，该文件并未使用到
import request from '../request'

const ctxBase = "/base/u"

function login(data) {
    return request.post(ctxBase + "/login", data)
}

// 管理员登录
function loginEnc(data) {
  return request.post(ctxBase + "/loginEnc", data)
}

// 退出。不区分登录用户的类型
function logout(data) {
    return request.post(ctxBase + "/logout", data)
}

// 获取管理员 token 有关的信息必须是同步的
function currentAdminInfo(data) {
  return request.post(ctxBase + "/adminInfo", data)
}

// 获取普通用户 token 有关的信息必须是同步的
function currentUserInfo(data) {
  return request.post(ctxBase + "/userInfo", data)
}

// 获取登录用户的权限。不区分登录用户的类型
function permissions(data) {
  return request.post(ctxBase + "/permissions", data)
}

// 普通用户登录
function signupEnc(data) {
  return request.post(ctxBase + "/signupEnc", data)
}

export default {
  login, loginEnc, logout, currentAdminInfo, currentUserInfo, 
  permissions, signupEnc,
}
