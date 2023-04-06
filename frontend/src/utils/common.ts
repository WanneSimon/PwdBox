/**
 * 加载图片
 * @param {*} url 
 */
export function loadImage(url:string) {
  return innerLoadImage(url)
}

function innerLoadImage(url:string) {
  return new Promise<string>(function(resolve, reject){
    let img = new Image()
    
    img.onload = function() {
      document.body.removeChild(img)
      resolve(url)
    }
    img.onerror = function(error) {
      document.body.appendChild(img)
      reject(error)
    }

    img.style.display = "none"
    img.src = url
    document.body.appendChild(img)
  }).catch(err => {
    return Promise.reject(err)
  })
}

/** 加载背景
 * @param url - 背景url
 * @param oldBg - 预设背景，可选
 * @param mainBg - 新的背景
 * @param callback - 成功加载后的其他操作,可选

// export function loadBackground(url, mainBg, oldBg) {
export function loadBackground(url:string, mainBg:string) {
  innerLoadImage(url)
    .then(img => {
      // const oldBg = this.$refs.originBg;
      // console.log("loaded-img", img, oldBg.classList)
      // if(oldBg) {
      //   oldBg.classList.add('fadeOut')
      // }

      // const mainBg = this.$refs.mainBg
      // mainBg.style.background = `#00ff00 url('${img}') no-repeat fixed center`
      mainBg.style.background = `url(${img})`
      mainBg.style.backgroundSize = '100% 100%'
      mainBg.style.backgroundRepeat = 'no-repeat'
      // mainBg.src = `${img}`
      mainBg.classList.add('fadeIn')
    }, (error) => {
      console.error("error-img", error)
    }).catch(err => {
      return Promise.reject(err)
    })
}
 */ 