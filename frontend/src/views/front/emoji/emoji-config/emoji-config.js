import { SaveConfig, FileOp, Config } from "/wailsjs/index.js"

export default {
  name: 'Emoji',
  components: { },
  inject: [ "Message" ],
  data () {
    return {
      config: null,

      emojis: [],
      loading: false,
    }
  },
  async mounted() { 
    await this.getConfig()

    console.log("config", this.config)
    console.log("emojis", this.emojis)
  },
  methods: {
    // 获取配置
    async getConfig() {
      this.loading = true
      let conf =  await Config().then(res => res)

      this.config = conf
      
      if(conf?.Emojis) {
        this.emojis = conf?.Emojis.map(e => e + "")
      }

      for(let i=this.emojis.length; i<3; i++) {
        this.pushNewConfig()
      }

      this.loading = false
    },

    async save() {
      let re =  await this.config.Save(res => res)
      let msg = "保存成功"
      if(!re) {
        msg = "保存失败"
      }
      console.log(msg)
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
      let re = await SaveConfig(this.config).then(res => res)
      // console.log("Save", re)
      if(re) {
        this.Message.success("保存成功")
      } else {
        this.Message.error("保存失败")
      }
    }
  }

}