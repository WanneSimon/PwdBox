import { Greet, FileOp, Config } from "/wailsjs/index.js"
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
      let conf =  await Config().then(res => res)
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
          // await this.loadFileData(fsFiltered, k, arr)

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

  }

}