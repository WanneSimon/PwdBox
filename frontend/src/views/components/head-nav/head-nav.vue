<template>
  <div class="nav-head float-block" style="--wails-draggable: drag">
      <label class="icon-label">
        <router-link :to="'/home'" title="Home" target="self">
          <el-button size="small" color="#626aef">Home</el-button>
        </router-link>
      </label>
      <label class="icon-label">
        <router-link :to="'/emoji'" title="Emoji">
          <el-button size="small" color="#626aef">Emoji</el-button>
        </router-link>
      </label>
      <label class="icon-label" title="小工具">
        <router-link :to="'/tools'">
          <el-button size="small" color="#626aef"> 小工具</el-button>
        </router-link>
      </label>
    
      <label class="right-pos">
        <!-- <el-button size="small" type="warning" @click="min" :icon="'Minus'">最小化</el-button>
        <el-button size="small" type="warning" @click="exit" :icon="'Close'">退出</el-button>
        <el-button size="small" @click="goBack" :icon="'ArrowLeftBold'">返回</el-button> -->
        <el-button size="small" @click="goBack" title="返回">
            <el-icon><ArrowLeftBold /></el-icon></el-button>
        <el-button size="small" type="warning" @click="min" title="最小化">
          <el-icon><Minus /></el-icon></el-button>
        <el-button size="small" type="danger" @click="exit" :icon="Close" title="关闭">
          <el-icon><Close /></el-icon></el-button>
      </label>
  </div>
</template>

<script>
import { Exit, Minimises } from "/wailsjs/index.js"
import { Minus, Close, ArrowLeftBold } from '@element-plus/icons-vue'

export default {
    name: 'HeadNav',
    components: { Minus, Close, ArrowLeftBold },
    data () {
        return {
            
        }
    },

    methods: {
      goBack() {
        this.$router.back()
      },
      exit() {
        Exit()
      },
      min(){
        Minimises()
      }
    }
}
</script>

<style scoped lang="scss">
.nav-head{
  border-radius: unset;
  padding: 0rem 0.4rem;
  height: 2.4rem;
  line-height: 2.4rem;
  position: relative;

  padding-left: 10rem;
  padding-right: 10rem;

  text-align: center;
  background-color: rgba($color: #000000, $alpha: 0.3);
  a {
    color: #606266 !important;
  }
  a:hover {
    color: #409eff !important;
  }

  .right-pos{
    position: absolute;
    right: 1px;
    top: 0;
    width: rem;
  }
}

:deep(.el-button+.el-button) {
  margin-left: 1px;
}
</style>
