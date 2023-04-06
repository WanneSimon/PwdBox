var _rootPath = import.meta.env.VITE_FRONT_BASE ? import.meta.env.VITE_FRONT_BASE : '/'
const ROOT_PATH = _rootPath.endsWith('/') ? _rootPath : _rootPath + '/'
// const ROOT_PATH = _rootPath

/**
 * 
 * @param {*} cname  不允许符号出现
 * @param {*} cvalue 不允许符号出现
 * @param {*} exdays 
 * @param {*} path 注意path要以 '/' 结尾，同时 path 要存在才能设置成功！
 */
function setCookie(cname:string,cvalue:any,exSeconds:number,path:string){
  if(!cname) {
    throw new Error('Invalid cookie')
  }

  var d = new Date();
  d.setTime(d.getTime()+(exSeconds*1000));
  // var expires = "expires="+d.toGMTString();
  var expires = "expires="+d.toUTCString();
  cvalue = cvalue ? cvalue.trim() : cvalue
  var _path = path ? `path=${path}` : ''
  // document.cookie = cname.trim() + "=" + cvalue + "; " + expires + "; "+_path;
  const newCookie = `${cname.trim()}=${cvalue}; ${expires}; ${_path}`;
  // const newCookie = `${cname.trim()}=${cvalue}; ${expires}`;
  // console.log("set-cookie", newCookie)
  document.cookie = newCookie
}

function setRootCookie(cname:string, cvalue:any, exSeconds:number){
  setCookie(cname, cvalue, exSeconds, ROOT_PATH)
}

function removeCookie(cname:string, path:string) {
  setCookie(cname, '', 0, path)
}

function removeRootCookie(cname:string) {
  setCookie(cname, '', 0, ROOT_PATH)
}

function getCookie(cname:string){
  var name = cname + "=";
  var ca = document.cookie.split(';');
  for(var i=0; i<ca.length; i++) {
      var c = ca[i].trim();
      if (c.indexOf(name)==0) {
        return c.substring(name.length,c.length);
      }
  }
  return "";
}

function allCookies(){
  // var ca = document.cookie.split(';');
  // for(var i=0; i<ca.length; i++) {
  //     var c = ca[i].trim();
  //     if (c.indexOf(name)==0) { return c.substring(name.length,c.length); }
  // }
  // console.log("cookies", document.cookie)
  return document.cookie.split(';').map(e => e.trim());
}

export default {
  setCookie, getCookie, allCookies, removeCookie, setRootCookie, removeRootCookie,
}