import { reactive, ref } from "vue"
import axios from 'axios'
import { ElPagination, ElImage } from 'element-plus'

export default {
  name: 'TheCatApiGallary',
  components: { ElPagination, ElImage },
  setup(props, context) {
    const url = ref('https://api.thecatapi.com/v1/images/search?limit=10&size=full')
    var historyPage = ref(1)
    var historyList = reactive({data: []}) 
    var list = reactive({data: []})


    const randomPictures = () => {
      list.data = []
      axios.get(url.value)
      .then(res => {
          // console.log("randoms", res)
          list.data = res.data // [{id, url, width, height}]

          if(list.data && list.data.length>0) {
            historyList.data.push(list.data)
          }
          historyPage.value = historyList.data.length
          // console.log("list", list)
      })
    }

    const showHistoryPage = (page) => {
      if(page >0 && page<=historyList.data.length) {
        list.data = historyList.data[page-1 ]
      }
    }

    return {
      list, historyPage, historyList, randomPictures, showHistoryPage,
    }
  },
  mounted() {
    this.randomPictures()
  }
}