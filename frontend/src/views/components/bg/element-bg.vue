<template>
  <!-- <img class="mainBg mainBg1" ref="mainBg_1" />
  <img class="mainBg mainBg2" ref="mainBg_2" /> -->
  <div class="mainBg mainBg1" ref="mainBg_1" ></div>
  <div class="mainBg mainBg2" ref="mainBg_2" ></div>
</template>

<script>
// 通过 position: fixed 和 z-index 渲染在最底部
import { loadImage, loadBackground } from '/src/utils/common.js'

export default {
  props: {
    urls: { type: Array, default: ()=>[] }, // 图片源
    interval: { type: Number, default: ()=> 10000 }, // 图片切换间隔， 默认10秒
    random: { type: Boolean, default: ()=> false }, // 随机切换图片
  },
  emits: [ 'beforeChange', 'changed' ],  // 背景图切换事件
  data() {
    return {
      loadedImages: [],
      activeIndex: -1, // 正在展示的图片下标 (loadedImages 的下标)
      lastShow: null,  // '1' 或 '2'.  上一次显示的是背景图1 还是 背景图2

      timer: null,
    }
  },
  watch: {
    urls() {
      // 清除定时器
      this.stopTimer()
      this.activeIndex = -1
      this.loadAllImages()
    }
  },
  mounted() {
    this.loadAllImages()
  },
  destroy() {
    this.stopTimer()
  },
  methods: {
    loadAllImages() {
      // console.log("1")
      if(!this.urls || this.urls.length==0) {
        return
      }

      for(let i=0; i<this.urls.length; i++) {
        let imageUrl = this.urls[i];
        loadImage(imageUrl).then(img => {      
          // 防止重复添加
          if(this.loadedImages.indexOf(img) == -1) {
            this.loadedImages.push(img)
          }
          
          // 开启定时器
          if(!this.timer) {
            this.changeBg(img, true) // 定时器不会立即执行，所以手动执行第一次
            this.setupTimer()
          }

        })
      }
    },

    // 开启定时器
    setupTimer() {
      if(this.timer){
        return
      }

      this.timer = setInterval(()=>{
        if(this.loadedImages.length == 0) {
          this.activeIndex = -1
          return
        }

        this.changeBg()
      }, this.interval)
    },

    /** 调整背景图
     * @param isFirst 是否是第一次加载，urls 发生
     * 
     *  */ 
    changeBg(isFirst) {
      const bg_1 = this.$refs.mainBg_1
      const bg_2 = this.$refs.mainBg_2

      if(this.activeIndex<0 || this.activeIndex>=this.loadedImages.length-1) {
        this.activeIndex = 0
      } else {
        if(this.random && this.loadedImages.length>1) {
          let rindex = Math.random()*this.loadedImages.length
          rindex = parseInt(rindex)
          this.activeIndex = rindex
        } else {
          this.activeIndex++
        }

      }

      let imgUrl = this.loadedImages[this.activeIndex]
      // console.log(this.activeIndex, imgUrl)
      // 基数显示第一个，偶数显示第二个
      let isFirstChange = isFirst ? true : false;
      this.$emit("beforeChange", imgUrl, isFirstChange)

      // 判断显示背景1还是背景2。
      if(!this.lastShow) {
        this.lastShow = '2'
      }

      if(this.lastShow == '2') { // 显示背景图1
        bg_1.style.backgroundImage = `url(${imgUrl})`
        // bg_1.src = `${imgUrl}`
        bg_2.classList.remove('fadeIn')
        bg_2.classList.add('fadeOut')
        bg_1.classList.remove('fadeOut')
        bg_1.classList.add('fadeIn')
        // bg_2.style.backgroundImage = ''
        this.lastShow = '1'
      } else {  // 显示背景图2
        bg_2.style.backgroundImage = `url(${imgUrl})`
        // bg_2.src = `${imgUrl}`
        bg_1.classList.remove('fadeIn')
        bg_1.classList.add('fadeOut')
        bg_2.classList.remove('fadeOut')
        bg_2.classList.add('fadeIn')
        // bg_1.style.backgroundImage = ''
        this.lastShow = '2'
      }

      this.$emit("changed", imgUrl, isFirstChange)
    },

    // 停止计时器
    stopTimer() {
      if(this.timer) {
        clearInterval(this.timer)
      }
    }
  }
}
</script>

<style scoped>
/** 淡入和淡出 */
.fadeOut{
	animation:fadeOut-frame 3s 1;
  animation-fill-mode: forwards; /* 不会重复播放 */
}

@keyframes fadeOut-frame{
   0%{opacity: 1;}
   10%{opacity: 0.8;}
   100%{opacity: 0; }
}

.fadeIn{
	animation:fadeIn-frame 3s 1;
  animation-fill-mode: forwards; /* 不会重复播放 */
}

@keyframes fadeIn-frame{
	0%{opacity:0; }
	100%{opacity: 1;}
	
}

.mainBg{
  position: fixed;
  width: 100vw;
  height: 100vh;
  /*z-index: 0; */
  top: 0;
  left: 0;
  background-size: 100% 100%;
  background-repeat: no-repeat;
  /* background-color: rgba(0, 0, 0, 0.2); */
}
.mainBg1{
  z-index: -1;
}
.mainBg2{
  z-index: -1;
}

@media only screen and (max-width: 500px) {
  .mainBg {
    /* min-width: 800px; */
    min-height: 600px;
    background-position: center;
    /* background-size: unset; */
  }
}

</style>