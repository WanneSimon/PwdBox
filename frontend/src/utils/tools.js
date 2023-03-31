/**
 * Parse the time to string
 * @param {(Object|string|number)} time
 * @param {string} cFormat
 * @returns {string | null}
 */
export function parseTime(time, cFormat) {
    if (arguments.length === 0 || !time) {
      return null
    }
    const format = cFormat || '{y}-{m}-{d} {h}:{i}:{s}'
    let date
    if (typeof time === 'object') {
      date = time
    } else {
      if ((typeof time === 'string')) {
        if ((/^[0-9]+$/.test(time))) {
          // support "1548221490638"
          time = parseInt(time)
        } else {
          // support safari
          // https://stackoverflow.com/questions/4310953/invalid-date-in-safari
          time = time.replace(new RegExp(/-/gm), '/')
        }
      }
  
      if ((typeof time === 'number') && (time.toString().length === 10)) {
        time = time * 1000
      }
      date = new Date(time)
    }
    const formatObj = {
      y: date.getFullYear(),
      m: date.getMonth() + 1,
      d: date.getDate(),
      h: date.getHours(),
      i: date.getMinutes(),
      s: date.getSeconds(),
      a: date.getDay()
    }
    const time_str = format.replace(/{([ymdhisa])+}/g, (result, key) => {
      const value = formatObj[key]
      // Note: getDay() returns 0 on Sunday
      if (key === 'a') { return ['日', '一', '二', '三', '四', '五', '六'][value ] }
      return value.toString().padStart(2, '0')
    })
    return time_str
}

/**
 * @param {number} time
 * @param {string} option
 * @returns {string}
 */
export function formatTime(time, option) {
    if (('' + time).length === 10) {
      time = parseInt(time) * 1000
    } else {
      time = +time
    }
    const d = new Date(time)
    const now = Date.now()
  
    const diff = (now - d) / 1000
  
    if (diff < 30) {
      return '刚刚'
    } else if (diff < 3600) {
      // less 1 hour
      return Math.ceil(diff / 60) + '分钟前'
    } else if (diff < 3600 * 24) {
      return Math.ceil(diff / 3600) + '小时前'
    } else if (diff < 3600 * 24 * 2) {
      return '1天前'
    } else if (diff < 3600 * 24 * 3) {
      return '2天前'
    } else if (diff < 3600 * 24 * 4) {
      return '3天前'
    } else if (diff < 3600 * 24 * 5) {
      return '4天前'
    } else if (diff < 3600 * 24 * 6) {
      return '5天前'
    } else if (diff < 3600 * 24 * 7) {
      return '6天前'
    } else if (diff < 3600 * 24 * 14) {
      return '一周前'
    }
    if (option) {
      return parseTime(time, option)
    } else {
      return (
        d.getMonth() +
        1 +
        '月' +
        d.getDate() +
        '日' +
        d.getHours() +
        '时' +
        d.getMinutes() +
        '分'
      )
    }
}


export function cutString(str, size, suffix) {
  if(!str || str.length <= size) {
    return str
  }

  suffix = suffix ? suffix : ''
  return str.substring(0,size) + suffix
}

// 数组浅拷贝
export function shallowCopyArry (arr, start, end) {
  let a = innerShallowCopyArry(arr, start, end)
  return a
}
function innerShallowCopyArry (arr, start, end) {
  // console.log(arr, start, end)
  if (end && start >= end) {
    throw new Error("复制数组，起始下标不能大于等于终止下标")
  }

  let re = []
  end = arr.length > end ? end : arr.length
  // console.log(arr, start, end)
  for( let i=start; i<end; i++) {
    re.push(arr[i])
  }
  return re
}


/**
 * 允许绝对布局的dom元素被拖拽。拖动太快了体验不好
 * https://article.itxueyuan.com/69oARg
 * @param {*} els 元素
 * @param {*} parentEl 父元素，用来限制拖动范围。不设置则不限制
 */
export function drageElements(els, parentEl) {

  // let drag = document.querySelector("#drag");//获取操作元素
  for(let i=0; i<els.length; i++) {
    let drag = els[i]
    drag.onmousedown = function (e) {//鼠标按下触发
        var disx = e.pageX - drag.offsetLeft;//获取鼠标相对元素距离
        var disy = e.pageY - drag.offsetTop;
        // console.log(e.pageX);
        // console.log(drag.offsetLeft);
        document.onmousemove = function (e) {//鼠标移动触发事件，元素移到对应为位置
          let x = e.pageX - disx;
          let y = e.pageY - disy;
          if(!parentEl) {
            drag.style.left = x + 'px';
            drag.style.top = y + 'px';
            return
          }

          if(x >0 && x<parentEl.clientWidth-drag.clientWidth) {
            drag.style.left = x + 'px';
          }
          if( y>0 && y<parentEl.clientHeight-drag.clientHeight) {
            drag.style.top = y + 'px';
          }
        }
        document.onmouseup = function(){//鼠标抬起，清除绑定的事件，元素放置在对应的位置
            document.onmousemove = null;
            document.onmousedown = null;
        };
        e.preventDefault();//阻止浏览器的默认事件
    };
  }

}

