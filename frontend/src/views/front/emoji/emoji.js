import { Greet, FileOp } from "/wailsjs/index.js"
// import { List } from "/wailsjs/go/env/FileOp"

export default {
  name: 'Emoji',
  components: { },
  data () {
    return {
      files: [], // {isDir, name, path, size}
    }
  },
  mounted() {
    this.listDir("D:\\Picture")
  },
  methods: {
    testList() {
      this.listDir("D:\\Picture")
    },
    async listDir(path) {
      let res = await Greet("1234").then(res => res)
      console.log("wait res", res)

      let fsList = await FileOp.ListImage(path).then(res => res)
      this.files = fsList
      console.log("fsList", fsList)
    },

  }

}