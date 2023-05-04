// import { Greet, FileOp, Config } from "@/../wailsjs/index"
import { Greet, FileOp, ConfigOp } from "@/../wailsjs/index"
// import byteString from '/src/utils/encrypt/byte-string.js'
import EmojiConfig from "./emoji-config/emoji-config.vue"

export default {
  name: 'Emoji',
  components: { EmojiConfig },
  data () {
    return {
      name: "" , // 搜索输入名称
      showConfig: false, 

      loading: false, // 是否正在加载
      paths: [], // 所有文件夹路径
      files: [], // {isDir, name, path, size}
    }
  },
  async mounted() { 
    // this.listDir("D:\\Picture")
    await this.getEmojiPath()
  },
  methods: {
    async testList() {
      let paths = await this.getEmojiPath()
      console.log("emojis", paths)

      // for(let i=0; i<paths.length; i++) {
      //   await this.listDir(paths[i])
      // }
    },
    async listDir(path) {
      let fsList = await FileOp.ListImage(path).then(res => res)
      // console.log("fsList", fsList)
      return fsList
    },

    // 获取文件夹路径
    async getEmojiPath() {
      // let conf =  await Config().then(res => res)
      console.log("ConfigOp", ConfigOp)
      let conf =  await ConfigOp.Get().then(res => res)
      console.log("ConfigOp-after", conf)

      this.paths = conf.Emojis
      console.log("this.paths", this.paths)
      return this.paths
    },
    refreshConfig() {
      // 异步刷新即可，不删除当前结果
      console.log("refresh!")
      this.getEmojiPath()
    },

    // 根据文件名搜索
    async searchByName(name) {
      if(name == undefined || name == null) {
        return
      }

      let arr = []
      this.loading = true
      for(let i=0; i< this.paths.length; i++) {
        let p = this.paths[i]
        let fs = await this.listDir(p)

        let fsFiltered = null
        if( name == "" ) {
          // arr = arr.concat(fs)
          fsFiltered = fs
        } else {
          fsFiltered = fs.filter(e => e.name.indexOf(name) != -1)
        }

        let limit = 99999999
        for(let k=0; k< fsFiltered.length && k<limit ; k++) {
          await this.loadFileData(fsFiltered, k, arr)
          /*
          let item = fsFiltered[k]
          item._pre = null
          item._url = null

          FileOp.Open(item.path).then(res => {
            let ext = item.ext
            ext = ext && ext.length>0 ? ext.substring(1) : ""
      
            item._pre = "data:image/"+ext+";base64,"
            item._url =  item._pre+ res
          })
          arr.push(item)
          */
        }
      }

      this.files = arr
      this.loading = false
      // console.log("files", this.files)
    },
    searchClick() {
      this.searchByName(this.name)
    },

    // 将文件用 base64 的形式读取
    async loadFileData(arr, index, newArr) {
      let item = arr[index]
      let data = await FileOp.Open(item.path).then(res => res)

      // let btData = byteString.stringToByte(data)
      // var f =  new File(btData, item.name)
      let ext = item.ext
      ext = ext && ext.length>0 ? ext.substring(1) : ""

      item._pre = "data:image/"+ext+";base64,"
      item._url =  item._pre+ data

      // console.log("data", data)
      newArr.push(item)
    },

    // 复制到截切版
    async copyToClipboard(item) {
      // 截掉 url 开头的部分 'data:image/jpg;base64,'
      console.log("item", item)
      let base64Data = item._url.substring(item._pre.length)

      // 用 'image/png' 梭哈，格式太多可能不支持
      const blobInput = this.convertBase64ToBlob(base64Data, 'image/png');
      const clipboardItemInput = new ClipboardItem({ 'image/png': blobInput });
      navigator.clipboard.write([clipboardItemInput]);
    },
    // base64 转Blob 对象
    convertBase64ToBlob(base64, type) {
      var bytes = window.atob(base64);
      var ab = new ArrayBuffer(bytes.length);
      var ia = new Uint8Array(ab);
      for (var i = 0; i < bytes.length; i++) {
        ia[i] = bytes.charCodeAt(i);
      }
      return new Blob([ab], { type: type });
    },

    // 查看剪切板的内容
    loadClipboard() {
      navigator.clipboard.read().then(res => {
        console.log("loadClipboard", res)
        let citem = res[0]
        console.log("citem", citem)
        console.log("types", citem.types)

        for(let i=0; i<citem.types.length; i++) {
          let type = citem.types[i]
          citem.getType(type).then(data => {
            // console.log(type, data)

            if(type.startsWith('text/')) {
              let reader = new FileReader()
              reader.onload = () => {
                console.log(type, reader.result)
              }
              reader.readAsBinaryString(data)
            } else {
              console.log(type, data)
            }

          })
        }
      })
    },
    async testCpClipboard() {
      let el = document.getElementById('test-img')
      let canvas = this.convertImageTagToCanvas(el)
      let base64PNG = canvas.toDataURL("image/png")

      // await this.imageToBlob(el.src, (blobData) => {
        // let canvas = this.convertImageTagToCanvas(el)
        // canvas.toBlob((blob) => { 
        //   console.log("blobData", blobData)
        //   // IE复制动图的时候
        //   let testItem = new ClipboardItem({ 
        //     'image/png': blobInput, //blobData,
        //     'text/plain': encodeURIComponent(el.src), // 文件 url
        //     'text/html': el.outerHTML, // 被复制图片的原生 html
        //   })
        //   navigator.clipboard.write([testItem]);
        // }, "image/png", 1.0);  

        // blobData = canvas.toDataURL("image/png");  

        console.log("blobData", base64PNG)
        const blobData = this.convertBase64ToBlob(base64PNG, 'image/png');
              // IE复制动图的时候
        let testItem = new ClipboardItem({ 
          'image/png': blobData,
          'text/plain': encodeURIComponent(el.src), // 文件 url
          'text/html': el.outerHTML, // 被复制图片的原生 html
        })
        navigator.clipboard.write([testItem]);
        
      // })
    },
    // url img地址，图片地址如果是网络图片，网络地址需要处理跨域
    // fn  函数，返回一个blob对象
    imageToBlob (url, fn) {
      if (!url || !fn) return false;
      var xhr = new XMLHttpRequest();
      xhr.open('get', url, true);
      xhr.responseType = 'blob';
      xhr.onload = function () {
        // 注意这里的this.response 是一个blob对象 就是文件对象
        fn(this.status == 200 ? this.response : false);
      }
      xhr.send();
      return true;
    },
    // 把image 转换为 canvas对象  
    convertImageTagToCanvas(image) {  
      image.setAttribute("crossOrigin",'Anonymous')
      // 创建canvas DOM元素，并设置其宽高和图片一样   
      var canvas = document.createElement("canvas");  
      canvas.width = image.width;  
      canvas.height = image.height;  
      // 坐标(0,0) 表示从此处开始绘制，相当于偏移。  
      canvas.getContext("2d").drawImage(image, 0, 0);    
      return canvas;  
    },

  }

}