// import { SaveConfig, FileOp, Config } from "/wailsjs/index.js"
import { FileOp, ConfigOp } from "@/../wailsjs/index"


export default {
  name: 'Emoji',
  components: { },
  emit: [ "saved", "close" ],
  inject: [ "Message", "Noti" ],
  data () {
    return {
      config: null,

      emojis: [],
      loading: false,
    }
  },
  async mounted() { 
    await this.getConfig()
  },
  methods: {
    // 获取配置
    async getConfig() {
      this.loading = true
      // let conf =  await Config().then(res => res)
      let conf = await ConfigOp.Get()

      this.config = conf
      
      if(conf?.Emojis) {
        this.emojis = conf?.Emojis.map(e => e + "")
      }

      for(let i=this.emojis.length; i<3; i++) {
        this.pushNewConfig()
      }

      this.loading = false
    },

    // 
    pushNewConfig() {
      if(!this.emojis) {
        this.emojis = []
      }

      this.emojis.push("")
    },

    async selectFolder(index) {
      let folder = await FileOp.SelectFolder().then(res => res)
      if(folder) {
        // return folder
        this.emojis[index] = folder
      }
    },

    deleteConfig(index) {
      this.emojis.splice(index, 1)
    },
    resetConfig() {
      // console.log("emojis", this.emojis)
      this.emojis = this.config.Emojis
      // console.log("emojis", this.emojis, this.config.Emojis)
    },
    async saveConfig() {
      // console.log("saveConfig", this.config)
      this.config.Emojis = this.emojis.filter(e => e)
      let re = await ConfigOp.SaveAppConfig(this.config).then(res => res)
      // console.log("Save", re)
      if(re) {
        // this.Message.success("保存成功")
        this.Noti.success({ message: "保存成功", position: 'bottom-right',})
        this.$emit("saved") // 不刷新自己组件的信息
      } else {
        // this.Message.error("保存失败")
        this.Noti.error({ message: "保存失败", position: 'bottom-right',})
      }
    },

    back() {
      this.$emit("close") // 关闭当前组件
    }
  }

}