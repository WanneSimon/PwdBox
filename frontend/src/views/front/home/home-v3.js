import axios from 'axios'
import MetaHeader from '/src/components/vue-meta/meta-header.vue'
import ElementBg from '/src/views/components/bg/element-bg.vue'

let baseUrl = import.meta.env.VITE_FRONT_BASE
baseUrl = baseUrl ? baseUrl : ''

export default {
  components: { MetaHeader, ElementBg, },
  setup() {

  },
  data() {
    return {
      sentence: null,
    }
  },
  computed: { 
  },
  mounted() {
    this.randomSentence()
  },
  methods: {
    // 随机哲理句
    randomSentence() {
      // 中文 https://api.xygeng.cn/one
      // 英文
      const randomTypes = {
        zh: 'https://api.xygeng.cn/one',
        en: 'https://api.quotable.io/random'
      }
      
      axios.get(randomTypes.zh).then(res => {
        // console.log("zh", res)
        if(res.data?.data) {
          // zh: {content, created_at, id, name, origin, tag, updated_at}
          this.sentence = res.data.data
        }
      })
    },
    toTools() {
      this.$router.push('/tools')
    }

  }
}