// import { Greet, FileOp, ConfigOp, DbOp, PlatformService } from "@/../wailsjs/index"
import { PlatformService } from "@/../wailsjs/index"
import { AddCircle32Regular, Edit32Filled, Delete28Filled,
  ArrowCounterclockwise28Filled, SubtractCircleArrowForward20Regular } from '@vicons/fluent'
import PlatformForm from './platform-form.vue'
import PlatformDetail from './platform-detail.vue'
import InitCheck from './init-check.vue'

export default {
  name: 'Pwdbox',
  components: { 
    AddCircle32Regular, Edit32Filled, Delete28Filled, ArrowCounterclockwise28Filled, SubtractCircleArrowForward20Regular,
    PlatformForm, PlatformDetail, InitCheck,
  },
  setup() {
    const testData = () => {
      list.value = [
        { name: '测试' },{ name: '测试' },{ name: '测试' },{ name: '测试' },
      ]
    }
  },
  data() {
    return {
      showDatas: false,

      loading: false,
      list: [],
      params: {
        username: null,
        phone: null, 
        email: null,
      },
      paramsPage: {
        page: 0,
        size: 12, // 12的方便显示
        total: 0,
      },
      hasMoreData: true,

      // 防止滚动事件出发太频繁
      lazyLoading: false,
    }
  },
  created() {

  },
  async mounted() { 
    // this.getPageData(1)
    // await this.startLoad()
  },
  // why OptionalAPI ? CompositionAPI makes some messy!
  methods: {
    testData() {
      list.value = [
        { name: '测试' },{ name: '测试' },{ name: '测试' },{ name: '测试' },
      ]
    },

    async getPageData(page) {
      let res = await PlatformService.PageList(null, page, this.paramsPage.size).then(res => res)
      // console.log("List", res)
      return res
    },
    async nextPage() {
      // console.log("nextPage")
      if(!this.hasMoreData) {
        return
      }

      let next = this.paramsPage.page + 1
      this.loading = true
      let res = await this.getPageData(next)
      // console.log("next", next, res)

      if(res.Data && res.Data.length > 0) {
        this.paramsPage.page = res.Page
        this.paramsPage.size = res.Size
        this.paramsPage.total = res.Total
        this.list = this.list.concat(res.Data) 
      //   this.hasMoreData = true
      // } else {
      //   this.hasMoreData = false
      }
      this.hasMoreData = this.paramsPage.page * this.paramsPage.size < this.paramsPage.total
      // console.log("hasMoreData", this.hasMoreData, this.paramsPage.page * this.paramsPage.size, this.paramsPage.total)
      this.loading = false
    },
    

    showPlatformForm(isUpdate, data) {
      let rcom = this.$refs.platformFormComponent
      // data = data ? data : {}
      // console.log("showPlatformForm", rcom)
      let cpData = data ? Object.assign({}, data) : null
      rcom.show(isUpdate, cpData)
    },
    showPlatformDetail(data) {
      let rcom = this.$refs.platformDetailComponent
      // let cpData = data ? Object.assign({}, data) : null
      rcom.show(data.id)
    },
    addPlatform(data) {
      this.list.value.unshift(data)
      this.$refs.platformFormComponent.close()
    },
    updatePlatform(data) {
      let arr = this.list
      if(!arr) {
        this.$refs.platformFormComponent.close()
        return
      }

      for(let i=0; i<arr.length; i++) {
        let item = arr[i]
        if(item.id == data.id) {
          arr[i] = data
          break
        }
      }
      this.$refs.platformFormComponent.close()
    },

    async refresh() {
      await this.startLoad()
      // this.$router.go(0)
    },
    // 一开始加载数据，直到铺满整个页面，或加载完所有数据
    async startLoad() {
      this.paramsPage.page = 0
      this.list = []
      this.hasMoreData = true
      this.showDatas = true

      let count = 1
      // console.log("startLoad")
      while(this.hasMoreData && this.loadMoreIsInViewPort()) {
        await this.lazyLoadMore()
        // console.log("startLoad-count", count)
        count++
      }

    },

    // 返回是否
    async lazyLoadMore () {
      if(!this.hasMoreData || this.lazyLoading) {
        return
      }

      this.lazyLoading = true
      let isShowInView = this.loadMoreIsInViewPort();
      if(isShowInView ) {
        await this.nextPage()
      }
      this.lazyLoading = false
    },
    // 判断加载更多的元素是否出现在视口中
    loadMoreIsInViewPort() {
      const screenHeight = window.innerHeight || document.documentElement.clientHeight
          || document.body.clientHeight;
      let el = document.getElementById("moreEl")
      let rect =  el.getBoundingClientRect()
      const top =rect && rect.top;
      // console.log("top", screenHeight, top, el.clientHeight, top+el.clientHeight)
      
      let isShowInView = top <= screenHeight;
      return isShowInView
    },
    
    async initData() {
      await this.startLoad()
    },

    async clearAesInfo() {
      await sessionStorage.removeItem("pwdbox")
      this.paramsPage.page = 0
      this.paramsPage.total = 0
      this.list = []
      this.showDatas = false
      
      let icRef = this.$refs.initCheckRef
      icRef.showAndCheck()
    }
  }

}